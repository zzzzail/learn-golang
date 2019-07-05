package main

import "fmt"

/*
defer 栈

推迟的函数调用会被压入一个栈中. 当外层函数返回时, 被推迟的函数会按照先进后出的顺序调用.

更多关于 defer 语句的信息，请阅读此博文。
http://blog.go-zh.org/defer-panic-and-recover
*/
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
