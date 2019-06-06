package main

import (
	"fmt"
)

func main() {

	fmt.Println(factorial(3))

}

// factorial(3)
// 3 * factorial(2) == 3 * 2 = 6
// 2 * factorial(1) == 2 * 1 = 2
// 1 * factorial(0) == 1 * 1 = 1
func factorial(num int) int {

	if num == 0 {
		return 1
	}

	return num * factorial(num-1)

}
