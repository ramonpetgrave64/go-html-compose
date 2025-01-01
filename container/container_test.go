package container

import (
	"bytes"
	"go-html-compose/util"
	"io"
	"testing"
)

type TestRenderable struct {
	// render.Renderable
	data []byte
}

func (r TestRenderable) Render(wr io.Writer) error {
	if _, err := wr.Write(r.data); err != nil {
		return err
	}
	return nil
}

func (r TestRenderable) StructuredRender(wr io.Writer, tabs int) error {
	if _, err := wr.Write(util.GetTabBytes(tabs)); err != nil {
		return err
	}
	if _, err := wr.Write(r.data); err != nil {
		return err
	}
	return nil
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
