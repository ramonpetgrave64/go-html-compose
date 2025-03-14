package doc

import "io"

// ParentElemFunc is a function that returns a *ParentElemStruct.
type ParentElemFunc func(content ...IContent) IContent

type IContent interface {
	RenderConent(wr io.Writer) (err error)
}

type IAttribute interface {
	RenderAttr(wr io.Writer) (err error)
}
