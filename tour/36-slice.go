package main

import "fmt"

/*
切片就像数组的引用
切片并不存储任何数据, 它只是描述了底层数组中的一段.
更改切片的元素会修改其底层数组中对应的元素.
与它共享底层数组的切片都会观测到这些修改.
*/
func main() {

	names := [4]string{
		"John",
		"Sunny",
		"Bibi",
		"Ringo",
	}
	fmt.Println(names)

	_names := names[0:2]
	_names[0] = "New John"
	_names[1] = "New Sunny"
	fmt.Println(_names)

	fmt.Println(names)

}
