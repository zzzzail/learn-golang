package main

import "fmt"

/*
零值
没有明确初始值的变量声明会被赋予它们的零值.
零值是:
	数值类型为 0
	布尔类型为 false
	字符串为 "" (空字符串)
*/
func main() {

	var flag bool
	var i int
	var str string

	fmt.Println(flag)
	fmt.Println(i)
	fmt.Println(str)

}
