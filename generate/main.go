package main

import (
	"bytes"
	"fmt"
	"io"
	"iter"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

const (
	dataPropSpecialCase = `// DataProp
// Element(s): Global attribute.
// Description: Lets you attach custom attributes to an HTML element.
// Value: Text.
func DataProp(property, value string) *doc.AttrStruct {
	return doc.Attr("data-" + property, value)
}`
	ariaPropSpecialCase = `// AriaProp
// Element(s): Global attribute.
// Description: Sets aria-* properties.
// Value: Text.
func AriaProp(property, value string) *doc.AttrStruct {
	return doc.Attr("aria-" + property, value)
}`

	RoleSpecialCase = `// Role
// Element(s): Global attribute.
// Description: Defines an explicit role for an element for use by assistive technologies.
// Value: Text.
func Role(value string) *doc.AttrStruct {
	return doc.Attr("role", value)
}`

	specURL = "https://html.spec.whatwg.org/multipage/indices.html"

	regularAttributesTableId      = "attributes-1"
	booleanAttributeTableValue    = "Boolean attribute"
	eventHandlerAttributesTableId = "ix-event-handlers"
	eventHandlerAttrTableValue    = "Event handler content attribute"

	doNotEdit = `// Code generated by "go run -C ../generate ./cmd/generate/"; DO NOT EDIT.
// HTML spec at https://html.spec.whatwg.org/multipage/indices.html`
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

func main() {
	if err := generate(); err != nil {
		panic(err)
	}
}

func generate() error {
	specFile, err := downloadFile(specURL, "spec.html")
	if err != nil {
		return err
	}
	defer specFile.Close()
	content, err := io.ReadAll(specFile)
	if err != nil {
		return err
	}
	if err := generateAttributes(bytes.NewReader(content)); err != nil {
		panic(err)
	}
	if err := generateElements(bytes.NewReader(content)); err != nil {
		panic(err)
	}
	return nil
}

func generateAttributes(specContent io.Reader) error {
	specNode, err := html.Parse(specContent)
	if err != nil {
		return err
	}
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

import "go-html-compose/pkg/doc"

// Special attributes

%s

%s

%s

// Regular Attributes

%s

// Event Handler Attributes

%s`, doNotEdit, ariaPropSpecialCase, dataPropSpecialCase, RoleSpecialCase, regularAttributeContent, eventHandlerAttributeContent); err != nil {
		return err
	}
	goFile, err := os.Create("../pkg/html/attrs/attrs.go")
	if err != nil {
		return err
	}
	defer goFile.Close()
	if _, err := goFile.Write(content.Bytes()); err != nil {
		return err
	}
	return nil
}

func getTableByCaption(doc *html.Node, captionText string) *html.Node {
	desc := seqSlice(doc.Descendants())
	return *sliceSelect(desc, func(node *html.Node) bool {
		caption := digChildData(node, "caption")
		if caption != nil && digAllText(caption) == captionText {
			return true
		}
		return false
	})
}

func getTableById(doc *html.Node, id string) *html.Node {
	return *seqSelect(doc.Descendants(), func(node *html.Node) bool {
		if node.Type != html.ElementNode {
			return false
		}
		return sliceSelect(node.Attr, func(attr html.Attribute) bool {
			return attr.Key == "id" && attr.Val == id
		}) != nil
	})
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
func %s(value %s) *doc.AttrStruct {
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

func seqSlice[S iter.Seq[T], T any](seq S) []T {
	slice := []T{}
	for item := range seq {
		slice = append(slice, item)
	}
	return slice
}

func sliceSelect[S ~[]T, T any](slice S, f func(T) bool) *T {
	for _, item := range slice {
		if f(item) {
			return &item
		}
	}
	return nil
}

func seqSelect[S iter.Seq[T], T any](seq S, f func(T) bool) *T {
	return sliceSelect(seqSlice(seq), f)
}

func downloadFile(url, output string) (*os.File, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}
	file, err := os.Create(output)
	if err != nil {
		return nil, err
	}
	// content, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }
	// cleanedContent := bytes.ReplaceAll(content, []byte(`"<code`), []byte(`<code`))
	// cleanedContent = bytes.ReplaceAll(cleanedContent, []byte(`</code>";`), []byte(`</code>`))
	// cleanedContent = bytes.ReplaceAll(cleanedContent, []byte(`</code>"`), []byte(`</code>`))
	// if _, err = file.Write(cleanedContent); err != nil {
	// 	return nil, err
	// }
	if _, err = io.Copy(file, resp.Body); err != nil {
		return nil, err
	}
	if err := file.Close(); err != nil {
		return nil, err
	}
	return os.Open(output)
}

// kebabToPascal converts a kebab-case name to pascal case.
func kebabToPascal(name string) string {
	pascalName := ""
	for _, part := range strings.Split(name, "-") {
		pascalName += strings.ToUpper(part[0:1]) + part[1:]
	}
	return pascalName
}

func digAllText(node *html.Node) string {
	texts := []string{}
	for child := range node.Descendants() {
		if child.Type == html.TextNode {
			texts = append(texts, child.Data)
		}
	}
	if len(texts) == 1 {
		return texts[0]
	}
	joinedText := strings.Join(texts, "")
	cleanedText := strings.Join(strings.Fields(joinedText), " ")
	return cleanedText
}

func digDescendantData(node *html.Node, datum ...string) *html.Node {
	child := node
	for _, d := range datum {
		found := false
		for n := range child.Descendants() {
			if n.Data == d {
				child = n
				found = true
				break
			}
		}
		if !found {
			return nil
		}
	}
	return child
}

func digChildData(node *html.Node, datum ...string) *html.Node {
	child := node
	for _, d := range datum {
		found := false
		for n := range child.ChildNodes() {
			if n.Data == d {
				child = n
				found = true
				break
			}
		}
		if !found {
			return nil
		}
	}
	return child
}
