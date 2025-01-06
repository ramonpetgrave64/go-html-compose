package doc

import (
	"bytes"
	"go-html-compose/util"
	"testing"
)

func Test_If(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		want      string
		ifContent *testRenderable
		condition bool
	}{
		{
			name: "basic",
			want: `my-words`,
			ifContent: &testRenderable{
				data: []byte(`my-words`),
			},
			condition: true,
		},
		{
			name: "basic",
			want: ``,
			ifContent: &testRenderable{
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

	mapFunc := func(item string) Renderable {
		return testRenderable{
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
			content := Container(rendrs...)
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
