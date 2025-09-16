package attrs

import "github.com/ramonpetgrave64/go-html-compose/pkg/html/strict/elems"

type SrcType interface {
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
