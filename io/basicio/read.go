package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	sampleReadFile()
}

func sampleReadFromString() {
	r := strings.NewReader("from string")
	data, _ := ReadFrom(r, 10)

	fmt.Println(data)
}

func sampleReadStdin() {
	fmt.Print("Pelase input from stdin:")

	data, _ := ReadFrom(os.Stdin, 10)
	fmt.Println(data)
}

func sampleReadFile() {
	file, _ := os.Open("main.go")
	defer file.Close()
	data, _ := ReadFrom(file, 10)

	fmt.Println(string(data))
}

// ReadFrom 读取文件
func ReadFrom(reader io.Reader, num int) ([]byte, error) {

	p := make([]byte, num)

	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}

	return p, err
}
