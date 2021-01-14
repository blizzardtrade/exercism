package main

func Reverse2(input string) string {

	str := []rune(input)
	for i, j := 0, len(str)-1; i < len(str)/2; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}
