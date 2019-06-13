package main

import "fmt"

/*
变量的初始化
变量声明可以包含初始值, 每个变量对应一个初始值.
如果初始值已存在, 则可以生路类型; 变量会从初始值中获取类型.
*/

var i, j int = 1, 2

func main() {

	var c, pyhton, java = true, false, "no!"

	fmt.Println(c, pyhton, java, i, j)

}
