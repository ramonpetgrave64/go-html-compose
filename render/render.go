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

func String(rendr Renderable) (text string, err error) {
	var buffer bytes.Buffer
	if err = rendr.Render(&buffer); err != nil {
		return
	}
	text = buffer.String()
	return
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
