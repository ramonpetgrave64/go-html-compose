package render

import (
	"bytes"
	"io"
)

type Renderable interface {
	Render(wr io.Writer)
	StructuredRenderWithTabs(wr io.Writer, tabs int)
}

func String(rendr Renderable) string {
	var buffer bytes.Buffer
	rendr.Render(&buffer)
	return buffer.String()
}

func Render(wr io.Writer, rendr Renderable) {
	rendr.Render(wr)
}

func StructuredRender(wr io.Writer, rendr Renderable) {
	rendr.StructuredRenderWithTabs(wr, 0)
}
