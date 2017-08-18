package main

import "fmt"

func main() {

	// 修改字符串
	s := "Hello"
	c := []byte(s)
	c[0] = 'h'
	s2 := string(c)
	fmt.Println("Value is", s, c, s2)

	// 字符串切片操作, 切片操作是截取字符串之后再拼接
	str := "Hello"
	str = "c" + str[2:2]
	fmt.Println(str)

	// 多行字符串, 换行也原样输出
	str2 := `
    <div>辣鸡任舸</div>
  `
	fmt.Println(str2)

	// rune是int32的别称, byte是uint8的别称.
	var num byte = 12
	var num2 rune = 13
	fmt.Println(num, num2)

}
