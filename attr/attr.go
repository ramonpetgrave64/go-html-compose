package attr

import (
	"fmt"
	"go-html-compose/render"
	"go-html-compose/util"
	"html"
	"io"
)

type AttributeStruct struct {
	render.Renderable
	Name  string
	Value string
}

func (a *AttributeStruct) Render(wr io.Writer) {
	wr.Write([]byte(fmt.Sprintf(`%s="%s"`, a.Name, html.EscapeString(a.Value))))
}

func (a *AttributeStruct) StructuredRender(wr io.Writer, tabs int) {
	wr.Write(util.GetTabBytes(tabs))
	wr.Write([]byte(fmt.Sprintf(`%s="%s"`, a.Name, html.EscapeString(a.Value))))
}

func Attr(name string, value string) *AttributeStruct {
	return &AttributeStruct{
		Name:  name,
		Value: value,
	}
}
