package elem

import (
	"io"

	"go-html-compose/pkg/attr"
	"go-html-compose/pkg/doc"
)

type ContentFunc func(elems ...doc.Renderable) *ParentTagStruct

type UnitTagStruct struct {
	Name       string
	Attributes []*attr.AttributeStruct
}

func (t *UnitTagStruct) Render(wr io.Writer) (err error) {
	if err = doc.WriteByteSlices(wr, []byte(`<`), []byte(t.Name)); err != nil {
		return
	}
	for _, attr := range t.Attributes {
		if err = doc.WriteByteSlices(wr, doc.SpaceContent); err != nil {
			return
		}
		if err = attr.Render(wr); err != nil {
			return
		}
	}
	err = doc.WriteByteSlices(wr, []byte(`>`))
	return
}

type ParentTagStruct struct {
	*UnitTagStruct
	Container *doc.ContainerStruct
}

func (t *ParentTagStruct) Render(wr io.Writer) (err error) {
	if err = t.UnitTagStruct.Render(wr); err != nil {
		return
	}
	if err = t.Container.Render(wr); err != nil {
		return err
	}
	err = doc.WriteByteSlices(wr, []byte(`</`), []byte(t.Name), []byte(`>`))
	return
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
