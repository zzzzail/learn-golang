package main

import (
	"fmt"
)

func main() {
	i := 10
	// 每一个case中默认有一个break, 匹配成功后不会继续向下执行其它case, 而是跳出这个switch
	// 可以使用fallthrough
	switch i {
	case 1:
		fmt.Println("i is equal to", 1)
	case 2:
		fmt.Println("i is equal to", 2)
	case 3:
		fmt.Println("i is equal to", 3)
	case 5:
		fmt.Println("i is equal to", 5)
	case 10:
		fmt.Println("i is equal to", 10)
		fallthrough
	default:
		fmt.Println("All I know is that i is an integer")
	}

}
