package elems

import "github.com/ramonpetgrave64/go-html-compose/pkg/doc"

// toIAttributes converts a slice of a specific attribute type to a slice of IAttribute.
func toIAttributes[T doc.IAttribute](attrs []T) []doc.IAttribute {
	iAttrs := make([]doc.IAttribute, len(attrs))
	for i, attr := range attrs {
		iAttrs[i] = attr
	}
	return iAttrs
}

func toIContent[T doc.IContent](content []T) []doc.IContent {
	iContent := make([]doc.IContent, len(content))
	for i, item := range content {
		iContent[i] = item
	}
	return iContent
}
