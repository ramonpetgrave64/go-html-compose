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
	strictElemTypesFilePath = "../html/strict/elems/types.go"
)

func generateStrictElements(specContent io.Reader) error {
	specNode, err := html.Parse(specContent)
	if err != nil {
		return err
	}
	elementsTable := getTableByCaption(specNode, elementsTableCaptionText)
	elements := extractElementsFromTable(elementsTable)

	var typeDefs, globalMethods []string
	globalMethods = append(globalMethods, "isGlobalAttribute()")
	for _, elem := range elements {
		typeDef := makeStrictElementType(elem)
		typeDefs = append(typeDefs, typeDef)

		funcName := kebabToPascal(elem.name)
		typeName := fmt.Sprintf("%sAttribute", funcName)
		globalMethods = append(globalMethods, fmt.Sprintf("is%s()", typeName))
	}

	var content bytes.Buffer
	if _, err = fmt.Fprintf(&content, `%s

package elems

import "github.com/ramonpetgrave64/go-html-compose/pkg/doc"

// GlobalAttribute can be used with any element.
type GlobalAttribute interface {
	doc.IAttribute
	%s
}

// Attribute types for specific elements
type (%s
)`, doNotEdit, strings.Join(globalMethods, "\n\t"), strings.Join(typeDefs, "\n\t")); err != nil {
		return err
	}

	return os.WriteFile(strictElemTypesFilePath, content.Bytes(), 0644)
}

func makeStrictElementType(elem *element) string {
	funcName := kebabToPascal(elem.name)
	typeName := fmt.Sprintf("%sAttribute", funcName)
	return fmt.Sprintf(`
	// %s is an attribute that can be used with the %s element.
	%s interface {
		doc.IAttribute
		is%s()
	}`, typeName, elem.name, typeName, typeName)
}
