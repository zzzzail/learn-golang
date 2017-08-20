package main

import (
	"fmt"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func add1(i *int) int {
	*i = *i + 1
	return *i
}

func main() {
	x := 3
	y := 4
	z := 5
	maxXY := max(x, y)
	maxYZ := max(y, z)

	fmt.Printf("max(%d, %d) = %d\n", x, y, maxXY)
	fmt.Printf("max(%d. %d) = %d\n", y, z, maxYZ)

	// 指针调用
	i := 11
	i2 := add1(&i) // 传入函数i的地址
	fmt.Println(i)
	fmt.Println(i2)

	/**
	 * 传指针的好处
	 * 1. 传指针可以使得 多个函数操作同一个函数
	 * 2. 传指针比较轻量级(8bytes), 只传递内存地址, 我们可以用指针传递体积大的结构体.
	 * 如果用参数值传递的话, 每次在copy上会花费相对较多的系统开销(内存和时间).
	 * 所以当你要传递大的结构体时候, 用指针是一个明智的选择.
	 * 3. Go 语言中 channel, slice, map 这三种类型的实现机制类似指针, 所以可以直接传递, 而不用取地址后传递指针.
	 * (注: 若函数需要改变slice的长度时, 则仍需要取地址传递指针)
	 */

}
