package strict

import (
	"io"

	"github.com/ramonpetgrave64/go-html-compose/pkg/doc"
	"github.com/ramonpetgrave64/go-html-compose/pkg/html/attrs"
	"github.com/ramonpetgrave64/go-html-compose/pkg/html/elems"
)

// util

type IAttribute = doc.IAttribute

func toIAttributes[T doc.IAttribute](attrs []T) []doc.IAttribute {
	converted := make([]doc.IAttribute, len(attrs))
	for i, attr := range attrs {
		converted[i] = attr
	}
	return converted
}

type attrWrapper struct {
	IAttribute
	globalAttr
	buttonAttr
	imgAttr
	scriptAttr
}

func newAttrWrapper(attr IAttribute) *attrWrapper {
	return &attrWrapper{IAttribute: attr}
}

func (w *attrWrapper) RenderAttr(wr io.Writer) (err error) {
	return w.IAttribute.RenderAttr(wr)
}

// func (w *attrWrapper) UnimplementedGlobalAttr() {}

// func (w *attrWrapper) UnimplementedImgAttr() {}

// func (w *attrWrapper) UnimplementedButtonAttr() {}

// func (w *attrWrapper) UnimplementedScriptAttr() {}

// elems

type globalAttr interface {
	doc.IAttribute
	global()
}

type imgAttr interface {
	doc.IAttribute
	img()
}

type buttonAttr interface {
	doc.IAttribute
	button()
}

type scriptAttr interface {
	doc.IAttribute
	script()
}

func Button(attrs ...buttonAttr) doc.ContContainerFunc {
	return elems.Button(toIAttributes(attrs)...)
}

func Img(attrs ...imgAttr) doc.IContent {
	return elems.Img(toIAttributes(attrs)...)
}

func Script(attrs ...imgAttr) doc.ContContainerFunc {
	return elems.Script(toIAttributes(attrs)...)
}

// attrs

type roleI interface {
	globalAttr
}

type altI interface {
	imgAttr
}

type srcI interface {
	imgAttr
	scriptAttr
}

type typeI interface {
	buttonAttr
}

type nameI interface {
	buttonAttr
}

func nameA(value string) nameI {
	return newAttrWrapper(attrs.Name(value))
}

func typeA(value string) typeI {
	return newAttrWrapper(attrs.Name(value))
}

func srcA(value string) srcI {
	return newAttrWrapper(attrs.Src(value))
}

func roleA(value string) srcI {
	return newAttrWrapper(attrs.Role(value))
}

// test

func Do() {
	n := nameA("my-name")
	t := typeA("my-type")

	s := srcA("my-src")

	r := roleA("my-role")

	b := Button(n, t)
	b()

	b1 := Button(n, t, s)
	b1()

	b2 := Button(n, t, r)
	b2()

	Img(s)

	script := Script(s, r)
	script()
}
