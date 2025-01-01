package attr

import (
	"fmt"
	"go-html-compose/render"
	"html"
	"io"
)

type AttributeStruct struct {
	// render.Renderable
	Name  []byte
	Value string
}

func (a *AttributeStruct) Render(wr io.Writer) error {
	if _, err := wr.Write([]byte(fmt.Sprintf(`%s="%s"`, a.Name, html.EscapeString(a.Value)))); err != nil {
		return err
	}
	return nil
}

func (a *AttributeStruct) StructuredRender(wr io.Writer, tabs int) error {
	var (
		equalSign = []byte(`=`)
		quote     = []byte(`"`)
		err       error
	)
	if err = render.WriteTabBytes(wr, tabs); err != nil {
		return err
	}
	if err = render.WriteByteSlices(
		wr,
		a.Name, equalSign, quote, []byte(html.EscapeString(a.Value)), quote,
	); err != nil {
		return err
	}
	return nil
}

func Attr(name []byte, value string) *AttributeStruct {
	return &AttributeStruct{
		Name:  name,
		Value: value,
	}
}
