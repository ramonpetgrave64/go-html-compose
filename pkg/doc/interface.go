package doc

import "io"

// ContContainerFunc is a function that returns a content container.
type ContContainerFunc func(children ...IContent) IContent

type IContent interface {
	RenderConent(wr io.Writer) (err error)
}

type IAttribute interface {
	RenderAttr(wr io.Writer) (err error)
}
