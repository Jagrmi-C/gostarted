package lib

import (
	"strconv"
	"strings"
	"fmt"
)

type Test_struct struct {
	ID		string
}

func IsIntegerX2(s string) (result string) {
	strValue := strings.TrimSpace(s)
	val, err := strconv.ParseInt(strValue, 10, 64)

	if err != nil {
		return s
	}

	return fmt.Sprint(val * 2, "\n")
}
