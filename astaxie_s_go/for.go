package main

import (
	"fmt"
)

func main() {

	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}

	fmt.Println("sum is equal to", sum)

	// 忽略 expression1, expression3
	sum2 := 1
	for sum2 < 1000 { // 相当于while
		sum2 += sum2
	}
	fmt.Println(sum2)

	// break continue
	for index := 10; index >= 0; index-- {
		if index == 5 {
			break // 跳出当前循环
		}
		fmt.Println("Break: ", index)
	}
	for index2 := 10; index2 >= 0; index2-- {
		if index2 == 5 {
			continue // 跳出本次循环
		}
		fmt.Println("continue", index2)
	}

	// for 配合 range 可以读取slice和map的数据
	m := map[string]int32{"One": 1, "Two": 2, "Three": 3, "Ten": 10}
	for k, v := range m {
		fmt.Println("map's Key:", k)
		fmt.Println("map's value:", v)
	}

}
