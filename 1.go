package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello World!")

	// uint8 : unsigned 8-bit integers (0 to 255)
	// uint16 : unsigned 16-bit integers (0 to 65535)
	// uint32 : unsigned 32-bit integers (0 to 4294967295)
	// uint64 : unsigned 64-bit integers (0 to 18446744073709551615)
	// int8 : unsignded 8-bit integers (-128 to 127)
	// int16 : unsigned 16-bit integers (-32768 to 32767)
	// int32 : unsigned 32-bit integers (-2147483648 to 2147483647)
	// int64 : unsigned 63-bit integers (-9223372036854775808 to 9223372036854775807)

	var age int = 40

	var favNum float64 = 1.6180339

	fmt.Println(age, favNum)

	var numOne = 1.000
	var num99 = .999

	fmt.Println(numOne - num99)

	fmt.Println("6 + 4 =", 6+4)
	fmt.Println("6 - 4 =", 6-4)
	fmt.Println("6 * 4 =", 6*4)
	fmt.Println("6 / 4 =", 6/4)
	fmt.Println("6 % 4 =", 6%4)

	const pi float64 = 3.14159265

	// var (
	// 	varA = 2
	// 	varB = 3
	// )

	var myName string = "Derek Banas"

	fmt.Println(len(myName))
	fmt.Println(myName + " is a robot")
	fmt.Println("i like \n \n")
	fmt.Println("Newlines")

	var isOver40 bool = true
	fmt.Println("bool", isOver40)
	fmt.Printf("%t \n", isOver40)

	fmt.Printf("%.3f \n", pi)
	fmt.Printf("%T \n", pi)

	fmt.Printf("%d \n", 100)
	fmt.Printf("%b \n", 100)
	fmt.Printf("%c \n", 44)
	fmt.Printf("%x \n", 17)
	fmt.Printf("%e \n", pi)

	fmt.Println("true && false =", true && false)
	fmt.Println("true || false =", true || false)
	fmt.Println("!true =", !true)

}
