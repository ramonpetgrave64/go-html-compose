package elem

import "go-html-compose/attr"

var (
	Doctype = UnitTag([]byte(`!DOCTYPE html"`))
)

func A(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`a`), attrs...)
}

func Article(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`article`), attrs...)
}

func Body(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`body`), attrs...)
}

func Button(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`button`), attrs...)
}

func Br(attrs ...*attr.AttributeStruct) *UnitTagStruct {
	return UnitTag([]byte(`br`), attrs...)
}

func Details(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`details`), attrs...)
}

func Div(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`div`), attrs...)
}

func Footer(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`footer`), attrs...)
}

func Head(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`head`), attrs...)
}

func Header(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`header`), attrs...)
}

func HTML(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`html`), attrs...)
}

func I(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`i`), attrs...)
}

func Main(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`main`), attrs...)
}

func Meta(attrs ...*attr.AttributeStruct) *UnitTagStruct {
	return UnitTag([]byte(`meta`), attrs...)
}

func Img(attrs ...*attr.AttributeStruct) *UnitTagStruct {
	return UnitTag([]byte(`img`), attrs...)
}

func Li(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`li`), attrs...)
}

func Link(attrs ...*attr.AttributeStruct) *UnitTagStruct {
	return UnitTag([]byte(`link`), attrs...)
}

func Nav(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`nav`), attrs...)
}

func Section(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`section`), attrs...)
}

func Span(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`span`), attrs...)
}

func Style(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`style`), attrs...)
}

func Summary(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`summary`), attrs...)
}

func Title(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`title`), attrs...)
}

func Ul(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag([]byte(`ul`), attrs...)
}
