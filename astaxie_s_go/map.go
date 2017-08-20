package main

import (
	"fmt"
)

func main() {

	// 声明一个key是字符串, 值为int的字典
	// 这种方式的声明需要在使用之前使用make初始化
	var numbers = make(map[string]int)

	numbers["one"] = 1 // 赋值
	numbers["ten"] = 10
	numbers["three"] = 3

	fmt.Println("第三个数字是:", numbers["three"])

	// map 左边是key, 右边是值

	/**
	 * 使用map注意点
	 * 1. map 是无序的, 每次打印出来的map都会不一样, 它不能通过index获取, 必须通过key获取
	 * 2. map 的长度是不固定的, 也就是和slice一样, 也是一种引用类型
	 * 3. 内置len函数同样适用于map, 返回map拥有key的数量
	 * 4. map的值可以很方便的修改, 通过numbers["one"] = 11可以很容易的把key为one的字典值修改为11
	 * 5. map和其它基本类型不同, 他不是thread-safe, 在多个go-routine存取时,必须使用mutex lock机制
	 */

	// map的初始化可以通过key:val的方式初始化值, 同时map内置有判断是否存在key的方式
	// 通过delete删除map的元素

	rating := map[string]float64{"C": 5, "Go": 4.5, "Pythone": 4.5, "C++": 2}
	// map有两个返回值, 如果不存在key那么ok为false, 如果存在ok为true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("csharp is in the map and dsharp rating is", csharpRating)
	} else {
		fmt.Println("we hanve no rating associated with C# in the map")
	}

	delete(rating, "C") // 删除key为C的元素
	fmt.Println(rating)

	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut" // 现在m["Hello"] = Salut

}
