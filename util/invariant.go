package util

import (
	"bytes"
	"fmt"
	"os"
)

func Invariant(condition bool, args ...interface{}) {
	if condition {
		return
	}

	var buffer bytes.Buffer

	buffer.WriteString("Invariant Violation: ")
	buffer.WriteString(fmt.Sprintf(args[0].(string), args[1:]...))

	fmt.Println(buffer.String())

	os.Exit(1)
}
