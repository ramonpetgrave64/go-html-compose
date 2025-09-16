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

type contentFunc func(attrs ...doc.IAttribute) doc.ContContainerFunc

func newContentWrapper2[
	A doc.IAttribute,
	C doc.IContent,
](contentFunc contentFunc, attrs []A, children []C) *contentWrapper {
	convertedAttributes := toIAttributes(attrs)
	convertedChildren := toIContent(children)
	content := contentFunc(convertedAttributes...)(convertedChildren...)
	return &contentWrapper{IContent: content}
}

func newContentWrapper3[P doc.IContent, A doc.IAttribute, C doc.IContent](
	contentFunc contentFunc, attrs []A,
) TypedContContainerFunc[P, C] {
	return func(children ...C) P {
		return newContentWrapper2(contentFunc, attrs, children).IContent.(P)
	}
}
