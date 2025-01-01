package text

import (
	"bytes"
	"go-html-compose/util"
	"testing"
)

func Test_Text(t *testing.T) {
	t.Parallel()

	tabs := 1
	tabContent := string(util.GetTabBytes(tabs))

	tests := []struct {
		name             string
		want             string
		content          *TextStruct
		structuredRender bool
	}{
		{
			name:             "basic: text",
			want:             `hello world`,
			content:          Text("hello world"),
			structuredRender: false,
		},
		{
			name:             "basic: structured text",
			want:             tabContent + `hello world`,
			content:          Text("hello world"),
			structuredRender: true,
		},
		{
			name:             "text: html escape",
			want:             `&lt;script&gt;alert(&#34;hello world&#34;)&lt;/script&gt;`,
			content:          Text(`<script>alert("hello world")</script>`),
			structuredRender: false,
		},
		{
			name:             "text: html escape: structured text",
			want:             tabContent + `&lt;script&gt;alert(&#34;hello world&#34;)&lt;/script&gt;`,
			content:          Text(`<script>alert("hello world")</script>`),
			structuredRender: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var buffer bytes.Buffer

			if tc.structuredRender {
				tc.content.StructuredRender(&buffer, tabs)
			} else {
				tc.content.Render(&buffer)
			}

			got := buffer.String()

			if tc.want != got {
				t.Error(util.TestContentDiffErr(tc.want, got))
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
			content: RawText("hello world"),
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
