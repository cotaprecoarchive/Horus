package util

import (
	"fmt"
	"os"
)

func Invariant(condition bool, args ...interface{}) {
	if condition {
		return
	}

	fmt.Println(
		"Invariant Violation: " +
			fmt.Sprintf(args[0].(string), args[1:]...),
	)

	os.Exit(1)
}
