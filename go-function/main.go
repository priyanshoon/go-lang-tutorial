package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string)  {
    fmt.Printf("good morning %v\n", n)
}

func sayBye(n string)  {
    fmt.Printf("goodbye %v\n", n)
}

func cycleNames(n []string, f func(string))  {
    for _, v := range n {
        f(v)
    }
}

// example testing
func heroCycle() (int, int) {
    return 0, 3
}

func circleArea(r float64) float64 {
    return math.Pi * r * r
}

func main()  {
    // len() append()

    sayGreeting("priyanshu")
    sayBye("priyanshu")

    fmt.Println("---------------------------------------------------------------")

    cycleNames([]string{"priyanshu", "rounak", "krish"}, sayGreeting)

    fmt.Println("---------------------------------------------------------------")

    cycleNames([]string{"priyanshu", "rounak", "krish"}, sayBye)

    fmt.Println("---------------------------------------------------------------")

    a1 := circleArea(10.3)
    a2 := circleArea(15)
    fmt.Printf("the area of circle 1 is %v\n", a1)
    fmt.Printf("the area of circle 2 is %v\n", a2)

    fmt.Printf("circle : %0.3f", a1)
}
