package util

import (
	"fmt"
)

var (
	NilContent     = []byte(``)
	TabContent     = []byte(`	`)
	SpaceContent   = []byte(` `)
	NewlineContent = []byte("\n")
)

func GetTabBytes(tabs int) []byte {
	tabBytes := []byte{}
	i := 0
	for i < tabs {
		tabBytes = append(tabBytes, TabContent...)
		i++
	}
	return tabBytes
}

func TestContentDiffErr(want string, got string) error {
	return fmt.Errorf("unexpected render value: \nwant: \n%s\n, got: \n%s\n", want, got)
}
