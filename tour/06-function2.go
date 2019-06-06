package main

import "fmt"

func main() {

	val := add(199, 298)
	fmt.Printf("value is %d", val)

}

/*
当连续两个或多个函数的已命名形参类型相同时, 除最后一个类型以外, 其它类型声明都可以省略.

在本例中
x int, y int
相当于
x, y int
*/
func add(x, y int) int {
	return x + y
}
