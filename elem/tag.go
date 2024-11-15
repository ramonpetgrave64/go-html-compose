package elem

import (
	"fmt"
	"go-html-compose/attr"
	"go-html-compose/render"
	"go-html-compose/util"
	"io"
)

type ContentFunc func(elems ...render.Renderable) *ParentTagStruct

type UnitTagStruct struct {
	render.Renderable
	Name       string
	Attributes []*attr.AttributeStruct
}

func (t *UnitTagStruct) Render(wr io.Writer) {
	wr.Write([]byte(fmt.Sprintf(`<%s`, t.Name)))
	for _, attr := range t.Attributes {
		wr.Write([]byte(` `))
		attr.Render(wr)
	}
	wr.Write([]byte(`>`))
}

func (t *UnitTagStruct) StructuredRenderWithTabs(wr io.Writer, tabs int) {
	spaces := util.GetTabBytes(tabs)
	wr.Write(spaces)
	wr.Write([]byte(fmt.Sprintf(`<%s`, t.Name)))
	for _, attr := range t.Attributes {
		wr.Write([]byte(` `))
		attr.Render(wr)
	}
	wr.Write([]byte(`>`))
}

type ParentTagStruct struct {
	render.Renderable
	*UnitTagStruct
	Children []render.Renderable
}

func (t *ParentTagStruct) Render(wr io.Writer) {
	t.UnitTagStruct.Render(wr)
	for _, elem := range t.Children {
		elem.Render(wr)
	}
	wr.Write([]byte(fmt.Sprintf(`</%s>`, t.Name)))
}

func (t *ParentTagStruct) StructuredRenderWithTabs(wr io.Writer, tabs int) {
	t.UnitTagStruct.StructuredRenderWithTabs(wr, tabs)
	for _, elem := range t.Children {
		wr.Write(util.NewlineContent)
		elem.StructuredRenderWithTabs(wr, tabs+1)
	}
	wr.Write(util.NewlineContent)
	wr.Write(util.GetTabBytes(tabs))
	wr.Write([]byte(fmt.Sprintf(`</%s>`, t.Name)))
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
			Children:      elems,
		}
	}
}
