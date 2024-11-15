package example

import (
	"bytes"
	a "go-html-compose/attr"
	e "go-html-compose/elem"
	r "go-html-compose/render"
	t "go-html-compose/text"
	"testing"
)

func Test_Example(tt *testing.T) {
	tt.Parallel()

	var buffer bytes.Buffer
	content := e.Div(
		a.Class("big world"),
		a.Style("ok"),
	)(
		e.Span()(
			t.Text("hello"),
			e.Img(a.Class("i")),
		),
		t.RawText(r.String(
			e.Span()(
				t.Text("world"),
			),
		)),
		t.RawText("raw<html>raw"),
		t.RawText(r.String(t.Text("g"))),
		t.Text("world"),
	)
	r.Render(&buffer, content)
	got := buffer.String()

	want := `<div class="big world" style="ok"><span>hello<img class="i"></span><span>world</span>raw<html>rawgworld</div>`

	if want != got {
		tt.Errorf("unexpected renfer value: \nwant: \n%s\n, got: \n%s\n", want, got)
	}

	buffer.Reset()
	r.StructuredRender(&buffer, content)
	got = buffer.String()

	want = `<div class="big world" style="ok">
	<span>
		hello
		<img class="i">
	</span>
	<span>world</span>
	raw<html>raw
	g
	world
</div>`

	if want != got {
		tt.Errorf("unexpected renfer value: \nwant: \n%s\n, got: \n%s\n", want, got)
	}
}
