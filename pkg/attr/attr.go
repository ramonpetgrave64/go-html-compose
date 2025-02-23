package attr

import (
	"html"
	"io"

	"go-html-compose/pkg/doc"
)

// AttributeStruct describes an HTML attribute.
type AttributeStruct struct {
	Name       string
	Value      *string
	skipRender bool
}

// Render renders the attribute to the io.Writer.
func (a *AttributeStruct) Render(wr io.Writer) error {
	if a.skipRender {
		return nil
	}
	var err error
	if err = doc.WriteByteSlices(
		wr,
		[]byte(a.Name), []byte(`="`), []byte(html.EscapeString(*a.Value)), []byte(`"`),
	); err != nil {
		return err
	}
	return nil
}

// Attr creates an AttributeStruct.
func Attr(name string, value *string) *AttributeStruct {
	return &AttributeStruct{
		Name:  name,
		Value: value,
	}
}

// BooleanAttr creates an attribute that holds boolean values and conditionally renders if the boolean is true.
func BooleanAttr(name string, boolean bool) *AttributeStruct {
	value := name
	attr := Attr(name, &value)
	attr.skipRender = !boolean
	return attr
}
