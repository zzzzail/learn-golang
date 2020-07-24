package main

import (
	"fmt"
	"os"
)

/*
main
*/
func main() {

	// 拼接字符串
	str := fmt.Sprintf("float %f", 3.1415926)
	fmt.Println(str)

	// F
	fmt.Fprintln(os.Stdout, "write to os.Stdout text.")

	d := Data{1, 2}
	fmt.Println(d)
}

type Data struct {
	X, Y int
}

func (d Data) String() string {
	return fmt.Sprintf("Data %v", d)
}
