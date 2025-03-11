package doc

import (
	"bytes"
	"testing"

	"go-html-compose/pkg/internal/test"
)

func Test_UnitTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		content Renderable
	}{
		{
			name:    "unit",
			want:    `<img>`,
			content: UnitTag("img"),
		},
		{
			name:    "unit with attribute",
			want:    `<img class="c1">`,
			content: UnitTag("img", Attr("class", "c1")),
		},
		{
			name:    "unit with multiple attributes",
			want:    `<img class="c1" aria-label="logo">`,
			content: UnitTag("img", Attr("class", "c1"), Attr("aria-label", "logo")),
		},
	}
	for _, tc := range tests {
		tc := tc

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

func Test_ParentTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		content Renderable
	}{
		{
			name:    "single parent",
			want:    `<div></div>`,
			content: ParentTag("div")(),
		},
		{
			name:    "single parent with single attribute",
			want:    `<div class="c1"></div>`,
			content: ParentTag("div", Attr("class", "c1"))(),
		},
		{
			name:    "single parent with multiple attributes",
			want:    `<div class="c1" aria-label="logo"></div>`,
			content: ParentTag("div", Attr("class", "c1"), Attr("aria-label", "logo"))(),
		},
		{
			name: "single nested",
			want: `<div><div></div></div>`,
			content: ParentTag("div")(
				ParentTag("div")(),
			),
		},
		{
			name: "multiple nested",
			want: `<div><div></div><img></div>`,
			content: ParentTag("div")(
				ParentTag("div")(),
				UnitTag("img"),
			),
		},
		{
			name: "deeply nested",
			want: `<div><div><span></span></div><img></div>`,
			content: ParentTag("div")(
				ParentTag("div")(
					ParentTag("span")(),
				),
				UnitTag("img"),
			),
		},
	}
	for _, tc := range tests {
		tc := tc

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
