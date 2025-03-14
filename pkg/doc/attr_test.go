package doc

import (
	"bytes"
	"testing"

	"go-html-compose/pkg/internal/test"
)

func Test_Attr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		content *AttrStruct
	}{
		{
			name:    "attribute",
			want:    `class="my-class"`,
			content: Attr("class", "my-class"),
		},
		{
			name:    "attribute: escaped",
			want:    `onclick="alert(&#39;Hi!&#39;)"`,
			content: Attr("onclick", "alert('Hi!')"),
		},
		{
			name:    "attribute: uescaped",
			want:    `onclick="alert('Hi!')"`,
			content: RawAttr("onclick", "alert('Hi!')"),
		},
		{
			name:    "boolean attribute: true",
			want:    `selected="selected"`,
			content: BooleanAttr("selected", true),
		},
		{
			name:    "boolean attribute: false",
			want:    ``,
			content: BooleanAttr("selected", false),
		},
	}
	for _, tc := range tests {
		tc := tc // create a new variable to hold the value of tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var buffer bytes.Buffer
			if err := tc.content.RenderAttr(&buffer); err != nil {
				t.Errorf("unexpected error: %s", err.Error())

			}
			got := buffer.String()
			test.TestDiffError(t, tc.want, got)
		})
	}
}
