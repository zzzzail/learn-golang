/*
每个 Go 程序都是由包 package 构成的.
程序从 main 包开始运行.
*/
package main

/*
本程序通过导入 import 路径 "fmt" 和 "math/rand" 来使用这两个包.

按照约定, 包名与导入路径的最后一个元素一致.
例如 "math/rand" 包中的源码均以 package rand 语句开始.

同时, 若引用 "math/rand" 包.
按照约定, 当前引用文件中任意地方都可使用导入路径中的最后一个元素 rand 作为包的引用.
例如 rand.Intn(number).
*/
import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("My favorite number is", rand.Intn(100000))

}
