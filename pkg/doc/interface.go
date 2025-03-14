package doc

import "io"

type IContent interface {
	RenderConent(wr io.Writer) (err error)
}

type IAttribute interface {
	RenderAttr(wr io.Writer) (err error)
}
