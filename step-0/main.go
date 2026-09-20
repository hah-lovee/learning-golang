package main

import (
	"fmt"
)

func main() {
	var x int8 = 127
	x++            // wrap-around
	fmt.Println(x) // will be -128

}

func safeAdd(a, b int8) (int8, error) {
	sum := a + b

	if sum > 0 && a < 0 && b < 0 {
		return 0, fmt.Errorf("wrap-around")
	}
	if sum < 0 && a > 0 && b > 0 {
		return 0, fmt.Errorf("wrap-around")
	}

	return sum, nil

}
