package elem

import (
	"go-html-compose/attr"
	"go-html-compose/doc"
	"go-html-compose/render"
	"io"
)

var (
	openBracket      = []byte(`<`)
	openBracketSlash = []byte(`</`)
	closeBracket     = []byte(`>`)
)

type ContentFunc func(elems ...render.Renderable) *ParentTagStruct

func writeOpeningTag(wr io.Writer, name []byte) error {
	if err := render.WriteByteSlices(wr, openBracket, name); err != nil {
		return err
	}
	return nil
}

func writeClosingTag(wr io.Writer, name []byte) error {
	if err := render.WriteByteSlices(wr, openBracketSlash, name, closeBracket); err != nil {
		return err
	}
	return nil
}

type UnitTagStruct struct {
	// render.Renderable
	Name       []byte
	Attributes []*attr.AttributeStruct
}

func (t *UnitTagStruct) Render(wr io.Writer) error {
	var err error
	if err = writeOpeningTag(wr, t.Name); err != nil {
		return err
	}
	for _, attr := range t.Attributes {
		wr.Write(render.SpaceContent)
		attr.Render(wr)
	}
	if err = render.WriteByteSlices(wr, closeBracket); err != nil {
		return err
	}
	return nil
}

func (t *UnitTagStruct) StructuredRender(wr io.Writer, tabs int) error {
	var err error
	if err = render.WriteTabBytes(wr, tabs); err != nil {
		return err
	}
	if err = writeOpeningTag(wr, t.Name); err != nil {
		return err
	}
	for idx, attr := range t.Attributes {
		if idx == 0 {
			if _, err = wr.Write(render.NewlineContent); err != nil {
				return err
			}
		}
		attr.StructuredRender(wr, tabs+1)
		if _, err := wr.Write(render.NewlineContent); err != nil {
			return err
		}
	}
	if len(t.Attributes) > 0 {
		if err = render.WriteTabBytes(wr, tabs); err != nil {
			return err
		}
	}
	if _, err = wr.Write(closeBracket); err != nil {
		return err
	}
	return nil
}

type ParentTagStruct struct {
	// render.Renderable
	*UnitTagStruct
	Container *doc.ContainerStruct
}

func (t *ParentTagStruct) Render(wr io.Writer) error {
	if err := t.UnitTagStruct.Render(wr); err != nil {
		return err
	}
	if err := t.Container.Render(wr); err != nil {
		return err
	}
	if err := writeClosingTag(wr, t.Name); err != nil {
		return err
	}
	return nil
}

func (t *ParentTagStruct) StructuredRender(wr io.Writer, tabs int) error {
	var err error
	if err = t.UnitTagStruct.StructuredRender(wr, tabs); err != nil {
		return err
	}
	if err := t.Container.StructuredRender(wr, tabs+1); err != nil {
		return err
	}
	if _, err = wr.Write(render.NewlineContent); err != nil {
		return err
	}
	if err = render.WriteTabBytes(wr, tabs); err != nil {
		return err
	}
	if err = writeClosingTag(wr, t.Name); err != nil {
		return err
	}
	return nil
}

func UnitTag(name []byte, attrs ...*attr.AttributeStruct) *UnitTagStruct {
	return &UnitTagStruct{
		Name:       name,
		Attributes: attrs,
	}
}

func ParentTag(name []byte, attrs ...*attr.AttributeStruct) ContentFunc {
	return func(elems ...render.Renderable) *ParentTagStruct {
		return &ParentTagStruct{
			UnitTagStruct: UnitTag(name, attrs...),
			Container:     doc.Container(elems...),
		}
	}
}
