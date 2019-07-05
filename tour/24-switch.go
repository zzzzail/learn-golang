package main

import (
	"fmt"
	"runtime"
)

/*
switch
switch 是编写一连串 if - else 语句的简便方法. 它运行第一个值等于条件表达式的 case 语句.

Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的 switch, 不过 Go 只运行
选定的 case, 而非之后所有的 case. 实际上, Go 自动提供了在这些语言中每个 case 后面所需要的,
不同在于 switch 的 case 无需为常量, 且取值不必为整数.
*/
func main() {

	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd
		// plan9, windows...
		fmt.Print("%s.\n", os)
	}

}
