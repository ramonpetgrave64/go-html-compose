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
	tagsFilePath = "../tag/tags/tags.go"

	elementsTableCaptionText  = "List of elements"
	unitElementsChildrenValue = "empty"

	docTypeSpecialCase = `// Doctype is a required preamble.
var Doctype = tag.UnitTag("!DOCTYPE html")`

	lastNonElementName = "autonomous custom elements"
)

type element struct {
	name        string
	description string
	parents     string
	children    string
	attributes  string
	isUnit      bool
}

func generateElements(specContent io.Reader) error {
	specNode, err := html.Parse(specContent)
	if err != nil {
		return err
	}
	elementsTable := getTableByCaption(specNode, elementsTableCaptionText)
	elements := extractElementsFromTable(elementsTable)
	allElemFuncs := []string{}
	for _, elem := range elements {
		allElemFuncs = append(allElemFuncs, makeElementFunc(elem))
	}
	allElementsContent := strings.Join(allElemFuncs, "\n")
	var content bytes.Buffer
	if _, err = fmt.Fprintf(&content, `%s

package tags

import "go-html-compose/pkg/attr"
import "go-html-compose/pkg/tag"

// Special Elements

%s

// Regular Elements

%s`, doNotEdit, docTypeSpecialCase, allElementsContent); err != nil {
		return err
	}
	goFile, err := os.Create(tagsFilePath)
	if err != nil {
		return err
	}
	defer goFile.Close()
	if _, err := goFile.Write(content.Bytes()); err != nil {
		return err
	}
	return nil
}

func extractElementsFromTable(table *html.Node) []*element {
	tbody := *seqSelect(table.ChildNodes(), func(node *html.Node) bool {
		return node.Data == "tbody"
	})
	elements := []*element{}
	for tr := range tbody.ChildNodes() {
		rowNodes := seqSlice(tr.ChildNodes())
		// special case for the last non-element
		name := digAllText(rowNodes[0])
		if name == lastNonElementName {
			continue
		}

		// special case for h1, h2, ..., h6, which share a single row
		if !strings.Contains(name, ",") {
			// special case for math[MathML] and [SVG]svg
			n1 := digAllText(digDescendantData(rowNodes[0], "a"))
			n2 := digAllText(digDescendantData(rowNodes[0], "code"))
			name = n1
			if len(n2) < len(n1) {
				name = n2
			}
		}

		description := digAllText(rowNodes[1])
		parents := digAllText(rowNodes[3])
		children := digAllText(rowNodes[4])
		attributes := digAllText(rowNodes[5])
		isUnit := children == unitElementsChildrenValue
		// special case for h1, h2, ..., h6, which share a single row
		for _, subName := range strings.Split(name, ",") {
			elem := &element{
				name:        strings.TrimSpace(subName),
				description: description,
				parents:     parents,
				children:    children,
				attributes:  attributes,
				isUnit:      isUnit,
			}
			elements = append(elements, elem)
		}
	}
	return elements
}

func makeElementFunc(elem *element) string {
	funcName := kebabToPascal(elem.name)
	tagFunc := "tag.ParentTag"
	returnType := "tag.ContentFunc"
	doc := makeElemDoc(elem)
	if elem.isUnit {
		tagFunc = "tag.UnitTag"
		returnType = "*tag.UnitTagStruct"
	}
	return fmt.Sprintf(
		`%s
func %s(attrs ...*attr.AttributeStruct) %s {
	return %s("%s", attrs...)
}
`,
		doc, funcName, returnType, tagFunc, elem.name,
	)
}

func makeElemDoc(elem *element) string {
	return fmt.Sprintf(`// %s
// Description: %s.
// Parents: %s.
// Children: %s.
// Attributes: %s`,
		kebabToPascal(elem.name), elem.description, elem.parents, elem.children, elem.attributes)
}

// func makeAttrDoc(attr *attribute) string {
// 	docSets := []string{}
// 	for _, props := range attr.propSets {
// 		docSet := fmt.Sprintf(`
// // Element(s): %s.
// // Description: %s.
// // Value: %s.`, props.elements, props.description, props.value)
// 		docSets = append(docSets, docSet)
// 	}
// 	doc := fmt.Sprintf(`// %s %s`, kebabToPascal(attr.name), strings.Join(docSets, "\n//"))
// 	return doc
// }
