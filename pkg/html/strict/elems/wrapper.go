package elems

import (
	"io"

	"github.com/ramonpetgrave64/go-html-compose/pkg/doc"
)

type contentWrapper struct {
	doc.IContent
	UlChild
}

func (contentWrapper *contentWrapper) RenderConent(wr io.Writer) (err error) {
	return contentWrapper.IContent.RenderConent(wr)
}

func newContentWrapper(cont doc.IContent) *contentWrapper {
	return &contentWrapper{IContent: cont}
}
