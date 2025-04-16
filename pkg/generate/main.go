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
	return sliceSelect(seqSlice[S, T](seq), f)
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
