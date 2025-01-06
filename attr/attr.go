package attr

import (
	"go-html-compose/doc"
	"html"
	"io"
)

var (
	equalSign = []byte(`=`)
	quote     = []byte(`"`)
)

type AttributeStruct struct {
	// doc.Renderable
	Name       string
	Value      *string
	skipRender bool
}

func (a *AttributeStruct) Render(wr io.Writer) error {
	if a.skipRender {
		return nil
	}
	var err error
	if err = doc.WriteByteSlices(
		wr,
		[]byte(a.Name), equalSign, quote, []byte(html.EscapeString(*a.Value)), quote,
	); err != nil {
		return err
	}
	return nil
}

func (a *AttributeStruct) StructuredRender(wr io.Writer, tabs int) error {
	var err error
	if err = doc.WriteTabBytes(wr, tabs); err != nil {
		return err
	}
	if err = a.Render(wr); err != nil {
		return err
	}
	return nil
}

func Attr(name string, value *string) *AttributeStruct {
	return &AttributeStruct{
		Name:  name,
		Value: value,
	}
}

func BooleanAttr(name string, boolean bool) *AttributeStruct {
	value := name
	attr := Attr(name, &value)
	attr.skipRender = !boolean
	return attr
}
