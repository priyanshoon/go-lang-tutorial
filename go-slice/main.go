package main

import (
	"fmt"
)

func main() {
	fmt.Println("GO SLICE")

	b := make([]int, 3)

	fmt.Println(b)
	fmt.Println(len(b), cap(b))

	b = append(b, 10)

	fmt.Println(b)
	fmt.Println(len(b), cap(b))

	b = append(b, 45)
	fmt.Println(b)
	fmt.Println(len(b), cap(b))

	b = append(b, 56)
	fmt.Println(b)
	fmt.Println(len(b), cap(b))

	b = append(b, 100)
	fmt.Println(b)
	fmt.Println(len(b), cap(b))
}
