package main

import "fmt"

/*
切片文法
切片文法类似于没有长度的数组文法.
这是一个数组文法:
[3]bool{true, false, true}
下面这样则会创建一个和上面相同的数组, 然后创建一个引用了它的切片:
[]bool{true, true, false}
*/
func main() {

	q := []int{1, 2, 3, 4, 5}
	fmt.Println(q)

	r := []bool{true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{1, false},
		{2, true},
		{3, false},
	}
	fmt.Println(s)

}
