package attr

import (
	"bytes"
	"go-html-compose/test"
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
			test.TestDiffError(t, tc.want, got)
		})
	}
}
