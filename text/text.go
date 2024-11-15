package text

import (
	"go-html-compose/util"
	"html"
	"io"
)

type TextStruct struct {
	Value string
}

func (text *TextStruct) Render(wr io.Writer) {
	wr.Write([]byte(html.EscapeString(text.Value)))
}

func (t *TextStruct) StructuredRenderWithTabs(wr io.Writer, tabs int) {
	wr.Write(util.GetTabBytes(tabs))
	t.Render(wr)
}

func Text(value string) *TextStruct {
	return &TextStruct{
		Value: value,
	}
}

type RawTextStruct struct {
	*TextStruct
}

func (t *RawTextStruct) Render(wr io.Writer) {
	wr.Write([]byte(t.Value))
}

func (t *RawTextStruct) StructuredRenderWithTabs(wr io.Writer, tabs int) {
	wr.Write(util.GetTabBytes(tabs))
	t.Render(wr)
}

func RawText(value string) *RawTextStruct {
	return &RawTextStruct{
		TextStruct: Text(value),
	}
}
