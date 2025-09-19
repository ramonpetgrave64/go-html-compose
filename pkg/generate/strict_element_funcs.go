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
	"github.com/ramonpetgrave64/go-html-compose/pkg/html/strict/internal/types/attrs"
	types "github.com/ramonpetgrave64/go-html-compose/pkg/html/strict/internal/types/elems"
)

%s
`, doNotEdit, allElementsContent); err != nil {
		return err
	}

	return os.WriteFile(strictElemFuncsFilePath, content.Bytes(), 0644)
}

func makeStrictElementFunc(elem *element) string {
	funcName := kebabToPascal(elem.name)
	strictAttrType := fmt.Sprintf("%sAttribute", funcName)

	// returnType := "doc.ContContainerFunc"
	returnType := fmt.Sprintf("types.ContContainerFunc[types.%s, types.%sChild]", funcName, funcName)
	if elem.isUnit {
		returnType = "doc.IContent"
	}

	doc := makeElemDoc(elem)
	if elem.isUnit {
		return fmt.Sprintf(`%s
func %s(attrs ...attrs.%s) %s {
	return types.%s{IContent: elems.%s(toIAttributes(attrs)...)}
}
`, doc, funcName, strictAttrType, returnType, funcName, funcName)
	}
	return fmt.Sprintf(`%s
func %s(attrs ...attrs.%s) %s {
	return func(children ...types.%sChild) types.%s {
		return types.%s{IContent: elems.%s(toIAttributes(attrs)...)(toIContent(children)...)}
	}
}
`, doc, funcName, strictAttrType, returnType, funcName, funcName, funcName, funcName)
}

// // Ul
// // Description: List.
// // Parents: flow.
// // Children: li; script-supporting elements.
// // Attributes: globals
// func Ul(attrs ...attrs.UlAttribute) types.ContContainerFunc[types.Ul, types.UlChild] {
// 	return func(children ...types.UlChild) types.Ul {
// 		return types.Ul{IContent: elems.Ul(toIAttributes(attrs)...)(toIContent(children)...)}
// 	}
// }
