package main

import "fmt"

func main() {
    var ages [3]int = [3]int{20, 19, 23}
    fmt.Println(ages, len(ages))

    names := [4]string{"Joe", "Love", "Beck", "Tuco"}
    names[1] = "Priyanshu"
    fmt.Println(names, len(names))

    // slices (use arrays under the hood)
    var scores = []int{100, 50, 60}
    scores[2] = 25
    scores = append(scores, 85)
    fmt.Println(scores, len(scores))

    // slices range
    rangeOne := names[1:3]
    rangeTwo := names[2:]
    rangeThree := names[:3]

    fmt.Println(rangeOne, rangeTwo, rangeThree)

    rangeOne = append(rangeOne, "kooper")

    fmt.Println(rangeOne)
}
