package doc

import (
	"html"
	"io"
)

type textStruct struct {
	// render.Renderable
	EscapeHTML bool
	Value      []byte
}

type rawTextStruct struct {
	*textStruct
}

func Text(value []byte) IContent {
	return newText(value, true)
}

func TextS(value string) IContent {
	return Text([]byte(value))
}

func RawText(value []byte) IContent {
	return &rawTextStruct{
		textStruct: newText(value, false),
	}
}

func RawTextS(value string) IContent {
	return RawText([]byte(value))
}

func (t *textStruct) RenderConent(wr io.Writer) (err error) {
	text := t.Value
	if t.EscapeHTML {
		text = []byte(html.EscapeString(string(text)))
	}
	_, err = wr.Write(text)
	return
}

func newText(value []byte, escapeHTML bool) *textStruct {
	return &textStruct{
		Value:      value,
		EscapeHTML: escapeHTML,
	}
}
