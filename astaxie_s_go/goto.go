package main

import (
	"fmt"
)

func main() {
	myFunc()
}

// 使用goto跳转到必须在当前函数内定义的标签
func myFunc() {
	i := 1
Here: // 标签, 大小写敏感
	fmt.Println(i)
	i++
	if i >= 100 {
		return
	}
	goto Here
}
