package elem

import (
	"bytes"
	"go-html-compose/attr"
	"go-html-compose/render"
	"testing"
)

func Test_StructuredRender(t *testing.T) {
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
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var buffer bytes.Buffer

			tc.content.StructuredRenderWithTabs(&buffer, 0)
			got := buffer.String()

			if tc.want != got {
				t.Errorf("unexpected renfer value: \nwant: \n%s\n, got: \n%s\n", tc.want, got)
			}
		})
	}
}
