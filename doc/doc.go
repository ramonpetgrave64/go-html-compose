package doc

import (
	"go-html-compose/render"
	"io"
)

type ContainerStruct struct {
	// render.Renderable
	Children []render.Renderable
}

func (c *ContainerStruct) Render(wr io.Writer) error {
	for _, elem := range c.Children {
		if err := elem.Render(wr); err != nil {
			return err
		}
	}
	return nil
}

func (c *ContainerStruct) StructuredRender(wr io.Writer, tabs int) error {
	for _, elem := range c.Children {
		if tabs > -1 {
			if _, err := wr.Write(render.NewlineContent); err != nil {
				return err
			}
		}
		if err := elem.StructuredRender(wr, tabs); err != nil {
			return err
		}
	}
	return nil
}

func Container(children ...render.Renderable) *ContainerStruct {
	return &ContainerStruct{
		Children: children,
	}
}
