package doc

import (
	"io"
)

type ContainerStruct struct {
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

func Container(children ...Renderable) *ContainerStruct {
	return &ContainerStruct{
		Children: children,
	}
}
