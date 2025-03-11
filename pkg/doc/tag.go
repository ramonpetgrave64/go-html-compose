package doc

import (
	"io"
)

type ContentFunc func(elems ...Renderable) *ParentTagStruct

type UnitTagStruct struct {
	Name       string
	Attributes []*AttributeStruct
	skipRender bool
}

func (t *UnitTagStruct) Render(wr io.Writer) (err error) {
	if t.skipRender {
		return
	}
	if err = WriteByteSlices(wr, []byte(`<`), []byte(t.Name)); err != nil {
		return
	}
	for _, attr := range t.Attributes {
		if err = WriteByteSlices(wr, []byte(` `)); err != nil {
			return
		}
		if err = attr.RenderTag(wr); err != nil {
			return
		}
	}
	err = WriteByteSlices(wr, []byte(`>`))
	return
}

type ParentTagStruct struct {
	*UnitTagStruct
	Container  *ContainerStruct
	skipRender bool
}

func (t *ParentTagStruct) Render(wr io.Writer) (err error) {
	if t.skipRender {
		return
	}
	if err = t.UnitTagStruct.Render(wr); err != nil {
		return
	}
	if err = t.Container.Render(wr); err != nil {
		return err
	}
	err = WriteByteSlices(wr, []byte(`</`), []byte(t.Name), []byte(`>`))
	return
}

func UnitTag(name string, attrs ...*AttributeStruct) *UnitTagStruct {
	return &UnitTagStruct{
		Name:       name,
		Attributes: attrs,
	}
}

func ParentTag(name string, attrs ...*AttributeStruct) ContentFunc {
	return func(elems ...Renderable) *ParentTagStruct {
		return &ParentTagStruct{
			UnitTagStruct: UnitTag(name, attrs...),
			Container:     Container(elems...),
		}
	}
}
