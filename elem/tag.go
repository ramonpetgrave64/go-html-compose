package elem

import (
	"go-html-compose/attr"
	"go-html-compose/doc"
	"io"
)

var (
	openBracket      = []byte(`<`)
	openBracketSlash = []byte(`</`)
	closeBracket     = []byte(`>`)
)

type ContentFunc func(elems ...doc.Renderable) *ParentTagStruct

func writeOpeningTag(wr io.Writer, name string) error {
	if err := doc.WriteByteSlices(wr, openBracket, []byte(name)); err != nil {
		return err
	}
	return nil
}

func writeClosingTag(wr io.Writer, name string) error {
	if err := doc.WriteByteSlices(wr, openBracketSlash, []byte(name), closeBracket); err != nil {
		return err
	}
	return nil
}

type UnitTagStruct struct {
	// doc.Renderable
	Name       string
	Attributes []*attr.AttributeStruct
}

func (t *UnitTagStruct) Render(wr io.Writer) error {
	var err error
	if err = writeOpeningTag(wr, t.Name); err != nil {
		return err
	}
	for _, attr := range t.Attributes {
		wr.Write(doc.SpaceContent)
		attr.Render(wr)
	}
	if _, err = wr.Write(closeBracket); err != nil {
		return err
	}
	return nil
}

func (t *UnitTagStruct) StructuredRender(wr io.Writer, tabs int) error {
	var err error
	if err = doc.WriteTabBytes(wr, tabs); err != nil {
		return err
	}
	if err = writeOpeningTag(wr, t.Name); err != nil {
		return err
	}
	for idx, attr := range t.Attributes {
		if idx == 0 {
			if _, err = wr.Write(doc.NewlineContent); err != nil {
				return err
			}
		}
		attr.StructuredRender(wr, tabs+1)
		if _, err := wr.Write(doc.NewlineContent); err != nil {
			return err
		}
	}
	if len(t.Attributes) > 0 {
		if err = doc.WriteTabBytes(wr, tabs); err != nil {
			return err
		}
	}
	if _, err = wr.Write(closeBracket); err != nil {
		return err
	}
	return nil
}

type ParentTagStruct struct {
	// doc.Renderable
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
	if _, err = wr.Write(doc.NewlineContent); err != nil {
		return err
	}
	if err = doc.WriteTabBytes(wr, tabs); err != nil {
		return err
	}
	if err = writeClosingTag(wr, t.Name); err != nil {
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
	return func(elems ...doc.Renderable) *ParentTagStruct {
		return &ParentTagStruct{
			UnitTagStruct: UnitTag(name, attrs...),
			Container:     doc.Container(elems...),
		}
	}
}
