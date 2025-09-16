package elems

import (
	"io"

	"github.com/ramonpetgrave64/go-html-compose/pkg/doc"
)

type ContContainerFunc[Parent, Child doc.IContent] func(children ...Child) Parent

type FlowContentType interface {
	doc.IContent
	isFlowContent()
}

type UlType struct {
	doc.IContent
	FlowContentType
}

func (e UlType) RenderConent(wr io.Writer) (err error) {
	return e.IContent.RenderConent(wr)
}

type UlChild interface {
	doc.IContent
	isUlChild()
}

type LiType struct {
	doc.IContent
	UlChild
}

func (e LiType) RenderConent(wr io.Writer) (err error) {
	return e.IContent.RenderConent(wr)
}

type LiChild interface {
	doc.IContent
	isLiChild()
}

type ScriptType struct {
	doc.IContent
}

func (e ScriptType) RenderConent(wr io.Writer) (err error) {
	return e.IContent.RenderConent(wr)
}

type ScriptChild interface {
	doc.IContent
	isScriptChild()
}
