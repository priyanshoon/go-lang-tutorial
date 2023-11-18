package main

import "fmt"

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

func main()  {
    // len() append()

    sayGreeting("priyanshu")
    sayBye("priyanshu")

    fmt.Println("---------------------------------------------------------------")

    cycleNames([]string{"priyanshu", "rounak", "krish"}, sayGreeting)

    fmt.Println("---------------------------------------------------------------")

    cycleNames([]string{"priyanshu", "rounak", "krish"}, sayBye)
}
