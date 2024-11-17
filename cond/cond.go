package cond

import (
	"go-html-compose/render"
	"go-html-compose/text"
)

var nilRenderable = text.RawText("")

func If(cond bool, rendr render.Renderable) render.Renderable {
	if cond {
		return rendr
	}
	return nilRenderable
}

func Map[T any](items []T, mapFunc func(T) render.Renderable) []render.Renderable {
	rendrs := make([]render.Renderable, len(items))
	for idx, item := range items {
		rendrs[idx] = mapFunc(item)
	}
	return rendrs
}
