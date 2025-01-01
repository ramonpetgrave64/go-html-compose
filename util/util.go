package util

import (
	"fmt"
)

func TestContentDiffErr(want string, got string) error {
	return fmt.Errorf("unexpected render value: \nwant: \n%s\n, got: \n%s\n", want, got)
}
