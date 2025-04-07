package doc

import (
	"html"
	"io"
)

// attrStruct describes an HTML attribute.
type attrStruct struct {
	Name       string
	Value      string
	raw        bool
	skipRender bool
}

// RenderAttr renders the attribute to the io.Writer.
func (a *attrStruct) RenderAttr(wr io.Writer) error {
	if a.skipRender {
		return nil
	}
	value := a.Value
	if !a.raw {
		value = html.EscapeString(a.Value)
	}
	var err error
	// TODO: decide if a.Name should be escaped.
	if err = writeByteSlices(
		wr,
		[]byte(a.Name), []byte(`="`), []byte(value), []byte(`"`),
	); err != nil {
		return err
	}
	return nil
}

// Attr creates an AttributeStruct.
func Attr(name, value string) *attrStruct {
	return &attrStruct{
		Name:  name,
		Value: value,
		raw:   false,
	}
}

// RawAttr creates a attribute that renders without escaping the value.
func RawAttr(name, value string) *attrStruct {
	return &attrStruct{
		Name:  name,
		Value: value,
		raw:   true,
	}
}

// BooleanAttr creates an attribute that holds boolean values and conditionally renders if the cond is true.
func BooleanAttr(name string, cond bool) *attrStruct {
	value := name
	attr := Attr(name, value)
	attr.skipRender = !cond
	return attr
}
