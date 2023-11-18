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
    fmt.Printf("the area of circle is %v\n", a1)

    fmt.Printf("circle : %0.3f", a1)
}
