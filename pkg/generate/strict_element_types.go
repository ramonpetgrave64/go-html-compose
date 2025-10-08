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
	strictElementTypesFilePath               = "../html/strict/internal/types/elems/elems.go"
	strictElementCategoriesFilePath          = "../html/strict/internal/types/elems/categories.go"
	elementContentCategoriesTableCaptionText = "List of element content categories"
)

func generateStrictElementTypes(specContent io.Reader) error {
	specNode, err := html.Parse(specContent)
	if err != nil {
		return err
	}
	elementsTable := getTableByCaption(specNode, elementsTableCaptionText)
	elements := extractElementsFromTable(elementsTable)

	elemMap := make(map[string]*element)
	for _, elem := range elements {
		elemMap[elem.name] = elem
	}

	parentElemMap := make(map[string][]string)

	for _, elem := range elements {
		parentStrings := strings.Split(elem.parents, ";")
		for _, parentString := range parentStrings {
			parentString = strings.TrimSpace(parentString)
			parentString = strings.TrimSuffix(parentString, ".")
			parentString = strings.ReplaceAll(parentString, "*", "")
			parentString = strings.ReplaceAll(parentString, ";", "")

			if _, ok := elemMap[parentString]; ok {
				parentElemMap[elem.name] = append(parentElemMap[elem.name], parentString)
			}

		}
	}

	categories, elementCategoryMap := extractElementCategories(specNode)

	categoryElementParents := extractCategoryElementParents(categories, elements)
	allCategoriesContent := generateCategoryTypesContent(categories, categoryElementParents)

	elemenStructStrings := []string{}
	for _, element := range elements {
		elemenStructStrings = append(
			elemenStructStrings,
			generateStrictElemStructString(
				element.name,
				parentElemMap[element.name],
				elementCategoryMap[element.name],
			),
		)

	}
	allElementsContent := strings.Join(elemenStructStrings, "\n")

	var content bytes.Buffer
	fmt.Fprintf(&content, `%s

package elems

import (
	"io"

	"github.com/ramonpetgrave64/go-html-compose/pkg/doc"
)

type ContContainerFunc[Parent, Child doc.IContent] func(children ...Child) Parent

%s`, doNotEdit, allElementsContent)

	if err := os.WriteFile(strictElementTypesFilePath, content.Bytes(), 0644); err != nil {
		return err
	}

	content.Reset()
	fmt.Fprintf(&content, `%s

package elems

import (
	"github.com/ramonpetgrave64/go-html-compose/pkg/doc"
)

%s`, doNotEdit, allCategoriesContent)

	if err := os.WriteFile(strictElementCategoriesFilePath, content.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

func generateStrictElemStructString(element string, parentElements []string, parentCategories []string) string {
	structName := kebabToPascal(element)
	parentTypeStrings := []string{}
	for _, parent := range parentElements {
		parentTypeStrings = append(parentTypeStrings, fmt.Sprintf("\n\t%sChild", kebabToPascal(parent)))
	}
	parentCategoryStrings := []string{}
	for _, category := range parentCategories {
		parentCategoryStrings = append(parentCategoryStrings, fmt.Sprintf("\n\t%sContent", phraseToPascal(category)))
	}
	parentTypes := strings.Join(parentTypeStrings, "")
	parentCategoriesTypes := strings.Join(parentCategoryStrings, "")
	return fmt.Sprintf(`type %s struct {
	doc.IContent%s%s
}

func (e %s) RenderConent(wr io.Writer) (err error) {
	return e.IContent.RenderConent(wr)
}

type %sChild interface {
	doc.IContent
	is%sChild()
}
`, structName, parentTypes, parentCategoriesTypes, structName, structName, structName)
}

func phraseToPascal(phrase string) string {
	phrase = kebabToPascal(phrase)
	pascalName := ""
	for _, part := range strings.Split(phrase, " ") {
		pascalName += strings.ToUpper(part[0:1]) + part[1:]
	}
	return pascalName
}

func generateCategoryTypesContent(categories []string, categoriesMap map[string][]string) string {
	allContent := []string{}
	// for category, parents := range categoriesMap {
	for _, category := range categories {
		parents := categoriesMap[category]
		name := phraseToPascal(category)
		parentTypesContent := []string{}
		for _, parent := range parents {
			parent := kebabToPascal(parent)
			parentTypesContent = append(parentTypesContent, fmt.Sprintf("\n\t%sChild", parent))
		}
		parentTypes := strings.Join(parentTypesContent, "")
		content := fmt.Sprintf(`type %sContent interface {
	doc.IContent
	is%sContent()%s
}
`, name, name, parentTypes)
		allContent = append(allContent, content)
	}
	return strings.Join(allContent, "\n")
}

func extractElementCategories(specNode *html.Node) ([]string, map[string][]string) {
	table := getTableByCaption(specNode, elementContentCategoriesTableCaptionText)
	tbody := *seqSelect(table.ChildNodes(), func(node *html.Node) bool {
		return node.Data == "tbody"
	})
	categories := []string{}
	elementCategoryMap := make(map[string][]string)
	for tr := range tbody.ChildNodes() {
		rowNodes := seqSlice(tr.ChildNodes())
		catergory := digAllText(rowNodes[0])
		catergory = normalizeCategory(catergory)
		categories = append(categories, catergory)
		elements := strings.Split(digAllText(rowNodes[1]), ";")
		for _, element := range elements {
			element = strings.TrimSpace(element)
			categories := elementCategoryMap[element]
			categories = append(categories, catergory)
			elementCategoryMap[element] = categories
		}
	}
	return categories, elementCategoryMap
}

func extractCategoryElementParents(categories []string, elements []*element) map[string][]string {
	categoryElementParentsMap := make(map[string][]string)
	for _, category := range categories {
		categoryElementParentsMap[category] = []string{}
	}
	for _, element := range elements {
		children := strings.Split(element.children, ";")
		for _, child := range children {
			child = normalizeCategory(child)
			if sliceContains(categories, child) {
				parentElements := categoryElementParentsMap[child]
				parentElement := kebabToPascal(element.name)
				parentElements = append(parentElements, parentElement)
				categoryElementParentsMap[child] = parentElements
			}
		}
	}
	return categoryElementParentsMap
}

func normalizeCategory(category string) string {
	category = strings.TrimSuffix(category, "*")
	category = strings.TrimSuffix(category, "content")
	category = strings.TrimSuffix(category, "elements")
	category = strings.TrimSpace(category)
	category = kebabToPascal(category)
	return category
}
