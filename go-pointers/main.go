package main

import "fmt"

func updateName(x string) {
    x = "sumo"
}

func main() {
    name := "tifa"
    updateName(name)
    fmt.Println("memory address of name is : ", &name)

    m := &name
    fmt.Println("memory address of m: ", m)

    fmt.Println(name)
}
