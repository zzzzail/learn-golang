package main

import "fmt"

func main() {

	val := add(1, 2)
	fmt.Printf("value = %d", val)

}

/*
	func 关键字用于声明一个函数.
	格式为:
	func function_name(args...) {
		running...
	}

	函数可以没有参数, 或多个参数.
	在本例中, add 接受两个 int 类型的参数 x 和 y.
	( 注意类型在变量名之后! )

	@see http://blog.go-zh.org/gos-declaration-syntax
*/
func add(x int, y int) int {
	return x + y
}
