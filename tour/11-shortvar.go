package main

import "fmt"

/*
短变量声明
在函数中, 简洁赋值语句 := 可在类型明确的地方替代 var 声明.
函数外的每个语句都必须以关键字开始 (var, func 等), 因此 := 结构不能在函数外使用.
*/

// 报错: syntax error: non-declaration statement outside function body
// java := "Yes!"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python := true, false

	fmt.Println(i, j, k, c, python)
}
