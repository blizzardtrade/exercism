package reverse

import (
	"fmt"
)

func Reverse(input string) string {
	str := ""
	for i := len(input) - 1; i >= 0; i-- {
		str += string(input[i])
	}
	fmt.Println(str)

	return str
}

