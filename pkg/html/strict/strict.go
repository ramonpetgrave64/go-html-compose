package strict

import (
	"github.com/ramonpetgrave64/go-html-compose/pkg/doc"
	"github.com/ramonpetgrave64/go-html-compose/pkg/html/attrs"
	"github.com/ramonpetgrave64/go-html-compose/pkg/html/elems"
)

// toIAttributes converts a slice of a specific attribute type to a slice of IAttribute.
func toIAttributes[T IAttribute](attrs []T) []IAttribute {
	iAttrs := make([]IAttribute, len(attrs))
	for i, attr := range attrs {
		iAttrs[i] = attr
	}
	return iAttrs
}

// Attribute type definitions for specific elements
type (
	// IAttribute is a convenience alias for doc.IAttribute.
	IAttribute = doc.IAttribute

	// GlobalAttribute can be used with any element.
	GlobalAttribute interface{ IAttribute }

	// ButtonAttribute is an attribute for <button> elements.
	ButtonAttribute interface{ IAttribute }

	// ImgAttribute is an attribute for <img> elements.
	ImgAttribute interface{ IAttribute }

	// ScriptAttribute is an attribute for <script> elements.
	ScriptAttribute interface{ IAttribute }
)

// Element wrapper functions with strict attribute checking

// Button creates a <button> element and only accepts ButtonAttribute types.
func Button(buttonAttrs ...ButtonAttribute) doc.ContContainerFunc {
	return elems.Button(toIAttributes(buttonAttrs)...)
}

// Img creates an <img> element and only accepts ImgAttribute types.
func Img(imgAttrs ...ImgAttribute) doc.IContent {
	return elems.Img(toIAttributes(imgAttrs)...)
}

// Script creates a <script> element and only accepts ScriptAttribute types.
func Script(scriptAttrs ...ScriptAttribute) doc.ContContainerFunc {
	return elems.Script(toIAttributes(scriptAttrs)...)
}

// Attribute constructor functions

// Global attributes

func Hidden(value string) GlobalAttribute { return attrs.Hidden(value) }

// Element-specific attributes

func Alt(value string) ImgAttribute { return attrs.Alt(value) }

func Name(value string) ButtonAttribute { return attrs.Name(value) }

func Src(value string) interface {
	ImgAttribute
	ScriptAttribute
} {
	return attrs.Src(value)
}

func Type(value string) interface {
	ButtonAttribute
	ScriptAttribute
} {
	return attrs.Type(value)
}

// Example usage demonstrating type safety
func Do() {
	// Valid attribute usage
	Button(
		Name("my-button"),
		Type("submit"),
		Hidden("hidden"), // Global attributes are accepted
	)

	Img(
		Src("image.png"),
		Alt("An image"),
		Hidden("hidden"),
	)

	Script(
		Src("main.js"),
		Type("module"),
		Hidden("hidden"),
	)

	// The following lines would cause a compile-time error
	// because the attribute is not valid for the element.

	// Button(Alt("this is for images")) // COMPILE ERROR
	// Img(Name("this is for buttons")) // COMPILE ERROR
}
