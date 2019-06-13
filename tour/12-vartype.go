package main

import (
	"fmt"
	"math/cmplx"
)

/*
Go 的基本类型有

bool
string
int  int8  int16  int32  int64
unit uint8 uint16 uint32 uint64

byte // uint8 的别名
rune // int32 的别名
	 // 表示一个 Unicode 码点

float32 float64

complex64 complex128

int, uint 和 uintptr 在 32 位系统上通常为 32 位宽, 在 64 位系统上则为 64 位宽.
当你需要一个整数值时应使用 int 类型, 除非你有特殊的理由使用固定大小或无符号的整数类型.

------------------------------------------------------------------------
int 类型是整数, 可以为负数
uint 类型是正整数 unsigned int, 无符号的整数
*/

var (
	ToBe   bool       = false
	i      int64      = -2
	ui     uint64     = 0
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	fmt.Printf("Type %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type %T Value: %v\n", i, i)
	fmt.Printf("Type %T Value: %v\n", ui, ui)
	fmt.Printf("Type %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T Value: %v\n", z, z)

}
