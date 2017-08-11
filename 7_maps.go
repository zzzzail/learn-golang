package main

import (
	"fmt"
)

func main() {

	presAge := make(map[string]int)

	presAge["TheodoreRoosevelt"] = 40
	// default : presAge["sth"] = 0
	// fmt.Println(presAge["sth"])
	fmt.Println(presAge["TheodoreRoosevelt"])
	fmt.Println(len(presAge))

	presAge["John F. Kennedy"] = 43
	fmt.Println(len(presAge))

	delete(presAge, "John F. Kennedy")

	fmt.Println(len(presAge), presAge)

}
