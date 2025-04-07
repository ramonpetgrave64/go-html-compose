# go-html-compose

A Go library for writing code that resembles the form and structure of html,
and then rendering to actual HTML at runtime.

## Usage

Write your content, like in the [test example](./pkg/html/html_test.go).

```go
package main

import (
	d "github.com/ramonpetgrave64/go-html-compose/pkg/doc"
	a "github.com/ramonpetgrave64/go-html-compose/pkg/html/attrs"
	e "github.com/ramonpetgrave64/go-html-compose/pkg/html/elems"
	"io"
)

func renderHomePage(wr io.Writer) error {
	content := content := e.Html(a.Lang("en"))(
		e.Head()(
			e.Meta(a.Charset("UTF-8")),
			e.Meta(a.Name("viewport"), a.Content("width=device-width, initial-scale=1.0")),
			e.Meta(a.Name("color-scheme"), a.Content("light dark")),
			e.Title()(d.TextS("Example HTML Page")),
			e.Link(a.Rel("stylesheet"), a.Href("https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css")),
		),
		e.Body()(
			e.Header(a.Class("container"))(
				e.H1()(d.TextS("Welcome to My Website")),
				e.Nav()(
					e.Ul()(
						e.Li()(e.A(a.Href("#home"))(d.TextS("Home"))),
						e.Li()(e.A(a.Href("#about"))(d.TextS("About"))),
						e.Li()(e.A(a.Href("#contact"))(d.TextS("Contact"))),
					),
				),
			),
			e.Main(a.Class("container"))(
				e.Section(a.Id("home"))(
					e.H2()(d.TextS("Home")),
					e.P()(d.TextS("This is the home section.")),
				),
				e.Section(a.Id("about"))(
					e.H2()(d.TextS("About")),
					e.P()(d.TextS("This is the about section.")),
				),
				e.Section(a.Id("contact"))(
					e.H2()(d.TextS("Contact")),
					e.P()(d.TextS("Please subscribe!🥺")),
					e.Form()(
						e.Div(a.Class("grid"))(
							e.Input(
								a.Type("text"),
								a.Name("firstname"),
								a.Placeholder("First name"),
								a.AriaProp("label", "First name"),
								a.Required(true),
							),
							e.Input(
								a.Type("email"),
								a.Name("email"),
								a.Placeholder("Email address"),
								a.AriaProp("label", "Email address"),
								a.Autocomplete("email"),
								a.Required(true),
							),
							e.Button(a.Type("submit"))(d.TextS("Subscribe")),
						),
						e.Fieldset()(
							e.Label(a.For("terms"))(
								e.Input(a.Type("checkbox"), a.Role("switch"), a.Id("terms"), a.Name("terms")),
								d.TextS("I agree to the "),
								e.A(a.Href("#"), a.Onclick("alert('Hello, World!')"))(d.TextS("Privacy Policy")),
							),
						),
					),
				),
			),
			e.Footer(a.Class("container"))(
				e.P()(d.TextS("© 2025 My Website")),
			),
		),
	)
	return content.Render(wr)
}
```

This produces a minified version of

```html
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta name="color-scheme" content="light dark" />
    <title>Example HTML Page</title>
    <link
      rel="stylesheet"
      href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css"
    />
  </head>

  <body>
    <header class="container">
      <h1>Welcome to My Website</h1>
      <nav>
        <ul>
          <li><a href="#home">Home</a></li>
          <li><a href="#about">About</a></li>
          <li><a href="#contact">Contact</a></li>
        </ul>
      </nav>
    </header>
    <main class="container">
      <section id="home">
        <h2>Home</h2>
        <p>This is the home section.</p>
      </section>
      <section id="about">
        <h2>About</h2>
        <p>This is the about section.</p>
      </section>
      <section id="contact">
        <h2>Contact</h2>
        <p>Please subscribe!🥺</p>
        <form>
          <div class="grid">
            <input
              type="text"
              name="firstname"
              placeholder="First name"
              aria-label="First name"
              required="required"
            /><input
              type="email"
              name="email"
              placeholder="Email address"
              aria-label="Email address"
              autocomplete="email"
              required="required"
            /><button type="submit">Subscribe</button>
          </div>
          <fieldset>
            <label for="terms"
              ><input type="checkbox" role="switch" id="terms" name="terms" />I
              agree to the
              <a href="#" onclick="alert('Hello, World!')"
                >Privacy Policy</a
              ></label
            >
          </fieldset>
        </form>
      </section>
    </main>
    <footer class="container">
      <p>© 2025 My Website</p>
    </footer>
  </body>
</html>
```

![example screenshot](./example_screenshot.png)

## Development

Generate and Test. See [generate](./pkg/generate/)

```
go generate ./...
go test ./...
```

Custom test helpers are in [`./pk/internal/test/utils.go`](./pkg/internal/test/utils.go).
