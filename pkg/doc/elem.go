// Package doc has tools to construct HTML content.
// See https://html.spec.whatwg.org/multipage/dom.html#kinds-of-content.
package doc

import (
	"io"
)

// childElemStruct describes elements that don't contain other elements.
// See https://html.spec.whatwg.org/multipage/dom.html#phrasing-content-2.
type childElemStruct struct {
	name       string
	attributes []IAttribute
}

// parentElemStruct describes a parent element.
type parentElemStruct struct {
	*childElemStruct
	children IContent
}

// func childElem creates a child element.
func childElem(name string, attrs ...IAttribute) *childElemStruct {
	return &childElemStruct{
		name:       name,
		attributes: attrs,
	}
}

// func ChildElem creates a child element.
func ChildElem(name string, attrs ...IAttribute) IContent {
	return childElem(name, attrs...)
}

// func ParentElem creates a ParentElemStruct.
func ParentElem(name string, attrs ...IAttribute) ContContainerFunc {
	return func(elems ...IContent) IContent {
		return &parentElemStruct{
			childElemStruct: childElem(name, attrs...),
			children:        ContContainer(elems...),
		}
	}
}

// RenderContent renders the element.
func (t *childElemStruct) RenderConent(wr io.Writer) (err error) {
	if err = writeByteSlices(wr, []byte(`<`), []byte(t.name)); err != nil {
		return
	}
	for _, attr := range t.attributes {
		if err = writeByteSlices(wr, []byte(` `)); err != nil {
			return
		}
		if err = attr.RenderAttr(wr); err != nil {
			return
		}
	}
	err = writeByteSlices(wr, []byte(`>`))
	return
}

// RenderContent renders the element.
func (t *parentElemStruct) RenderConent(wr io.Writer) (err error) {
	if err = t.childElemStruct.RenderConent(wr); err != nil {
		return
	}
	if err = t.children.RenderConent(wr); err != nil {
		return err
	}
	err = writeByteSlices(wr, []byte(`</`), []byte(t.name), []byte(`>`))
	return
}
