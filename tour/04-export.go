package main

import (
	"fmt"
	"math"
)

func main() {

	/*
		关于导出名:

		在 Go 文件中, 如果一个名字首字母为大写开头, 那么它就是已导出的.
		例如, Pizza 就是一个已导出名, Pi 也同样, 它导出字 math 包.

		在导出一个包时, 你只能引用其已导出的名字. 任何"未导出"的名字在该包外均无法访问.

		若使用 math.pi, 则会报 undefined: math 的错误.
	*/

	// 错误的示范
	// fmt.Printf("Pi = %g\n", math.pi)
	fmt.Printf("Pi = %g\n", math.Pi)

}
