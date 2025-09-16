package elems

import "github.com/ramonpetgrave64/go-html-compose/pkg/doc"

type TypedContContainerFunc[Parent, Child doc.IContent] func(children ...Child) Parent

type UlChild interface {
	doc.IContent
	isUlChild()
}

type LiType interface {
	doc.IContent
	UlChild
}

type LiChild interface {
	doc.IContent
	isLiChild()
}

type ScriptChild interface {
	doc.IContent
	isScriptChild()
}
