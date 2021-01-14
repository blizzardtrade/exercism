package main

import (
	"errors"
)

//Square calculate how many grains were on a given square, and
func Square(input int) (uint64, error) {
	if input <= 0 || input > 64 {
		err := errors.New("mistake")
		return 0, err
	}

	var square uint64 = 1

	for i := 2; i <= input; i++ {
		square *= 2
	}

	return square, nil
}

//Total is the total number of grains on the chessboard
func Total() uint64 {
	var total uint64

	for i := 1; i <= 64; i++ {
		temp, _ := Square(i)
		total += temp
	}
	return total
}
