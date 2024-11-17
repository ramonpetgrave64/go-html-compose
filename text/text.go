package text

import (
	"go-html-compose/render"
	"go-html-compose/util"
	"html"
	"io"
)

type TextStruct struct {
	render.Renderable
	EscapeHTML bool
	Value      string
}

func (t *TextStruct) Render(wr io.Writer) {
	text := t.Value
	if t.EscapeHTML {
		text = html.EscapeString(text)
	}
	wr.Write([]byte(text))
}

func (t *TextStruct) StructuredRender(wr io.Writer, tabs int) {
	wr.Write(util.GetTabBytes(tabs))
	t.Render(wr)
}

func NewText(value string, escapeHTML bool) *TextStruct {
	return &TextStruct{
		Value:      value,
		EscapeHTML: escapeHTML,
	}
}

func Text(value string) *TextStruct {
	return NewText(value, true)
}

type RawTextStruct struct {
	*TextStruct
}

func RawText(value string) *RawTextStruct {
	return &RawTextStruct{
		TextStruct: NewText(value, false),
	}
}
