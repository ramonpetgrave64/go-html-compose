package elem

import (
	"fmt"
	"go-html-compose/attr"
	"go-html-compose/container"
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
	render.Renderable
	Name       string
	Attributes []*attr.AttributeStruct
}

func (t *UnitTagStruct) Render(wr io.Writer) {
	wr.Write(tagOpening(t.Name))
	for _, attr := range t.Attributes {
		wr.Write(util.SpaceContent)
		attr.Render(wr)
	}
	wr.Write([]byte(`>`))
}

func (t *UnitTagStruct) StructuredRender(wr io.Writer, tabs int) {
	wr.Write(util.GetTabBytes(tabs))
	wr.Write(tagOpening(t.Name))
	for idx, attr := range t.Attributes {
		if idx == 0 {
			wr.Write(util.NewlineContent)
		}
		attr.StructuredRender(wr, tabs+1)
		wr.Write(util.NewlineContent)
	}

	if len(t.Attributes) > 0 {
		wr.Write(util.GetTabBytes(tabs))
	}
	wr.Write([]byte(`>`))
}

type ParentTagStruct struct {
	render.Renderable
	*UnitTagStruct
	Container *container.ContainerStruct
	Children  []render.Renderable
}

func (t *ParentTagStruct) Render(wr io.Writer) {
	t.UnitTagStruct.Render(wr)
	t.Container.Render(wr)
	wr.Write(closingTag(t.Name))
}

func (t *ParentTagStruct) StructuredRender(wr io.Writer, tabs int) {
	t.UnitTagStruct.StructuredRender(wr, tabs)
	for _, elem := range t.Children {
		wr.Write(util.NewlineContent)
		elem.StructuredRender(wr, tabs+1)
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
			Container: &container.ContainerStruct{
				Children: elems,
			},
			Children: elems,
		}
	}
}
