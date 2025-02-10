package doc

func If(cond bool, rendr Renderable) Renderable {
	return IfElse(cond, rendr, RawText([]byte(``)))
}

func IfElse(cond bool, rendrIfTrue Renderable, renderIfFalse Renderable) Renderable {
	if cond {
		return rendrIfTrue
	}
	return renderIfFalse
}

func Map[T any](items []T, mapFunc func(T) Renderable) []Renderable {
	rendrs := make([]Renderable, len(items))
	for idx, item := range items {
		rendrs[idx] = mapFunc(item)
	}
	return rendrs
}

func MapToContainer[T any](items []T, mapFunc func(T) Renderable) *ContainerStruct {
	rendrs := make([]Renderable, len(items))
	for idx, item := range items {
		rendrs[idx] = mapFunc(item)
	}
	return Container(rendrs...)
}
