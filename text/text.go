package text

import (
	"go-html-compose/render"
	"html"
	"io"
)

type TextStruct struct {
	// render.Renderable
	EscapeHTML bool
	Value      []byte
}

func (t *TextStruct) Render(wr io.Writer) error {
	text := t.Value
	if t.EscapeHTML {
		text = []byte(html.EscapeString(string(text)))
	}
	if _, err := wr.Write(text); err != nil {
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

func newText(value []byte, escapeHTML bool) *TextStruct {
	return &TextStruct{
		Value:      value,
		EscapeHTML: escapeHTML,
	}
}

func Text(value []byte) *TextStruct {
	return newText(value, true)
}

type RawTextStruct struct {
	*TextStruct
}

func RawText(value []byte) *RawTextStruct {
	return &RawTextStruct{
		TextStruct: newText(value, false),
	}
}
