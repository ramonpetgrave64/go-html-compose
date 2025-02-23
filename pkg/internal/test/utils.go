package test

import (
	"fmt"
)

func TestContentDiffErr(want, got string) error {
	return fmt.Errorf("unexpected render value: \nwant: \n%s\n, got: \n%s", want, got)
}

func Diff(want, got string) string {
	if want != got {
		return fmt.Sprintf(`
-%s
+%s
`, want, got)
	}
	return ""
}

type T interface {
	Errorf(format string, args ...any)
	Helper()
}

func TestDiffError(t T, want, got string) {
	t.Helper()
	if diff := Diff(want, got); diff != "" {
		t.Errorf("unexpected value (-want, +got): %s", diff)
	}
}
