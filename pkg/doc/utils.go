package doc

import (
	"io"
)

func writeByteSlices(wr io.Writer, slices ...[]byte) error {
	var err error
	for _, slice := range slices {
		if _, err = wr.Write(slice); err != nil {
			return err
		}
	}
	return nil
}
