package attr

import (
	"bytes"
	"go-html-compose/util"
	"testing"
)

func Test_Render(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		content *AttributeStruct
	}{
		{
			name:    "basic",
			want:    `class="my-class"`,
			content: Class("my-class"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var buffer bytes.Buffer

			if err := tc.content.Render(&buffer); err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}

			got := buffer.String()

			if tc.want != got {
				t.Error(util.TestContentDiffErr(tc.want, got))
			}
		})
	}
}

func Test_StructuredRenderWithTabs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		content *AttributeStruct
		tabs    int
	}{
		{
			name:    "structured: 0 tabs",
			want:    `class="my-class"`,
			content: Class("my-class"),
			tabs:    0,
		},
		{
			name:    "structured: 1 tabs",
			want:    string(util.GetTabBytes(1)) + `class="my-class"`,
			content: Class("my-class"),
			tabs:    1,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var buffer bytes.Buffer

			if err := tc.content.StructuredRender(&buffer, tc.tabs); err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}

			got := buffer.String()

			if tc.want != got {
				t.Error(util.TestContentDiffErr(tc.want, got))
			}
		})
	}
}
