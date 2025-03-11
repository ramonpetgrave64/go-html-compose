// Package doc has tools to construct HTML content.
// See https://html.spec.whatwg.org/multipage/dom.html#kinds-of-content.
package doc

import (
	"io"
)

// ChildElemStruct describes elements that don't contain other elements.
// See https://html.spec.whatwg.org/multipage/dom.html#phrasing-content-2.
type ChildElemStruct struct {
	Name       string
	Attributes []*AttributeStruct
}

// ParentElemStruct describes a parent element.
type ParentElemStruct struct {
	*ChildElemStruct
	Children *ContContainerStruct
}

// ParentElemFunc is a function that returns a *ParentElemStruct.
type ParentElemFunc func(content ...IContent) *ParentElemStruct

// func ChildElem creates a ChildElemStruct.
func ChildElem(name string, attrs ...*AttributeStruct) *ChildElemStruct {
	return &ChildElemStruct{
		Name:       name,
		Attributes: attrs,
	}
}

// func ParentElem creates a ParentElemStruct.
func ParentElem(name string, attrs ...*AttributeStruct) ParentElemFunc {
	return func(elems ...IContent) *ParentElemStruct {
		return &ParentElemStruct{
			ChildElemStruct: ChildElem(name, attrs...),
			Children:        ContContainer(elems...),
		}
	}
}

// RenderContent renders the element.
func (t *ChildElemStruct) RenderConent(wr io.Writer) (err error) {
	if err = WriteByteSlices(wr, []byte(`<`), []byte(t.Name)); err != nil {
		return
	}
	for _, attr := range t.Attributes {
		if err = WriteByteSlices(wr, []byte(` `)); err != nil {
			return
		}
		if err = attr.RenderTag(wr); err != nil {
			return
		}
	}
	err = WriteByteSlices(wr, []byte(`>`))
	return
}

// RenderContent renders the element.
func (t *ParentElemStruct) RenderConent(wr io.Writer) (err error) {
	if err = t.ChildElemStruct.RenderConent(wr); err != nil {
		return
	}
	if err = t.Children.RenderConent(wr); err != nil {
		return err
	}
	err = WriteByteSlices(wr, []byte(`</`), []byte(t.Name), []byte(`>`))
	return
}
