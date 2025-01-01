package text

import (
	"go-html-compose/render"
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
	var err error
	if err = render.WriteTabBytes(wr, tabs); err != nil {
		return err
	}
	if err = t.Render(wr); err != nil {
		return err
	}
	return nil
}

func newText(value string, escapeHTML bool) *TextStruct {
	return &TextStruct{
		Value:      value,
		EscapeHTML: escapeHTML,
	}
}

func Text(value string) *TextStruct {
	return newText(value, true)
}

type RawTextStruct struct {
	*TextStruct
}

func RawText(value string) *RawTextStruct {
	return &RawTextStruct{
		TextStruct: newText(value, false),
	}
}
