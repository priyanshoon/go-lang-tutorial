package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ans := make([][]uint8, dy)
	for i := range ans {
		ans[i] = make([]uint8, dx)
		for j := range ans[i] {
			ans[i][j] = uint8((i ^ j) * (i ^ j))
		}
	}
	return ans
}

func main() {
	pic.Show(Pic)
}
