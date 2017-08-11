package main

import (
	"fmt"
)

func main() {

	// rect1 := Rectangle{letfX: 0, topY: 1, width: 20, height: 30}
	rect1 := Rectangle{0, 1, 20, 30}
	fmt.Println(rect1)
	fmt.Println("Rectangel is", rect1.width, "wide")
	fmt.Println("Area of rectangle =", rect1.area())

}

type Rectangle struct {
	letfX  float64
	topY   float64
	width  float64
	height float64
}

func (rect *Rectangle) area() float64 {

	return rect.width * rect.height

}
