package doc

import (
	"io"
)

type IContent interface {
	RenderConent(wr io.Writer) (err error)
}
