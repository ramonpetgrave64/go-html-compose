package doc

import (
	"go-html-compose/render"
	"go-html-compose/util"
	"io"
)

type DocumentStruct struct {
	// render.Renderable
	Children []render.Renderable
}

func (c *DocumentStruct) Render(wr io.Writer) error {
	for _, elem := range c.Children {
		if err := elem.Render(wr); err != nil {
			return err
		}
	}
	return nil
}

func (c *DocumentStruct) StructuredRender(wr io.Writer, tabs int) error {
	for _, elem := range c.Children {
		if tabs > -1 {
			if _, err := wr.Write(util.NewlineContent); err != nil {
				return err
			}
		}
		if err := elem.StructuredRender(wr, tabs); err != nil {
			return err
		}
	}
	return nil
}

func Document(children ...render.Renderable) *DocumentStruct {
	return &DocumentStruct{
		Children: children,
	}
}
