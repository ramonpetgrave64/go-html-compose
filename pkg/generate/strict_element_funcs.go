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
	strictElemFuncsFilePath = "../html/strict/elems/elems.go"
)

func generateStrictElementFuncs(specContent io.Reader) error {
	specNode, err := html.Parse(specContent)
	if err != nil {
		return err
	}
	elementsTable := getTableByCaption(specNode, elementsTableCaptionText)
	elements := extractElementsFromTable(elementsTable)

	allElemFuncs := []string{}
	for _, elem := range elements {
		allElemFuncs = append(allElemFuncs, makeStrictElementFunc(elem))
	}

	allElementsContent := strings.Join(allElemFuncs, "\n")

	var content bytes.Buffer
	if _, err = fmt.Fprintf(&content, `%s

package elems

import (
	"github.com/ramonpetgrave64/go-html-compose/pkg/doc"
	"github.com/ramonpetgrave64/go-html-compose/pkg/html/elems"
)

// toIAttributes converts a slice of a specific attribute type to a slice of IAttribute.
func toIAttributes[T doc.IAttribute](attrs []T) []doc.IAttribute {
	iAttrs := make([]doc.IAttribute, len(attrs))
	for i, attr := range attrs {
		iAttrs[i] = attr
	}
	return iAttrs
}

%s
`, doNotEdit, allElementsContent); err != nil {
		return err
	}

	return os.WriteFile(strictElemFuncsFilePath, content.Bytes(), 0644)
}

func makeStrictElementFunc(elem *element) string {
	funcName := kebabToPascal(elem.name)
	attrVarName := strings.ToLower(funcName[:1]) + funcName[1:] + "Attrs"
	strictAttrType := fmt.Sprintf("%sAttribute", funcName)

	returnType := "doc.ContContainerFunc"
	if elem.isUnit {
		returnType = "doc.IContent"
	}

	doc := makeElemDoc(elem)

	return fmt.Sprintf(`%s
func %s(%s ...%s) %s {
	return elems.%s(toIAttributes(%s)...)
}
`, doc, funcName, attrVarName, strictAttrType, returnType, funcName, attrVarName)
}
