package doc

import (
	"io"
)

type ContainerStruct struct {
	// render.Renderable
	Children []Renderable
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
			if _, err := wr.Write(NewlineContent); err != nil {
				return err
			}
		}
		if err := elem.StructuredRender(wr, tabs); err != nil {
			return err
		}
	}
	return nil
}

func Container(children ...Renderable) *ContainerStruct {
	return &ContainerStruct{
		Children: children,
	}
}
