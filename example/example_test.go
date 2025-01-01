package example

import (
	"bytes"
	a "go-html-compose/attr"
	e "go-html-compose/elem"
	r "go-html-compose/render"
	t "go-html-compose/text"
	"go-html-compose/util"
	"testing"
)

func Test_Example(tt *testing.T) {
	tt.Parallel()

	worldString, err := r.Bytes(
		e.Span()(
			t.Text([]byte(`world`)),
		),
	)
	if err != nil {
		tt.Errorf("unexpected error: %s", err.Error())
	}

	otherString, err := r.Bytes(
		t.Text([]byte(`g`)),
	)
	if err != nil {
		tt.Errorf("unexpected error: %s", err.Error())
	}

	var buffer bytes.Buffer
	content := e.Div(
		a.Class("big world"),
		a.Style("ok"),
	)(
		e.Span()(
			t.Text([]byte(`hello`)),
			e.Img(a.Class("i")),
		),
		t.RawText(worldString),
		t.RawText([]byte(`raw<html>raw`)),
		t.RawText(otherString),
		t.Text([]byte(`world`)),
	)
	if err := r.Render(&buffer, content); err != nil {
		tt.Errorf("unexpected error: %s", err.Error())
	}
	got := buffer.String()

	want := `<div class="big world" style="ok"><span>hello<img class="i"></span><span>world</span>raw<html>rawgworld</div>`

	if want != got {
		tt.Error(util.TestContentDiffErr(want, got))
	}

	buffer.Reset()
	if err := r.StructuredRender(&buffer, content); err != nil {
		tt.Errorf("unexpected error: %s", err.Error())
	}
	got = buffer.String()

	want = `<div
	class="big world"
	style="ok"
>
	<span>
		hello
		<img
			class="i"
		>
	</span>
	<span>world</span>
	raw<html>raw
	g
	world
</div>`

	if want != got {
		tt.Error(util.TestContentDiffErr(want, got))
	}
}
