package main

import "fmt"

func updateName(x *string) {
    *x = "sumo"
}

func main() {
    name := "tifa"
    //updateName(name)

    m := &name

    updateName(m)
    fmt.Println(name)

//    fmt.Println(name)
}
