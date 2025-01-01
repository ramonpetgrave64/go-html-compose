package cond

import (
	"go-html-compose/doc"
	"go-html-compose/render"
	"go-html-compose/text"
)

var nilRenderable = text.RawText("")

func If(cond bool, rendr render.Renderable) render.Renderable {

	return IfElse(cond, rendr, nilRenderable)
}

func IfElse(cond bool, rendrIfTrue render.Renderable, renderIfFalse render.Renderable) render.Renderable {
	if cond {
		return rendrIfTrue
	}
	return renderIfFalse
}

func Map[T any](items []T, mapFunc func(T) render.Renderable) []render.Renderable {
	rendrs := make([]render.Renderable, len(items))
	for idx, item := range items {
		rendrs[idx] = mapFunc(item)
	}
	return rendrs
}

func MapToContainer[T any](items []T, mapFunc func(T) render.Renderable) *doc.ContainerStruct {
	rendrs := make([]render.Renderable, len(items))
	for idx, item := range items {
		rendrs[idx] = mapFunc(item)
	}
	return doc.Container(rendrs...)
}
