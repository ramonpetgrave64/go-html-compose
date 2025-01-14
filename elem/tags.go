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

func Fieldset(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("fieldset", attrs...)
}

func Footer(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("footer", attrs...)
}

func Form(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("form", attrs...)
}

func H1(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("h1", attrs...)
}

func H2(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("h2", attrs...)
}

func H3(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("h3", attrs...)
}

func H4(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("h4", attrs...)
}

func H5(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("h5", attrs...)
}

func H6(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("h6", attrs...)
}

func Head(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("head", attrs...)
}

func Header(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("header", attrs...)
}

func Hr(attrs ...*attr.AttributeStruct) *UnitTagStruct {
	return UnitTag("hr", attrs...)
}

func HTML(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("html", attrs...)
}

func I(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("i", attrs...)
}

func Input(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("input", attrs...)
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

func Label(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("label", attrs...)
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

func Ol(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("ol", attrs...)
}

func Option(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("option", attrs...)
}

func P(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("p", attrs...)
}

func Script(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("script", attrs...)
}

func Section(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("section", attrs...)
}

func Select(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("select", attrs...)
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

func Table(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("table", attrs...)
}

func TBody(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("tbody", attrs...)
}

func TD(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("td", attrs...)
}

func TH(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("th", attrs...)
}

func THead(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("thead", attrs...)
}

func Title(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("title", attrs...)
}

func TR(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("tr", attrs...)
}

func Ul(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("ul", attrs...)
}
