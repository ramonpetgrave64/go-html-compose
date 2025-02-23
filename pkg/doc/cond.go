package doc

// If conditionally renders the Renderable.
func If(cond bool, rendr Renderable) Renderable {
	return IfElse(cond, rendr, RawText([]byte(``)))
}

// If conditionally renders either Renderable.
func IfElse(cond bool, rendrIfTrue, renderIfFalse Renderable) Renderable {
	if cond {
		return rendrIfTrue
	}
	return renderIfFalse
}

// MapToContainer maps the slice to a Renderables into a ContainerStruct.
func MapToContainer[T any](items []T, mapFunc func(T) Renderable) *ContainerStruct {
	rendrs := make([]Renderable, len(items))
	for idx, item := range items {
		rendrs[idx] = mapFunc(item)
	}
	return Container(rendrs...)
}
