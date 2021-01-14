package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Valid(input string) bool {
	var number []int
	input = strings.ReplaceAll(input, " ", "")

	r := []rune(input)
	for _, v := range r {
		if unicode.IsDigit(v) != true || len(input) == 1 || len(input) == 0 {
			return false
		}
		n, _ := strconv.Atoi(string(v))
		number = append(number, n)
	}

	sum := 0
	count := 0

	for i := len(number) - 1; i >= 0; i-- {
		count++
		fmt.Println(i, number[i])
		if count%2 == 0 {
			number[i] = number[i] * 2

			if number[i] > 9 {
				number[i] = number[i] - 9
			}
		}
		sum += number[i]

	}

	return sum%10 == 0
}
