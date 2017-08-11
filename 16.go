package main

import (
	"fmt"
)

func main() {

	sPtr := new(int)

	changeYValNow(sPtr)

	fmt.Println("y =", *sPtr)

}

func changeYValNow(yPtr *int) {
	*yPtr = 1
}
