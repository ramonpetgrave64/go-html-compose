package doc

import (
	"bytes"
	"io"
	"testing"

	"go-html-compose/pkg/internal/test"
)

type testRenderable struct {
	data []byte
}

func (r testRenderable) RenderConent(wr io.Writer) error {
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
		content ContContainerStruct
	}{
		{
			name:    "empty",
			want:    ``,
			content: *ContContainer(),
		},
		{
			name:    "single",
			want:    `ok`,
			content: *ContContainer(testRenderable{data: []byte(`ok`)}),
		},
		{
			name:    "multiple",
			want:    `okgo`,
			content: *ContContainer(testRenderable{data: []byte(`ok`)}, testRenderable{data: []byte(`go`)}),
		},
	}
	for _, tc := range tests {
		tc := tc // create a new variable to hold the value of tc

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
