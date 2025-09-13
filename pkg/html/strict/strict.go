package strict

import (
	sa "github.com/ramonpetgrave64/go-html-compose/pkg/html/strict/attrs"
	se "github.com/ramonpetgrave64/go-html-compose/pkg/html/strict/elems"
)

// Example usage demonstrating type safety
func Do() {
	// Valid attribute usage
	se.Button(
		sa.Name("my-button"),
		sa.Type("submit"),
		sa.Hidden("hidden"), // Global attributes are accepted
		sa.Class("ok"),
	)()

	se.Img(
		sa.Src("image.png"),
		sa.Alt("An image"),
		sa.Hidden("hidden"),
	)

	se.Script(
		sa.Src("main.js"),
		sa.Type("module"),
		sa.Hidden("hidden"),
	)()

	se.Script(
		sa.Src("main.js"),
		sa.Type("module"),
		sa.Hidden("hidden"),
		// sa.Enctype("ok"),
	)()

	// li := se.Li(sa.Class("ok"))()
	// li.RenderConent(nil)
	se.Ul()(
		// li,
		se.Li(sa.Class("ok"))(),
		// se.Button()(),
	)

	// The following lines would cause a compile-time error
	// because the attribute is not valid for the element.

	// Button(Alt("this is for images")) // COMPILE ERROR
	// Img(Name("this is for buttons")) // COMPILE ERROR
}
