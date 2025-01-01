package render

import (
	"bytes"
	"io"
)

type Renderable interface {
	Render(wr io.Writer) (err error)
	// RenderWithDelim(wr io.Writer, delim []byte)
	StructuredRender(wr io.Writer, tabs int) (err error)
}

func String(rendr Renderable) string {
	var buffer bytes.Buffer
	rendr.Render(&buffer)
	return buffer.String()
}

func Render(wr io.Writer, rendr Renderable) error {
	if err := rendr.Render(wr); err != nil {
		return err
	}
	return nil
}

// StructuredRender calls StructuredRender(wr, 0)
func StructuredRender(wr io.Writer, rendr Renderable) error {
	if err := rendr.StructuredRender(wr, 0); err != nil {
		return err
	}
	return nil
}
