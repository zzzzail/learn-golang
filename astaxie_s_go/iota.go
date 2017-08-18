package main

import (
	"fmt"
)

const (
	x = iota // x=0
	y = iota // y=1
	z = iota // z=2

	/**
	 * 常量声明省略值时,默认和之前一个值的字面相同.
	 * 这里隐式地说w = iota, 因此w == 3.
	 * 其实上面y和z可同样不用"= iota"
	 */
	w
)

const v = iota // 每遇到一个const关键字, iota就会重置, 此时v=0

const (
	h, i, j = iota, iota, iota // h=0, i=0, j=0 iota 在同一行值相同
)

const (
	a       = iota // a=0
	b       = "B"
	c       = iota             // c=2
	d, e, f = iota, iota, iota // d=3, e=3, f=3
	g       = iota             // g=4
)

func main() {
	/**
	 * go 程序设计规则
	 * 大写字母开头的变量和方法是可导出的, 也就是其他包可以读取的, 是公有变量;
	 * 小写字母开头的就是不可导出的, 是私有变量.
	 *
	 */
	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
}
