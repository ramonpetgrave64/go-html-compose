## Generate

This module generates the [`pkg/attr/attrs.go`](../attr/attrs.go) file from [`./spec.html`](./spec.html), called with `go generate ./...` by [`pkg/internal/generate.go`](../internal/generate.go). The spec is auto-downloaded from https://html.spec.whatwg.org/multipage/indices.html.

**Special Cases**

- `aria-*` is hardcoded to `AriaProp(propert, value string)`
- `data-*` is hardcoded to `DataProp(propert, value string)`
- `role` is hardcoded to `Role(value string)`.
