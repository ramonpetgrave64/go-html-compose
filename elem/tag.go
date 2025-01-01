package elem

import (
	"fmt"
	"go-html-compose/attr"
	"go-html-compose/doc"
	"go-html-compose/render"
	"go-html-compose/util"
	"io"
)

type ContentFunc func(elems ...render.Renderable) *ParentTagStruct

func tagOpening(name string) []byte {
	return []byte(fmt.Sprintf(`<%s`, name))
}

func closingTag(name string) []byte {
	return []byte(fmt.Sprintf(`</%s>`, name))
}

type UnitTagStruct struct {
	// render.Renderable
	Name       string
	Attributes []*attr.AttributeStruct
}

func (t *UnitTagStruct) Render(wr io.Writer) error {
	if _, err := wr.Write(tagOpening(t.Name)); err != nil {
		return err
	}
	for _, attr := range t.Attributes {
		wr.Write(util.SpaceContent)
		attr.Render(wr)
	}
	if _, err := wr.Write([]byte(`>`)); err != nil {
		return err
	}
	return nil
}

func (t *UnitTagStruct) StructuredRender(wr io.Writer, tabs int) error {
	if _, err := wr.Write(util.GetTabBytes(tabs)); err != nil {
		return err
	}
	if _, err := wr.Write(tagOpening(t.Name)); err != nil {
		return err
	}
	for idx, attr := range t.Attributes {
		if idx == 0 {
			if _, err := wr.Write(util.NewlineContent); err != nil {
				return err
			}
		}
		attr.StructuredRender(wr, tabs+1)
		if _, err := wr.Write(util.NewlineContent); err != nil {
			return err
		}
	}

	if len(t.Attributes) > 0 {
		if _, err := wr.Write(util.GetTabBytes(tabs)); err != nil {
			return err
		}
	}
	if _, err := wr.Write([]byte(`>`)); err != nil {
		return err
	}
	return nil
}

type ParentTagStruct struct {
	// render.Renderable
	*UnitTagStruct
	Document *doc.ContainerStruct
	Children []render.Renderable
}

func (t *ParentTagStruct) Render(wr io.Writer) error {
	t.UnitTagStruct.Render(wr)
	t.Document.Render(wr)
	if _, err := wr.Write(closingTag(t.Name)); err != nil {
		return err
	}
	return nil
}

func (t *ParentTagStruct) StructuredRender(wr io.Writer, tabs int) error {
	if err := t.UnitTagStruct.StructuredRender(wr, tabs); err != nil {
		return err
	}
	for _, elem := range t.Children {
		if _, err := wr.Write(util.NewlineContent); err != nil {
			return err
		}
		if err := elem.StructuredRender(wr, tabs+1); err != nil {
			return err
		}
	}
	if _, err := wr.Write(util.NewlineContent); err != nil {
		return err
	}
	if _, err := wr.Write(util.GetTabBytes(tabs)); err != nil {
		return err
	}
	if _, err := wr.Write([]byte(fmt.Sprintf(`</%s>`, t.Name))); err != nil {
		return err
	}
	return nil
}

func UnitTag(name string, attrs ...*attr.AttributeStruct) *UnitTagStruct {
	return &UnitTagStruct{
		Name:       name,
		Attributes: attrs,
	}
}

func ParentTag(name string, attrs ...*attr.AttributeStruct) ContentFunc {
	return func(elems ...render.Renderable) *ParentTagStruct {
		return &ParentTagStruct{
			UnitTagStruct: UnitTag(name, attrs...),
			Document:      doc.Container(elems...),
			Children:      elems,
		}
	}
}
