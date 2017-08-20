package main

import (
	"fmt"
)

func main() {

	// 动态数组
	// slice 并不是真正意义上的动态数组, 而是一个引用类型.
	// slice 总是指向一个底层 array, slice 声明也可以像 array 一样, 只会不需要长度.
	// var fslice []int
	// slice := []byte{'a', 'b', 'c', 'd'}

	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// slice 可以从一个数组或一个已存在的slice中再次声明.
	// a 指向两个含有byte的slice
	var a, b []byte
	// a指向数组的第3个元素开始, 并到底5个元素结束
	// 现在a包含 a[2], a[3], a[4]
	a = ar[2:5]

	// a[3], a[4]
	b = ar[3:5]

	// slice 是引用类型, 所以当引用改变其中元素的值时, 其它的所有引用都会改变该值.

	// slice 的内置函数
	fmt.Println(len(a), len(b), len(ar))
	fmt.Println(cap(a), cap(b), cap(ar))

}
