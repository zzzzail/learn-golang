package main

import "fmt"

/*
类型推导
在声明一个变量而不指定其类型时(即使不带类型的 := 语法或 var = 表达式语句),
变量的类型由右值推导得出.

当右值声明了类型时, 新变量的类型与其相同:
var i int
j := i // j 也是一个 int 值

不过当邮编包含未指明类型的数值常量时, 新变量的类型就可能是 int, float64 或 complex128 了,
这取决于常量的精度.
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128

*/
func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)

	i := 1<<32 - 1
	fmt.Printf("i is of type %T\n", i)

	s := "string"
	fmt.Printf("s is of type %T\n", s)

	f := 3.1415926
	fmt.Printf("f is of type %T\n", f)

	g := 0.778 + 0.75i
	fmt.Printf("g is of type %T\n", g)
}
