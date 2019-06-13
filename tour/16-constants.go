package main

import "fmt"

/*
常量

常量的声明与变量类似, 只不过是使用 const 关键字.
常量可以是字符、字符串、布尔值或数值.
常量不能用 := 语法声明.
*/

const Pi float64 = 3.14

func main() {

	const World string = "World!"
	fmt.Println("Hello ", World)

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Printf("Pi is type of %T\n", Pi)

}
