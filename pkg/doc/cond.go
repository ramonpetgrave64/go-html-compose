package doc

var (
	nilAttribute = BooleanAttr("", false)
)

// IfAttr conditonally renders either attribute.
func IfElseAttr(condition bool, ifTrue, ifFalse *AttrStruct) *AttrStruct {
	if condition {
		return ifTrue
	}
	return ifFalse
}

// IfAttr conditonally renders the attribute.
func IfAttr(condition bool, ifTrue *AttrStruct) *AttrStruct {
	return IfElseAttr(condition, ifTrue, nilAttribute)
}

// IfCont conditionally renders cont.
func IfCont(cond bool, ifTrue IContent) IContent {
	return IfElseCont(cond, ifTrue, RawText([]byte(``)))
}

// If conditionally renders either element.
func IfElseCont(cond bool, ifTrue, ifFalse IContent) IContent {
	if cond {
		return ifTrue
	}
	return ifFalse
}

// MapToContContainer maps the slice to a Renderables into a FlowStruct.
func MapToContContainer[T any](items []T, mapFunc func(T) IContent) *ContContainerStruct {
	rendrs := make([]IContent, len(items))
	for idx, item := range items {
		rendrs[idx] = mapFunc(item)
	}
	return ContContainer(rendrs...)
}
