package test

import (
	"bytes"
	a "go-html-compose/attr"
	d "go-html-compose/doc"
	e "go-html-compose/elem"
	"testing"
)

func Test_Example(tt *testing.T) {
	tt.Parallel()

	worldString, err := d.Bytes(
		e.Span()(
			d.Text([]byte(`world`)),
		),
	)
	if err != nil {
		tt.Errorf("unexpected error: %s", err.Error())
	}

	otherString, err := d.Bytes(
		d.Text([]byte(`g`)),
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
			d.Text([]byte(`hello`)),
			e.Img(a.Class("i")),
		),
		d.RawText(worldString),
		d.RawText([]byte(`raw<html>raw`)),
		d.RawText(otherString),
		d.Text([]byte(`world`)),
	)
	if err := d.Render(&buffer, content); err != nil {
		tt.Errorf("unexpected error: %s", err.Error())
	}
	got := buffer.String()

	want := `<div class="big world" style="ok"><span>hello<img class="i"></span><span>world</span>raw<html>rawgworld</div>`

	if want != got {
		tt.Error(TestContentDiffErr(want, got))
	}

	buffer.Reset()
	if err := d.StructuredRender(&buffer, content); err != nil {
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
		tt.Error(TestContentDiffErr(want, got))
	}
}
