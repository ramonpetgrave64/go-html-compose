package doc

import (
	"bytes"
	"testing"

	"go-html-compose/pkg/internal/test"
)

func Test_If(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		want      string
		ifContent *testContentStruct
		condition bool
	}{
		{
			name: "true",
			want: `my-words`,
			ifContent: &testContentStruct{
				data: []byte(`my-words`),
			},
			condition: true,
		},
		{
			name: "false",
			want: ``,
			ifContent: &testContentStruct{
				data: []byte(`my-words`),
			},
			condition: false,
		},
	}
	for _, tc := range tests {
		tc := tc // create a new variable to hold the value of tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var buffer bytes.Buffer
			content := IfCont(tc.condition, tc.ifContent)
			if err := content.RenderConent(&buffer); err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}
			got := buffer.String()
			test.TestDiffError(t, tc.want, got)
		})
	}
}

func Test_MapToContainer(t *testing.T) {
	t.Parallel()

	mapFunc := func(item string) IContent {
		return testContentStruct{
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
		tc := tc // create a new variable to hold the value of tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var buffer bytes.Buffer
			content := MapToContContainer(tc.items, mapFunc)
			if err := content.RenderConent(&buffer); err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}
			got := buffer.String()
			test.TestDiffError(t, tc.want, got)
		})
	}
}
