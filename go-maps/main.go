package main

import "fmt"

func main() {
    menu := map[string]float64 {
        "soup": 4.99,
        "pie": 7.99,
        "salad": 6.99,
        "toffee pudding": 3.55,
    }

    fmt.Println(menu)
    fmt.Println(menu["pie"])

    // looping maps
    for k, v := range menu {
        fmt.Println(k, "-", v)
    }

    // ints as ket types
    phonebook := map[int]string{
        9987654320: "priyanshu",
        1234567889: "rounak",
        5647389200: "krish",
        9192837464: "ankit",
    }

    fmt.Println(phonebook)
    fmt.Println(phonebook[1234567890])

    phonebook[1234567890] = "rohit"
    fmt.Println(phonebook)
}
