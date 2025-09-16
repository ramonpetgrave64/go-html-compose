package strict

import (
	"bytes"
	"testing"

	"github.com/ramonpetgrave64/go-html-compose/pkg/doc"
	"github.com/ramonpetgrave64/go-html-compose/pkg/html/strict/attrs"
	"github.com/ramonpetgrave64/go-html-compose/pkg/html/strict/elems"
	"github.com/ramonpetgrave64/go-html-compose/pkg/internal/test"
)

func Test_AllowedTypes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		content doc.IContent
	}{
		{
			name:    "allowed attribute",
			want:    `<li value="my-value"></li>`,
			content: elems.Li(attrs.Value("my-value"))(),
		},
		{
			name:    "global attribute",
			want:    `<li onclick="alert('hello')"></li>`,
			content: elems.Li(attrs.Onclick("alert('hello')"))(),
		},
		{
			name: "allowed child element",
			want: `<ul><li></li></ul>`,
			content: elems.Ul()(
				elems.Li()(),
			),
		},
	}
	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var buffer bytes.Buffer
			if err := tc.content.RenderConent(&buffer); err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}
			got := buffer.String()
			test.TestDiffError(t, tc.want, got)
		})
	}
}

func Test_ParentTypes(t *testing.T) {
	t.Parallel()

	t.Run("allowed parent type", func(t *testing.T) {
		t.Parallel()

		li := elems.Li()()
		if _, ok := li.(elems.UlChild); !ok {
			t.Errorf("expected li to be a UlChild, but it's not")
		}
	})

	t.Run("disallowed parent type", func(t *testing.T) {
		t.Parallel()

		li := elems.Li()()
		if _, ok := li.(elems.ScriptChild); ok {
			t.Errorf("expected li to not be ScriptChild, but it is")
		}
	})
}
