package container

import (
	"bytes"
	"go-html-compose/render"
	"go-html-compose/util"
	"io"
	"testing"
)

type TestRenderable struct {
	render.Renderable
	data []byte
}

func (r TestRenderable) Render(wr io.Writer) {
	wr.Write(r.data)
}

func (r TestRenderable) StructuredRender(wr io.Writer, tabs int) {
	wr.Write(util.GetTabBytes(tabs))
	wr.Write(r.data)
}

func Test_Container(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		want    string
		content ContainerStruct
		tabs    int
	}{
		{
			name:    "empty",
			want:    ``,
			content: *Container(),
			tabs:    0,
		},
		{
			name: "single",
			want: `
ok`,
			content: *Container(TestRenderable{data: []byte(`ok`)}),
		},
		{
			name: "multiple",
			want: `
ok
go`,
			content: *Container(TestRenderable{data: []byte(`ok`)}, TestRenderable{data: []byte(`go`)}),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var buffer bytes.Buffer

			tc.content.StructuredRender(&buffer, 0)
			got := buffer.String()

			if tc.want != got {
				t.Error(util.TestContentDiffErr(tc.want, got))
			}
		})
	}
}
