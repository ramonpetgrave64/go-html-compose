package attrs

import (
	"io"

	"github.com/ramonpetgrave64/go-html-compose/pkg/doc"
)

type Abbr struct {
	doc.IAttribute
	ThAttribute
}

func (a Abbr) RenderAttr(wr io.Writer) (err error) {
	return a.IAttribute.RenderAttr(wr)
}

type Src struct {
	doc.IAttribute
	AudioAttribute
	EmbedAttribute
	IframeAttribute
	ImgAttribute
	InputAttribute
	ScriptAttribute
	SourceAttribute
	TrackAttribute
	VideoAttribute
}

func (a Src) RenderAttr(wr io.Writer) (err error) {
	return a.IAttribute.RenderAttr(wr)
}

type Value struct {
	doc.IAttribute
	ButtonAttribute
	OptionAttribute
	DataAttribute
	InputAttribute
	LiAttribute
	MeterAttribute
	ProgressAttribute
}

func (a Value) RenderAttr(wr io.Writer) (err error) {
	return a.IAttribute.RenderAttr(wr)
}
