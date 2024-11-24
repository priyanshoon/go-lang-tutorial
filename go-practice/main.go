package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("This is go practice")

	var visited byte = 'A'
	var arr [10][10]byte

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			arr[i][j] = '+'
		}
	}

	p1, p2 := rand.Int()%9, rand.Int()%9
	arr[p1][p2] = visited

	for visited != 'Z' {
		movement := rand.Int() % 4

		if movement == 0 && p2 > 0 && arr[p1][p2-1] == '+' {
			p2--
			visited++
			arr[p1][p2] = visited
		} else if movement == 1 && p2 < 9 && arr[p1][p2+1] == '+' {
			p2++
			visited++
			arr[p1][p2] = visited
		} else if movement == 2 && p1 > 0 && arr[p1-1][p2] == '+' {
			p1--
			visited++
			arr[p1][p2] = visited
		} else if movement == 3 && p1 < 9 && arr[p1+1][p2] == '+' {
			p1++
			visited++
			arr[p1][p2] = visited
		} else if arr[p1-1][p2] != '+' && arr[p1+1][p2] != '+' && arr[p1][p2-1] != '+' && arr[p1][p2+1] != '+' {
			fmt.Println("done")
			break
		} else if p2 == 0 && arr[p1][p2+1] != '+' && arr[p1-1][p2] != '+' && arr[p1+1][p2] != '+' {
			break
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%s ", string(arr[i][j]))
		}
		fmt.Println()
	}
	hello()
}
