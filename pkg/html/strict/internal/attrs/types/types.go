package types

import (
	"io"

	"github.com/ramonpetgrave64/go-html-compose/pkg/doc"
	"github.com/ramonpetgrave64/go-html-compose/pkg/html/strict/elems"
)

type Src struct {
	doc.IAttribute
	elems.AudioAttribute
	elems.EmbedAttribute
	elems.IframeAttribute
	elems.ImgAttribute
	elems.InputAttribute
	elems.ScriptAttribute
	elems.SourceAttribute
	elems.TrackAttribute
	elems.VideoAttribute
}

func (a Src) RenderAttr(wr io.Writer) (err error) {
	return a.IAttribute.RenderAttr(wr)
}

type Value struct {
	doc.IAttribute
	elems.ButtonAttribute
	elems.OptionAttribute
	elems.DataAttribute
	elems.InputAttribute
	elems.LiAttribute
	elems.MeterAttribute
	elems.ProgressAttribute
}

func (a Value) RenderAttr(wr io.Writer) (err error) {
	return a.IAttribute.RenderAttr(wr)
}
