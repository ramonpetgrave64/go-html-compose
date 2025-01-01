package attr

import (
	"fmt"
	"go-html-compose/util"
	"html"
	"io"
)

type AttributeStruct struct {
	// render.Renderable
	Name  string
	Value string
}

func (a *AttributeStruct) Render(wr io.Writer) error {
	if _, err := wr.Write([]byte(fmt.Sprintf(`%s="%s"`, a.Name, html.EscapeString(a.Value)))); err != nil {
		return err
	}
	return nil
}

func (a *AttributeStruct) StructuredRender(wr io.Writer, tabs int) error {
	if _, err := wr.Write(util.GetTabBytes(tabs)); err != nil {
		return err
	}
	if _, err := wr.Write([]byte(fmt.Sprintf(`%s="%s"`, a.Name, html.EscapeString(a.Value)))); err != nil {
		return err
	}
	return nil
}

func Attr(name string, value string) *AttributeStruct {
	return &AttributeStruct{
		Name:  name,
		Value: value,
	}
}
