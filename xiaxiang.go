package golang

func main() {

}

func level1() {
	x := 20
	level2()
}

func level2() {
	level3()
}

func level3() {
	// 这时候 l3 函数作用域中,
	// 是可以访问得到 父函数的父函数的 x 变量的扩展方法 x_align()

	// 换个方法 x_align() 提取出一个关键字 refer
	// x = refer x
	// 这样需要理解 refer 关键字.
	// 修改 level1 方法中的 x 变量时, 也可找到 level3 中的 x 变量引用.
}
