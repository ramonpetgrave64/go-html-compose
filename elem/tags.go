package elem

import "go-html-compose/attr"

var (
	Doctype = UnitTag("!DOCTYPE html")
)

func A(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("a", attrs...)
}

func Article(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("article", attrs...)
}

func Body(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("body", attrs...)
}

func Button(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("button", attrs...)
}

func Br(attrs ...*attr.AttributeStruct) *UnitTagStruct {
	return UnitTag("br", attrs...)
}

func Details(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("details", attrs...)
}

func Div(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("div", attrs...)
}

func Footer(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("footer", attrs...)
}

func Head(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("head", attrs...)
}

func Header(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("header", attrs...)
}

func HTML(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("html", attrs...)
}

func I(attrs ...*attr.AttributeStruct) *UnitTagStruct {
	return UnitTag("i", attrs...)
}

func Main(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("main", attrs...)
}

func Meta(attrs ...*attr.AttributeStruct) *UnitTagStruct {
	return UnitTag("meta", attrs...)
}

func Img(attrs ...*attr.AttributeStruct) *UnitTagStruct {
	return UnitTag("img", attrs...)
}

func Li(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("li", attrs...)
}

func Link(attrs ...*attr.AttributeStruct) *UnitTagStruct {
	return UnitTag("link", attrs...)
}

func Nav(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("nav", attrs...)
}

func Section(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("section", attrs...)
}

func Span(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("span", attrs...)
}

func Style(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("style", attrs...)
}

func Summary(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("summary", attrs...)
}

func Title(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("title", attrs...)
}

func Ul(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("ul", attrs...)
}
