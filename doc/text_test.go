package doc

import (
	"bytes"
	"go-html-compose/test"
	"testing"
)

func Test_Text(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		content *TextStruct
	}{
		{
			name:    "text",
			want:    `hello world`,
			content: Text([]byte(`hello world`)),
		},
		{
			name:    "html escape",
			want:    `&lt;script&gt;alert(&#34;hello world&#34;)&lt;/script&gt;`,
			content: Text([]byte(`<script>alert("hello world")</script>`)),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var buffer bytes.Buffer
			tc.content.Render(&buffer)
			got := buffer.String()
			if tc.want != got {
				t.Error(test.TestContentDiffErr(tc.want, got))
			}
		})
	}
}

func Test_RawText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		content *RawTextStruct
	}{
		{
			name:    "basic: rawtext",
			want:    `hello world`,
			content: RawText([]byte(`hello world`)),
		},
		{
			name:    "rawtext: no html escape",
			want:    `<script>alert("hello world")</script>`,
			content: RawText([]byte(`<script>alert("hello world")</script>`)),
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
				t.Error(test.TestContentDiffErr(tc.want, got))
			}
		})
	}
}
