package main

import "fmt"

func updateName(x string) string {
    x = "sumo"
    return x
}

func updateMenu(y map[string]float64) {
    y["coffee"] = 2.99

}

func main()  {
    // group A type => string, ints, bools, floats, arrays, structs
    name := "tifa"
    name = updateName(name)
    fmt.Println(name)

    // group b types => slices, maps, functions
    menu := map[string] float64 {
        "pie": 5.95,
        "ice cream": 3.99,
    }

    updateMenu(menu)
    fmt.Println(menu)

}
