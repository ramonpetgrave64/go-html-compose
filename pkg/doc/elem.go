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
	children *ContContainerStruct
}

// func ChildElem creates a ChildElemStruct.
func ChildElem(name string, attrs ...IAttribute) *childElemStruct {
	return &childElemStruct{
		name:       name,
		attributes: attrs,
	}
}

// func ParentElem creates a ParentElemStruct.
func ParentElem(name string, attrs ...IAttribute) ParentElemFunc {
	return func(elems ...IContent) IContent {
		return &parentElemStruct{
			childElemStruct: ChildElem(name, attrs...),
			children:        ContContainer(elems...),
		}
	}
}

// RenderContent renders the element.
func (t *childElemStruct) RenderConent(wr io.Writer) (err error) {
	if err = WriteByteSlices(wr, []byte(`<`), []byte(t.name)); err != nil {
		return
	}
	for _, attr := range t.attributes {
		if err = WriteByteSlices(wr, []byte(` `)); err != nil {
			return
		}
		if err = attr.RenderAttr(wr); err != nil {
			return
		}
	}
	err = WriteByteSlices(wr, []byte(`>`))
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
	err = WriteByteSlices(wr, []byte(`</`), []byte(t.name), []byte(`>`))
	return
}
