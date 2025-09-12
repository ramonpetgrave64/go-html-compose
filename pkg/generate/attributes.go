package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

const (
	regularAttributesTableId      = "attributes-1"
	booleanAttributeTableValue    = "Boolean attribute"
	eventHandlerAttributesTableId = "ix-event-handlers"
	eventHandlerAttrTableValue    = "Event handler content attribute"

	dataPropSpecialCase = `// DataProp
// Element(s): Global attribute.
// Description: Lets you attach custom attributes to an HTML element.
// Value: Text.
func DataProp(property, value string) doc.IAttribute {
	return doc.Attr("data-"+property, value)
}
`
	ariaPropSpecialCase = `// AriaProp
// Element(s): Global attribute.
// Description: Sets aria-* properties.
// Value: Text.
func AriaProp(property, value string) doc.IAttribute {
	return doc.Attr("aria-"+property, value)
}
`

	roleSpecialCase = `// Role
// Element(s): Global attribute.
// Description: Defines an explicit role for an element for use by assistive technologies.
// Value: Text.
func Role(value string) doc.IAttribute {
	return doc.Attr("role", value)
}
`
)

type attribute struct {
	name           string
	propSets       []*propSet
	isBoolean      bool
	isEventHandler bool
}

type propSet struct {
	elements    string
	description string
	value       string
}

func generateAttributes(specContent io.Reader) error {
	specNode, err := html.Parse(specContent)
	if err != nil {
		return err
	}

	specialAttributesContent := strings.Join(
		[]string{ariaPropSpecialCase, dataPropSpecialCase, roleSpecialCase},
		"\n",
	)

	regularAttributesTable := getTableById(specNode, regularAttributesTableId)
	regularAttributes := extractAttributesFromTable(regularAttributesTable)
	regularAttributeFuncs := []string{}
	for _, attr := range regularAttributes {
		attributeFunc := makeAttributeFunc(attr)
		regularAttributeFuncs = append(regularAttributeFuncs, attributeFunc)
	}
	regularAttributeContent := strings.Join(regularAttributeFuncs, "\n")

	eventHandlerAttributesTable := getTableById(specNode, eventHandlerAttributesTableId)
	eventHanlderAttributes := extractAttributesFromTable(eventHandlerAttributesTable)
	eventHandlerAttributeFuncs := []string{}
	for _, attr := range eventHanlderAttributes {
		attributeFunc := makeAttributeFunc(attr)
		eventHandlerAttributeFuncs = append(eventHandlerAttributeFuncs, attributeFunc)
	}
	eventHandlerAttributeContent := strings.Join(eventHandlerAttributeFuncs, "\n")

	var content bytes.Buffer
	if _, err = fmt.Fprintf(&content, `%s

// Package attrs contains auto-generated Attr functions from the spec at
// https://html.spec.whatwg.org/multipage/indices.html.
// Some attributes will have multiple specifications that depend on the element with which they are used.
package attrs

import "github.com/ramonpetgrave64/go-html-compose/pkg/doc"

// Special attributes

%s
// Regular Attributes

%s
// Event Handler Attributes

%s`, doNotEdit, specialAttributesContent, regularAttributeContent, eventHandlerAttributeContent); err != nil {
		return err
	}
	goFile, err := os.Create("../html/attrs/attrs.go")
	if err != nil {
		return err
	}
	if _, err := goFile.Write(content.Bytes()); err != nil {
		return err
	}
	if err := goFile.Close(); err != nil {
		return err
	}
	return nil
}

func extractAttributesFromTable(table *html.Node) []*attribute {
	tbody := *seqSelect(table.ChildNodes(), func(node *html.Node) bool {
		return node.Data == "tbody"
	})
	attributesMap := map[string]*attribute{}
	attributesSlice := []*attribute{}
	for tr := range tbody.ChildNodes() {
		rowNodes := seqSlice(tr.ChildNodes())
		name := digChildData(rowNodes[0], "code").FirstChild.Data
		value := strings.TrimSpace(digAllText(rowNodes[3]))
		isBoolean := value == booleanAttributeTableValue
		isEventHandler := value == eventHandlerAttrTableValue
		props := &propSet{
			elements:    strings.TrimSpace(digAllText(rowNodes[1])),
			description: strings.TrimSpace(digAllText(rowNodes[2])),
			value:       value,
		}
		if attr, ok := attributesMap[name]; !ok {
			attributesMap[name] = &attribute{
				name:           name,
				propSets:       []*propSet{props},
				isBoolean:      isBoolean,
				isEventHandler: isEventHandler,
			}
			attributesSlice = append(attributesSlice, attributesMap[name])
		} else {
			if attr.isBoolean != isBoolean {
				panic(fmt.Errorf("expected all propsets to have equal isBoolean: %s", attr.name))
			}
			attr.propSets = append(attr.propSets, props)
		}
	}
	return attributesSlice
}

func makeAttributeFunc(attr *attribute) string {
	funcName := kebabToPascal(attr.name)
	valueType := "string"
	attrFunc := "doc.Attr"
	doc := makeAttrDoc(attr)
	if attr.isBoolean {
		valueType = "bool"
		attrFunc = "doc.BooleanAttr"
	} else if attr.isEventHandler {
		attrFunc = "doc.RawAttr"
	}

	return fmt.Sprintf(
		`%s
func %s(value %s) doc.IAttribute {
	return %s("%s", value)
}
`,
		doc, funcName, valueType, attrFunc, attr.name,
	)
}

func makeAttrDoc(attr *attribute) string {
	head := fmt.Sprintf(`// %s
//`, kebabToPascal(attr.name))
	docSets := []string{}
	for _, props := range attr.propSets {
		docSet := fmt.Sprintf(`
// Element(s): %s.
//
// Description: %s.
//
// Value: %s.`, props.elements, props.description, props.value)
		docSets = append(docSets, docSet)
	}
	doc := fmt.Sprintf(`%s%s`, head, strings.Join(docSets, "\n//"))
	return doc
}
