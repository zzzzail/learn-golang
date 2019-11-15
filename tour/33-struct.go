package main

import "fmt"

/*
结构体文法
结构体文法通过直接列出字段的值来新分配一个结构体.
使用 Name: 语法可以仅列出部分字段. (字段名的顺序无关)
特殊的前缀 & 使用一个指向结构体的指针.
*/

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y: 0 被隐式赋予
	v3 = Vertex{Y: 2}  // X: 0 被隐式赋予
	v4 = Vertex{}      // X: 0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体(指针)
)

func main() {

	fmt.Println(v1, v2, v3, v4, p)

}
