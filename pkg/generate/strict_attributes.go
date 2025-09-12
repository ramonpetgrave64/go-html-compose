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
	"github.com/ramonpetgrave64/go-html-compose/pkg/html/strict/elems"
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

	allElements := []string{}
	isGlobal := false
	for _, ps := range attr.propSets {
		if ps.elements == "HTML elements" {
			isGlobal = true
			break
		}
		elems := strings.Split(ps.elements, ";")
		for _, el := range elems {
			trimmed := strings.TrimSpace(el)
			// The spec sometimes includes contextual information in parentheses for an
			// element, e.g., "source (in picture)". This is not valid for a Go
			// identifier, so we strip it out.
			if i := strings.Index(trimmed, "("); i != -1 {
				trimmed = trimmed[:i]
			}
			trimmed = strings.TrimSpace(trimmed)
			// The strict package is not meant to work with custom elements, so we skip them.
			if trimmed == "" || trimmed == "form-associated custom elements" {
				continue
			}
			allElements = append(allElements, trimmed)
		}
	}

	var types []string
	if isGlobal {
		types = append(types, "elems.GlobalAttribute")
	} else {
		for _, elName := range allElements {
			types = append(types, fmt.Sprintf("elems.%sAttribute", kebabToPascal(elName)))
		}
	}
	returnType := fmt.Sprintf("interface{\n\t%s\n}", strings.Join(types, "\n\t"))

	doc := makeAttrDoc(attr)
	underlyingFuncCall := fmt.Sprintf("newAttrWrapper(attrs.%s(value))", underlyingFuncName)

	return fmt.Sprintf(`%s
func %s(value %s) %s {
	return %s
}
`, doc, funcName, valueType, returnType, underlyingFuncCall)
}
