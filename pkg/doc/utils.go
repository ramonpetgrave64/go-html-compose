package doc

import (
	"bytes"
	"io"
)

func Bytes(rendr IContent) ([]byte, error) {
	var buffer bytes.Buffer
	if err := rendr.RenderConent(&buffer); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
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
