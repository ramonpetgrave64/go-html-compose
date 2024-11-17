package render

import (
	"bytes"
	"io"
)

type Renderable interface {
	Render(wr io.Writer)
	// RenderWithDelim(wr io.Writer, delim []byte)
	StructuredRender(wr io.Writer, tabs int)
}

func String(rendr Renderable) string {
	var buffer bytes.Buffer
	rendr.Render(&buffer)
	return buffer.String()
}

func Render(wr io.Writer, rendr Renderable) {
	rendr.Render(wr)
}

// StructuredRender calls StructuredRender(wr, 0)
func StructuredRender(wr io.Writer, rendr Renderable) {
	rendr.StructuredRender(wr, 0)
}
