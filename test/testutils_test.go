package test

import (
	"testing"
)

func Test_CleanFormat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "empty",
			text: "",
			want: "",
		},
		{
			name: "tabs",
			text: "\t\t",
			want: "",
		},
		{
			name: "newlines",
			text: "\n\n\n",
			want: "",
		},
		{
			name: "mix",
			text: "\n\t\n",
			want: "",
		},
		{
			name: "normal",
			text: `
	<html>
		<div>
		</div>
	</html>
			`,
			want: "<html><div></div></html>",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := CleanFormat(tc.text)
			TestDiffError(t, tc.want, got)
		})
	}
}

func Test_Diff(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		a      string
		b      string
		result bool
	}{
		{
			name:   "empty",
			a:      "",
			b:      "",
			result: false,
		},
		{
			name:   "equal",
			a:      "abc",
			b:      "abc",
			result: false,
		},
		{
			name:   "unequal",
			a:      "x",
			b:      "abc",
			result: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// TODO: test against the return value of Diff()
			result := Diff(tc.a, tc.b) != ""
			if tc.result != result {
				t.Error("unexpected diff")
			}
		})
	}
}
