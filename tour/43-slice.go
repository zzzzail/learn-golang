package main

import "fmt"

/*
向切片追加元素
为切片追加新的元素是种常用的操作，为此 Go 提供了内建的 append 函数。内建函数的文档对此函数有详细的介绍。

func append(s []T, vs ...T) []T
append 的第一个参数 s 是一个元素类型为 T 的切片，其余类型为 T 的值将会追加到该切片的末尾。

append 的结果是一个包含原切片所有元素加上新添加元素的切片。

当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。
*/
func main() {

	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)

	/*
		1. 注意声明格式
		var arr [10]int // 声明数组
		var arr [10]int{1, 2, 3} // 声明数组并赋值
		var arr [10]int = [10]int{1, 2, 3} // ok
		var arr = [10]int{1, 2, 3} // ok
		arr := [10]int // 声明数组
		arr := [10]int{1, 2, 3} // 声明数组并赋值
		第一种是使用 var 关键字声明, 其参数类型紧跟在参数名称之后.
		第二种是使用 := 号声明, 参数类型在等号后, 值之前.

		2. 注意 array 和 slice 是不同类型!
		var b = [10]int{1, 2, 3}
		b = b[:] // cannot use b[:] (type []int) as type [10]int in assignment
		如果声明一个数组后, 把该数组值变成 slice 类型, 则会报错
		不能把类型 []int 作为类型 [10]int 来使用!
		啊喂, 不一样吗?
	*/
	var b = [10]int{1, 2, 3}
	b = b[:]
	b1 := b[:]
	printSlice(b1)

	// 怎么追加?
	b1 = append(b1, 10)
	printSlice(b1)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
