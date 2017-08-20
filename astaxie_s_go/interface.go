package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int8
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float64
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I'm %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("La la la ...", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I'm %s work at %s, Call me on %s\n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {

	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	// 定义men类型的变量i
	var i Men
	i = mike // i 能存储 Student
	fmt.Println("This is mike, a Student: ")
	i.SayHi()
	i.Sing("November rain...")

	i = tom // i 也能保存 Employee
	fmt.Println("This is tom, an Employee: ")
	i.SayHi()
	i.Sing("Born to be wild...")

	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike
	for _, val := range x {
		val.SayHi()
	}

}
