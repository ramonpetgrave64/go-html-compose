package attr

import (
	"html"
	"io"

	"go-html-compose/pkg/doc"
)

type AttributeStruct struct {
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
		[]byte(a.Name), []byte(`="`), []byte(html.EscapeString(*a.Value)), []byte(`"`),
	); err != nil {
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
