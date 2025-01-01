package attr

import (
	"strconv"
)

func Charset(value string) *AttributeStruct {
	return Attr([]byte(`charset`), value)
}

func Class(value string) *AttributeStruct {
	return Attr([]byte(`class`), value)
}

func Content(value string) *AttributeStruct {
	return Attr([]byte(`content`), value)
}

func DataTheme(value string) *AttributeStruct {
	return Attr([]byte(`data-theme`), value)
}

func Href(value string) *AttributeStruct {
	return Attr([]byte(`href`), value)
}

func HTTPEquiv(value string) *AttributeStruct {
	return Attr([]byte(`http-equiv`), value)
}

func Lang(value string) *AttributeStruct {
	return Attr([]byte(`lang`), value)
}

func Name(value string) *AttributeStruct {
	return Attr([]byte(`name`), value)
}

func Open(val bool) *AttributeStruct {
	return Attr([]byte(`open`), strconv.FormatBool(val))
}

func Rel(value string) *AttributeStruct {
	return Attr([]byte(`rel`), value)
}

func Role(value string) *AttributeStruct {
	return Attr([]byte(`role`), value)
}

func Sizes(value string) *AttributeStruct {
	return Attr([]byte(`sizes`), value)
}

func Src(value string) *AttributeStruct {
	return Attr([]byte(`src`), value)
}

func Style(value string) *AttributeStruct {
	return Attr([]byte(`style`), value)
}

func Type(value string) *AttributeStruct {
	return Attr([]byte(`type`), value)
}
