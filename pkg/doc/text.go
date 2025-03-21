package doc

import (
	"html"
	"io"
)

type TextStruct struct {
	// render.Renderable
	EscapeHTML bool
	Value      []byte
}

func (t *TextStruct) Render(wr io.Writer) (err error) {
	text := t.Value
	if t.EscapeHTML {
		text = []byte(html.EscapeString(string(text)))
	}
	err = WriteByteSlices(wr, text)
	return
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

func TextS(value string) *TextStruct {
	return Text([]byte(value))
}

type RawTextStruct struct {
	*TextStruct
}

func RawText(value []byte) *RawTextStruct {
	return &RawTextStruct{
		TextStruct: newText(value, false),
	}
}

func RawTextS(value string) *RawTextStruct {
	return RawText([]byte(value))
}
