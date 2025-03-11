package doc

import (
	"html"
	"io"
)

// AttributeStruct describes an HTML attribute.
type AttributeStruct struct {
	Name       string
	Value      string
	raw        bool
	skipRender bool
}

// Render renders the attribute to the io.Writer.
func (a *AttributeStruct) Render(wr io.Writer) error {
	if a.skipRender {
		return nil
	}
	value := a.Value
	if !a.raw {
		value = html.EscapeString(a.Value)
	}
	var err error
	if err = WriteByteSlices(
		wr,
		[]byte(a.Name), []byte(`="`), []byte(value), []byte(`"`),
	); err != nil {
		return err
	}
	return nil
}

// Attr creates an AttributeStruct.
func Attr(name string, value string) *AttributeStruct {
	return &AttributeStruct{
		Name:  name,
		Value: value,
		raw:   false,
	}
}

// RawAttr creates a attribute that renderes without escaping the value.
func RawAttr(name string, value string) *AttributeStruct {
	return &AttributeStruct{
		Name:  name,
		Value: value,
		raw:   true,
	}
}

// BooleanAttr creates an attribute that holds boolean values and conditionally renders if the cond is true.
func BooleanAttr(name string, cond bool) *AttributeStruct {
	value := name
	attr := Attr(name, value)
	return IfAttr(attr, cond)
}

// IfAttr makes an existing Attr conditionally renderable.
func IfAttr(attr *AttributeStruct, cond bool) *AttributeStruct {
	attr.skipRender = !cond
	return attr
}
