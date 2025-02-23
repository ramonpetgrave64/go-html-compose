package elem

import (
	"bytes"
	"testing"

	"go-html-compose/pkg/attr"
	"go-html-compose/pkg/doc"
	"go-html-compose/pkg/internal/test"
)

func Test_UnitTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		content doc.Renderable
	}{
		{
			name:    "unit",
			want:    `<img>`,
			content: Img(),
		},
		{
			name:    "unit with attribute",
			want:    `<img class="c1">`,
			content: Img(attr.Class("c1")),
		},
		{
			name:    "unit with multiple attributes",
			want:    `<img class="c1" aria-label="logo">`,
			content: Img(attr.Class("c1"), attr.AriaLabel("logo")),
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

func Test_ParentTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		content doc.Renderable
	}{
		{
			name:    "single parent",
			want:    `<div></div>`,
			content: Div()(),
		},
		{
			name:    "single parent with single attribute",
			want:    `<div class="c1"></div>`,
			content: Div(attr.Class("c1"))(),
		},
		{
			name:    "single parent with multiple attributes",
			want:    `<div class="c1" aria-label="logo"></div>`,
			content: Div(attr.Class("c1"), attr.AriaLabel("logo"))(),
		},
		{
			name: "single nested",
			want: `<div><div></div></div>`,
			content: Div()(
				Div()(),
			),
		},
		{
			name: "multiple nested",
			want: `<div><div></div><img></div>`,
			content: Div()(
				Div()(),
				Img(),
			),
		},
		{
			name: "deeply nested",
			want: `<div><div><span></span></div><img></div>`,
			content: Div()(
				Div()(
					Span()(),
				),
				Img(),
			),
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
