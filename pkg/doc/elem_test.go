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
		content IContent
	}{
		{
			name:    "unit",
			want:    `<img>`,
			content: ChildElem("img"),
		},
		{
			name:    "unit with attribute",
			want:    `<img class="c1">`,
			content: ChildElem("img", Attr("class", "c1")),
		},
		{
			name:    "unit with multiple attributes",
			want:    `<img class="c1" aria-label="logo">`,
			content: ChildElem("img", Attr("class", "c1"), Attr("aria-label", "logo")),
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

func Test_ParentTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		content IContent
	}{
		{
			name:    "single parent",
			want:    `<div></div>`,
			content: ParentElem("div")(),
		},
		{
			name:    "single parent with single attribute",
			want:    `<div class="c1"></div>`,
			content: ParentElem("div", Attr("class", "c1"))(),
		},
		{
			name:    "single parent with multiple attributes",
			want:    `<div class="c1" aria-label="logo"></div>`,
			content: ParentElem("div", Attr("class", "c1"), Attr("aria-label", "logo"))(),
		},
		{
			name: "single nested",
			want: `<div><div></div></div>`,
			content: ParentElem("div")(
				ParentElem("div")(),
			),
		},
		{
			name: "multiple nested",
			want: `<div><div></div><img></div>`,
			content: ParentElem("div")(
				ParentElem("div")(),
				ChildElem("img"),
			),
		},
		{
			name: "deeply nested",
			want: `<div><div><span></span></div><img></div>`,
			content: ParentElem("div")(
				ParentElem("div")(
					ParentElem("span")(),
				),
				ChildElem("img"),
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
