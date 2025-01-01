package cond

import (
	"bytes"
	"errors"
	"go-html-compose/doc"
	"go-html-compose/render"
	"go-html-compose/util"
	"io"
	"testing"
)

type TestRenderable struct {
	// render.Renderable
	data []byte
}

func (r TestRenderable) Render(wr io.Writer) error {
	if _, err := wr.Write(r.data); err != nil {
		return err
	}
	return nil
}

func (r TestRenderable) StructuredRender(wr io.Writer, tabs int) error {
	return errors.New("not implemented")
}

func Test_If(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		want      string
		ifContent *TestRenderable
		condition bool
	}{
		{
			name: "basic",
			want: `my-words`,
			ifContent: &TestRenderable{
				data: []byte(`my-words`),
			},
			condition: true,
		},
		{
			name: "basic",
			want: ``,
			ifContent: &TestRenderable{
				data: []byte(`my-words`),
			},
			condition: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var buffer bytes.Buffer
			content := If(tc.condition, tc.ifContent)
			if err := content.Render(&buffer); err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}

			got := buffer.String()

			if tc.want != got {
				t.Error(util.TestContentDiffErr(tc.want, got))
			}
		})
	}
}

func Test_Map(t *testing.T) {
	t.Parallel()

	mapFunc := func(item string) render.Renderable {
		return TestRenderable{
			data: []byte(`*item: ` + item),
		}
	}

	tests := []struct {
		name  string
		want  string
		items []string
	}{
		{
			name:  "0 items",
			want:  ``,
			items: []string{},
		},
		{
			name:  "1 item",
			want:  `*item: my-words`,
			items: []string{`my-words`},
		},
		{
			name:  "2 items",
			want:  `*item: my-words*item: are good`,
			items: []string{`my-words`, `are good`},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var buffer bytes.Buffer
			rendrs := Map(tc.items, mapFunc)
			content := doc.Document(rendrs...)
			if err := content.Render(&buffer); err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}

			got := buffer.String()

			if len(tc.items) != len(rendrs) {
				t.Errorf("expected equal items: want: %d, got: %d", len(tc.items), len(rendrs))
			}

			if tc.want != got {
				t.Error(util.TestContentDiffErr(tc.want, got))
			}
		})
	}
}
