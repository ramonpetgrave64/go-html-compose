package text

import (
	"go-html-compose/util"
	"html"
	"io"
)

type TextStruct struct {
	// render.Renderable
	EscapeHTML bool
	Value      string
}

func (t *TextStruct) Render(wr io.Writer) error {
	text := t.Value
	if t.EscapeHTML {
		text = html.EscapeString(text)
	}
	if _, err := wr.Write([]byte(text)); err != nil {
		return err
	}
	return nil
}

func (t *TextStruct) StructuredRender(wr io.Writer, tabs int) error {
	if _, err := wr.Write(util.GetTabBytes(tabs)); err != nil {
		return err
	}
	if err := t.Render(wr); err != nil {
		return err
	}
	return nil
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
