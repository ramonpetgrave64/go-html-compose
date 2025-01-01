package elem

import (
	"bytes"
	"go-html-compose/attr"
	"go-html-compose/render"
	"go-html-compose/util"
	"testing"
)

func Test_StructuredRenderWithTabs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		content render.Renderable
	}{
		{
			name: "basic: html",
			want: `<html>
</html>`,
			content: HTML()(),
		},
		{
			name: "basic: nested: single",
			want: `<html>
	<div>
	</div>
</html>`,
			content: HTML()(
				Div()(),
			),
		},
		{
			name:    "basic: tag: single",
			want:    `<img class="big">`,
			content: Img(attr.Class("big")),
		},
		{
			name: "basic: tag: multiple",
			want: `<img class="big" src="https://example.com/favicon">`,
			content: Img(
				attr.Class("big"),
				attr.Src("https://example.com/favicon"),
			),
		},
		{
			name: "basic nested: single: attrubute: single",
			want: `<html>
	<div class="my-class">
	</div>
</html>`,
			content: HTML()(
				Div(attr.Class("my-class"))(),
			),
		},
		{
			name: "basic nested: deep",
			want: `<html>
	<div>
		<div>
			<span>
			</span>
			<img>
		</div>
	</div>
</html>`,
			content: HTML()(
				Div()(
					Div()(
						Span()(),
						Img(),
					),
				),
			),
		},
		{
			name: "basic nested: multiple",
			want: `<html>
	<div>
	</div>
	<div>
	</div>
</html>`,
			content: HTML()(
				Div()(),
				Div()(),
			),
		},
		{
			name:    "unit tag",
			content: Div()(Img()),
			want: `<div>
	<img>
</div>`,
		},
		{
			name:    "unit tag: single attribute",
			content: Div()(Img(attr.Class("my-class"))),
			want: `<div>
	<img
		class="my-class"
	>
</div>`,
		},
		{
			name:    "unit tag: multiple attributes",
			content: Div()(Img(attr.Class("my-class"), attr.Src("my-src"))),
			want: `<div>
	<img
		class="my-class"
		src="my-src"
	>
</div>`,
		},
	}
	tests = tests[len(tests)-3:]
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var buffer bytes.Buffer

			if err := tc.content.StructuredRender(&buffer, 0); err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}
			got := buffer.String()

			if tc.want != got {
				t.Error(util.TestContentDiffErr(tc.want, got))
			}
		})
	}
}

func Test_Render(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		content render.Renderable
	}{
		{
			name:    "basic: html",
			want:    `<html></html>`,
			content: HTML()(),
		},
		{
			name: "basic: nested: single",
			want: `<html><div></div></html>`,
			content: HTML()(
				Div()(),
			),
		},
		{
			name:    "basic: tag: single",
			want:    `<img class="big">`,
			content: Img(attr.Class("big")),
		},
		{
			name: "basic: tag: multiple",
			want: `<img class="big" src="https://example.com/favicon">`,
			content: Img(
				attr.Class("big"),
				attr.Src("https://example.com/favicon"),
			),
		},
		{
			name: "basic nested: single: attrubute: single",
			want: `<html><div class="my-class"></div></html>`,
			content: HTML()(
				Div(attr.Class("my-class"))(),
			),
		},
		{
			name: "basic nested: deep",
			want: `<html><div><div><span></span><img></div></div></html>`,
			content: HTML()(
				Div()(
					Div()(
						Span()(),
						Img(),
					),
				),
			),
		},
		{
			name: "basic nested: multiple",
			want: `<html><div></div><div></div></html>`,
			content: HTML()(
				Div()(),
				Div()(),
			),
		},
		{
			name:    "unit tag",
			content: Div()(Img()),
			want:    `<div><img></div>`,
		},
		{
			name:    "unit tag: single attribute",
			content: Div()(Img(attr.Class("my-class"))),
			want:    `<div><img class="my-class"></div>`,
		},
		{
			name:    "unit tag: multiple attributes",
			content: Div()(Img(attr.Class("my-class"), attr.Src("my-src"))),
			want:    `<div><img class="my-class" src="my-src"></div>`,
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
