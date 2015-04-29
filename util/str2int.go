package util

import "strconv"

func Str2int(str string) int {
	var val, err = strconv.Atoi(str)

	if err != nil {
		return 0
	}

	return val
}
