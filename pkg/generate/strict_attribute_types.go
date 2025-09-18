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
	strictAttrTypesFilePath = "../html/strict/internal/types/attrs/attrs.go"
)

func generateStrictAttributeTypes(specContent io.Reader) error {
	specNode, err := html.Parse(specContent)
	if err != nil {
		return err
	}

	regularAttributesTable := getTableById(specNode, regularAttributesTableId)
	regularAttributes := extractAttributesFromTable(regularAttributesTable)

	eventHandlerAttributesTable := getTableById(specNode, eventHandlerAttributesTableId)
	eventHandlerAttributes := extractAttributesFromTable(eventHandlerAttributesTable)

	allAttributes := append(regularAttributes, eventHandlerAttributes...)

	strictAttrStructs := []string{}
	for _, attr := range allAttributes {
		strictAttrStructs = append(strictAttrStructs, makeStrictAttributeStruct(attr))
	}

	allStrictAttrsContent := strings.Join(strictAttrStructs, "\n\n")

	var content bytes.Buffer
	if _, err = fmt.Fprintf(&content, `%s

package attrs

import (
	"io"

	"github.com/ramonpetgrave64/go-html-compose/pkg/doc"
)

%s
`, doNotEdit, allStrictAttrsContent); err != nil {
		return err
	}

	return os.WriteFile(strictAttrTypesFilePath, content.Bytes(), 0644)
}

func makeStrictAttributeStruct(attr *attribute) string {
	structName := kebabToPascal(attr.name)

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

	var embeddedTypes []string
	if isGlobal {
		embeddedTypes = append(embeddedTypes, "GlobalAttribute")
	} else {
		for _, elName := range allElements {
			embeddedTypes = append(embeddedTypes, fmt.Sprintf("%sAttribute", kebabToPascal(elName)))
		}
	}

	return fmt.Sprintf(`type %s struct {
	doc.IAttribute
	%s
}

func (a %s) RenderAttr(wr io.Writer) (err error) {
	return a.IAttribute.RenderAttr(wr)
}`, structName, strings.Join(embeddedTypes, "\n\t"), structName)
}
