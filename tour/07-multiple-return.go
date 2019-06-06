package main

import "fmt"

func main() {

	a := "World!"
	b := "Hello"

	a, b = swap(a, b)

	fmt.Printf("%s %s", a, b)

}

/*
函数可以返回任意数量的返回值.
swap 函数返回了两个字符串.
*/
func swap(a, b string) (string, string) {
	return b, a
}
