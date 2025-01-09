package doc

import (
	"bytes"
	"go-html-compose/test"
	"io"
	"testing"
)

type testRenderable struct {
	// render.Renderable
	data []byte
}

func (r testRenderable) Render(wr io.Writer) error {
	if _, err := wr.Write(r.data); err != nil {
		return err
	}
	return nil
}

func (r testRenderable) StructuredRender(wr io.Writer, tabs int) error {
	var err error
	if err = WriteTabBytes(wr, tabs); err != nil {
		return err
	}
	if _, err = wr.Write(r.data); err != nil {
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
			content: *Container(testRenderable{data: []byte(`ok`)}),
		},
		{
			name: "multiple",
			want: `
ok
go`,
			content: *Container(testRenderable{data: []byte(`ok`)}, testRenderable{data: []byte(`go`)}),
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
				t.Error(test.TestContentDiffErr(tc.want, got))
			}
		})
	}
}
