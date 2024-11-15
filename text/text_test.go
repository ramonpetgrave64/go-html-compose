package text

import (
	"bytes"
	"go-html-compose/render"
	"testing"
)

func Test_Text(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		content render.Renderable
	}{
		{
			name:    "basic: text",
			want:    `hello world`,
			content: Text("hello world"),
		},
		{
			name:    "basic: rawtext",
			want:    `hello world`,
			content: RawText("hello world"),
		},
		{
			name:    "text: html escape",
			want:    `&lt;script&gt;alert(&#34;hello world&#34;)&lt;/script&gt;`,
			content: Text(`<script>alert("hello world")</script>`),
		},
		{
			name:    "rawtext: no html escape",
			want:    `<script>alert("hello world")</script>`,
			content: RawText(`<script>alert("hello world")</script>`),
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
