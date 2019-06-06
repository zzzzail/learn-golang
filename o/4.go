package main

import "fmt"

func main() {

	var myAge int = 16

	switch myAge {
	case 16:
		fmt.Println("Go Drive")
	case 18:
		fmt.Println("Go Vote")
	default:
		fmt.Println("You Have Fun")
	}

}
