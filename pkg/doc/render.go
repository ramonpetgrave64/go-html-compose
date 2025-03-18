package doc

import (
	"bytes"
	"io"
)

type Renderable interface {
	Render(wr io.Writer) (err error)
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

func WriteByteSlices(wr io.Writer, slices ...[]byte) error {
	var err error
	for _, slice := range slices {
		if _, err = wr.Write(slice); err != nil {
			return err
		}
	}
	return nil
}
