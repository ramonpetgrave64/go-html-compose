package test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func CleanFormat(original string) string {
	re := regexp.MustCompile(`(<\w+)\s+(\w+)`)
	cleaned := re.ReplaceAllString(original, `$1 $2`)
	re = regexp.MustCompile(`(")\s+(\w+)`)
	cleaned = re.ReplaceAllString(cleaned, `$1 $2`)
	re = regexp.MustCompile(`(")\s+(>)`)
	cleaned = re.ReplaceAllString(cleaned, `$1$2`)
	cleaned = strings.ReplaceAll(cleaned, "\n", "")
	cleaned = strings.ReplaceAll(cleaned, "\t", "")
	return cleaned
}

func TestContentDiffErr(want string, got string) error {
	return fmt.Errorf("unexpected render value: \nwant: \n%s\n, got: \n%s\n", want, got)
}

func Diff(want string, got string) string {
	if want != got {
		return fmt.Sprintf(`
	- %s	
	+ %s
`, want, got)
	}
	return ""
}

func TestDiffError(t *testing.T, want string, got string) {
	t.Helper()
	if diff := Diff(want, got); diff != "" {
		t.Errorf("unexpected value (-want, +got): %s", diff)
	}
}
