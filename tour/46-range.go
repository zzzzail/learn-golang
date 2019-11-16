package main

import "fmt"

func main() {

	pic := Pic(10, 10)
	for i, v := range pic {
		fmt.Printf("i=%d, v=%v\n", i, v)
	}

}

// Pic 画图个图片哟
func Pic(dx, dy int) [][]uint8 {

	value := make([][]uint8, dy)

	for y := 0; y < dy; y++ {

		xarr := make([]uint8, dx)

		for x := 0; x < dx; x++ {
			// xarr[x] = uint8((x + y) / 2)
			// xarr[x] = uint8(x * y)
			xarr[x] = uint8(x ^ y)

		}

		value[y] = xarr
	}

	return value
}
