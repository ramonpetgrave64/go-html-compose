package attr

import (
	"bytes"
	"testing"
)

func Test_Attribute(t *testing.T) {
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

			tc.content.Render(&buffer)
			got := buffer.String()

			if tc.want != got {
				t.Errorf("unexpected renfer value: \nwant: \n%s\n, got: \n%s\n", tc.want, got)
			}
		})
	}
}
