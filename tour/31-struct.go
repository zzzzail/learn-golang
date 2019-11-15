package main

import "fmt"

/*
结构体字段
结构体字段使用点号来访问.
*/

type Vertex struct {
	X int
	Y int
}

func main() {

	v := Vertex{1, 2}
	v.X = 40
	fmt.Println(v)

}
