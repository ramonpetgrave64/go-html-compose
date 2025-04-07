# Generate

This module generates the [`pkg/html/attrs/attrs.go`](../html/attrs/attrs.go) and [`pkg/html/elems/elems.go`](../html/elems/elems.go) files from [`./spec.html`](./spec.html), called with `go generate ./...` by [`pkg/generate.go`](../generate.go). The spec is auto-downloaded from https://html.spec.whatwg.org/multipage/indices.html.

(TODO: Confirm) We keep this as a module so that `go build` of the library does not include this code.
This module should not be released to Go module proxies.

## Special Cases

**Attributes**

- `aria-*` is hardcoded to `AriaProp(propert, value string)`
- `data-*` is hardcoded to `DataProp(propert, value string)`
- `role` is hardcoded to `Role(value string)`

**Tags**

- The `DOCTYPE` is hardcoded to `var Doctype = UnitTag("!DOCTYPE html")`
