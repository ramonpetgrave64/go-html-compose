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
			<span>
			</span>
		</div>
	</html>
			`,
			want: "<html><div><span></span></div></html>",
		},
		{
			name: "attributes",
			text: `
	<html>
		<div class="blah">
		</div>
		<div
			class="abc"
		>
		</div>
	</html>
			`,
			want: `<html><div class="blah"></div><div class="abc"></div></html>`,
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
		name string
		a    string
		b    string
		diff string
	}{
		{
			name: "empty",
			a:    "",
			b:    "",
			diff: "",
		},
		{
			name: "equal",
			a:    "abc",
			b:    "abc",
			diff: "",
		},
		{
			name: "unequal",
			a:    "x",
			b:    "abc",
			diff: `
-x
+abc
`,
		},
		{
			name: "empty, unequal",
			a:    "",
			b:    "abc",
			diff: `
-
+abc
`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			diff := Diff(tc.a, tc.b)
			if tc.diff != diff {
				t.Errorf("unexpected diff: (-want, +got):\n-%s\n+%s", tc.diff, diff)
			}
		})
	}
}

type testT struct {
	errorFCalled bool
}

func (t *testT) Errorf(format string, args ...any) {
	t.errorFCalled = true
}

func (t testT) Helper() {}

func Test_TestDiffError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		a            string
		b            string
		errorFCalled bool
	}{
		{
			name:         "diff",
			a:            "x",
			b:            "abc",
			errorFCalled: true,
		},
		{
			name:         "no diff",
			a:            "x",
			b:            "x",
			errorFCalled: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			testT := testT{}
			TestDiffError(&testT, tc.a, tc.b)
			if tc.errorFCalled != testT.errorFCalled {
				t.Errorf("unexpected ErrorFCalled: (-want, +got):\n-%t\n+%t", tc.errorFCalled, testT.errorFCalled)
			}
		})
	}
	t.Run("diff", func(t *testing.T) {
		testT := &testT{}
		TestDiffError(testT, "x", "abc")

	})
}
