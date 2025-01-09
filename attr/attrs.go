package attr

func Action(value string) *AttributeStruct {
	return Attr("action", &value)
}

func AriaLabel(value string) *AttributeStruct {
	return Attr("aria-label", &value)
}

func Charset(value string) *AttributeStruct {
	return Attr("charset", &value)
}

func Class(value string) *AttributeStruct {
	return Attr("class", &value)
}

func Content(value string) *AttributeStruct {
	return Attr("content", &value)
}

func DataTheme(value string) *AttributeStruct {
	return Attr("data-theme", &value)
}

func Href(value string) *AttributeStruct {
	return Attr("href", &value)
}

func HTTPEquiv(value string) *AttributeStruct {
	return Attr("http-equiv", &value)
}

func Lang(value string) *AttributeStruct {
	return Attr("lang", &value)
}

func Method(value string) *AttributeStruct {
	return Attr("method", &value)
}

func Name(value string) *AttributeStruct {
	return Attr("name", &value)
}

func Open(value bool) *AttributeStruct {
	return BooleanAttr("open", value)
}

func Placeholder(value string) *AttributeStruct {
	return Attr("placeholder", &value)
}

func ReadOnly(value bool) *AttributeStruct {
	return BooleanAttr("readonly", value)
}

func Rel(value string) *AttributeStruct {
	return Attr("rel", &value)
}

func Required(value bool) *AttributeStruct {
	return BooleanAttr("required", value)
}

func Role(value string) *AttributeStruct {
	return Attr("role", &value)
}

func Selected(value bool) *AttributeStruct {
	return BooleanAttr("selected", value)
}

func Scope(value string) *AttributeStruct {
	return Attr("scope", &value)
}

func Sizes(value string) *AttributeStruct {
	return Attr("sizes", &value)
}

func Src(value string) *AttributeStruct {
	return Attr("src", &value)
}

func Style(value string) *AttributeStruct {
	return Attr("style", &value)
}

func Type(value string) *AttributeStruct {
	return Attr("type", &value)
}

func Value(value string) *AttributeStruct {
	return Attr("value", &value)
}
