package attrs

import (
	"io"

	"github.com/ramonpetgrave64/go-html-compose/pkg/doc"
	"github.com/ramonpetgrave64/go-html-compose/pkg/html/strict/elems"
)

type attrWrapper struct {
	doc.IAttribute
	elems.GlobalAttribute // GlobalAttribute implements all other element attribute types
}

func (attrWrapper *attrWrapper) RenderAttr(wr io.Writer) (err error) {
	return attrWrapper.IAttribute.RenderAttr(wr)
}

func newAttrWrapper(attr doc.IAttribute) *attrWrapper {
	return &attrWrapper{IAttribute: attr}
}
