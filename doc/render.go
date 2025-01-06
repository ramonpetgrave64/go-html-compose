package doc

import (
	"bytes"
	"io"
)

var (
	TabContent     = []byte(`	`) // tab character
	SpaceContent   = []byte(` `) // space character
	NewlineContent = []byte("\n")
)

type Renderable interface {
	Render(wr io.Writer) (err error)
	// RenderWithDelim(wr io.Writer, delim []byte)
	StructuredRender(wr io.Writer, tabs int) (err error)
}

func Bytes(rendr Renderable) ([]byte, error) {
	var buffer bytes.Buffer
	if err := rendr.Render(&buffer); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
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

func WriteTabBytes(wr io.Writer, tabs int) error {
	var err error
	for range tabs {
		if _, err = wr.Write(TabContent); err != nil {
			return err
		}
	}
	return nil
}

func WriteByteSlices(wr io.Writer, slices ...[]byte) error {
	var err error
	for _, slice := range slices {
		if _, err = wr.Write(slice); err != nil {
			return err
		}
	}
	return nil
}
