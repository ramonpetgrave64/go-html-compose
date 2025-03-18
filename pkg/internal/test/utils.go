package test

import (
	"fmt"
)

// Diff shows a non-empty diff message if the two strings are not equal. e.g.,
// -<want text>
// +<got text>
func Diff(want, got string) string {
	if want != got {
		return fmt.Sprintf(`
-%s
+%s
`, want, got)
	}
	return ""
}

// testingT is an interface for testing.T, for testing TestDiffError().
type testingT interface {
	Errorf(format string, args ...any)
	Helper()
}

// TestDiffError given a testingT, calls Errorf if their is a difference between want and got.
func TestDiffError(t testingT, want, got string) {
	t.Helper()
	if diff := Diff(want, got); diff != "" {
		t.Errorf("unexpected value (-want, +got): %s", diff)
	}
}
