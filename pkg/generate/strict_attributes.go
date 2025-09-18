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
	strictAttrsFilePath = "../html/strict/attrs/attrs.go"
)

func generateStrictAttributes(specContent io.Reader) error {
	specNode, err := html.Parse(specContent)
	if err != nil {
		return err
	}

	regularAttributesTable := getTableById(specNode, regularAttributesTableId)
	regularAttributes := extractAttributesFromTable(regularAttributesTable)

	eventHandlerAttributesTable := getTableById(specNode, eventHandlerAttributesTableId)
	eventHandlerAttributes := extractAttributesFromTable(eventHandlerAttributesTable)

	allAttributes := append(regularAttributes, eventHandlerAttributes...)

	strictAttrFuncs := []string{}
	for _, attr := range allAttributes {
		strictAttrFuncs = append(strictAttrFuncs, makeStrictAttributeFunc(attr))
	}

	allStrictAttrsContent := strings.Join(strictAttrFuncs, "\n")

	var content bytes.Buffer
	if _, err = fmt.Fprintf(&content, `%s

package attrs

import (
	"github.com/ramonpetgrave64/go-html-compose/pkg/html/attrs"
	types "github.com/ramonpetgrave64/go-html-compose/pkg/html/strict/internal/types/attrs"
)

%s
`, doNotEdit, allStrictAttrsContent); err != nil {
		return err
	}

	return os.WriteFile(strictAttrsFilePath, content.Bytes(), 0644)
}

func makeStrictAttributeFunc(attr *attribute) string {
	funcName := kebabToPascal(attr.name)
	underlyingFuncName := funcName

	valueType := "string"
	if attr.isBoolean {
		valueType = "bool"
	}

	returnType := fmt.Sprintf("types.%s", funcName)

	doc := makeAttrDoc(attr)
	underlyingFuncCall := fmt.Sprintf("types.%s{IAttribute: attrs.%s(value)}", funcName, underlyingFuncName)

	return fmt.Sprintf(`%s
func %s(value %s) %s {
	return %s
}
`, doc, funcName, valueType, returnType, underlyingFuncCall)
}
