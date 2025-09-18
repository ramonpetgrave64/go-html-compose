package elems

import (
	"bytes"
	"testing"

	"github.com/ramonpetgrave64/go-html-compose/pkg/doc"
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
			name:    "single element",
			want:    `<li></li>`,
			content: Li()(),
		},
		{
			name: "allowed child element",
			want: `<ul><li></li></ul>`,
			content: Ul()(
				Li()(),
			),
		},
		// {
		// 	// not a real test case, should have compile-time error
		// 	name: "disallowed child element",
		// 	content: Script()(
		// 		Li()(),
		// 	),
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

		var li any = Li()()
		if _, ok := li.(UlChild); !ok {
			t.Errorf("expected li to be a UlChild, but it's not")
		}
	})

	t.Run("disallowed parent type", func(t *testing.T) {
		t.Parallel()

		var li any = Li()()
		if _, ok := li.(ScriptChild); ok {
			t.Errorf("expected li to not be ScriptChild, but it is")
		}
	})
}
