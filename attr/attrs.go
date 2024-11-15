package attr

func Class(value string) *AttributeStruct {
	return Attr("class", value)
}

func Style(value string) *AttributeStruct {
	return Attr("style", value)
}

func Src(value string) *AttributeStruct {
	return Attr("src", value)
}
