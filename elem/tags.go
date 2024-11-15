package elem

import "go-html-compose/attr"

func Span(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("span", attrs...)
}

func Div(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("div", attrs...)
}

func Img(attrs ...*attr.AttributeStruct) *UnitTagStruct {
	return UnitTag("img", attrs...)
}

func HTML(attrs ...*attr.AttributeStruct) ContentFunc {
	return ParentTag("html", attrs...)
}
