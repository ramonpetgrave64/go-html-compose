package attrs

import (
	"bytes"
	"testing"

	"github.com/ramonpetgrave64/go-html-compose/pkg/doc"
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
			content: elems.Li(Value("my-value"))(),
		},
		{
			name:    "global attribute",
			want:    `<li onclick="alert('hello')"></li>`,
			content: elems.Li(Onclick("alert('hello')"))(),
		},
		// {
		// 	// not a real test case, should have compile-time error
		// 	name:    "disallowed child element",
		// 	content: elems.Li(Src("my-value"))(),
		// },
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

		var value any = Value("my-value")
		if _, ok := value.(elems.LiAttribute); !ok {
			t.Errorf("expected value to be a UlAttribute, but it's not")
		}
	})

	t.Run("disallowed parent type", func(t *testing.T) {
		t.Parallel()

		var value any = Value("my-value")
		if _, ok := value.(elems.ScriptAttribute); ok {
			t.Errorf("expected value to not be ScriptAttribute, but it is")
		}
	})
}
