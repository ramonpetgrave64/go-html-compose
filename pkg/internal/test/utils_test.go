package test

import (
	"testing"
)

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

type mockT struct {
	errorFCalled bool
}

func (t *mockT) Errorf(format string, args ...any) {
	t.errorFCalled = true
}

func (t mockT) Helper() {}

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
		tc := tc // create a new variable to hold the value of tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			testMockT := mockT{}
			TestDiffError(&testMockT, tc.a, tc.b)
			if tc.errorFCalled != testMockT.errorFCalled {
				t.Errorf("unexpected ErrorFCalled: (-want, +got):\n-%t\n+%t", tc.errorFCalled, testMockT.errorFCalled)
			}
		})
	}
}
