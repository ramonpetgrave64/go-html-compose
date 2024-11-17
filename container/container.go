package container

import (
	"go-html-compose/render"
	"go-html-compose/util"
	"io"
)

type ContainerStruct struct {
	render.Renderable
	Children []render.Renderable
}

func (c *ContainerStruct) Render(wr io.Writer) {
	for _, elem := range c.Children {
		elem.Render(wr)
	}
}

func (c *ContainerStruct) StructuredRender(wr io.Writer, tabs int) {
	for _, elem := range c.Children {
		if tabs > -1 {
			wr.Write(util.NewlineContent)
		}
		elem.StructuredRender(wr, tabs)
	}
}

func Container(children ...render.Renderable) *ContainerStruct {
	return &ContainerStruct{
		Children: children,
	}
}
