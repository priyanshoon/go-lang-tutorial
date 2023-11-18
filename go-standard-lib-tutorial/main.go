package main

import (
	"fmt"
	"sort"
	// "strings"
)

func main() {
    // greeting := "hello their friends"

    //    fmt.Println(strings.Contains(greeting, "hello!"))
    //    fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
    //    fmt.Println(strings.ToUpper(greeting))
    //    fmt.Println(strings.Index(greeting, "th"))
    //    fmt.Println(strings.Split(greeting, " "))

    //    the original value is unchanged
    //    fmt.Println("original string value = ", greeting)

    ages := []int{45, 29, 20, 34, 30, 75, 60, 50}
    sort.Ints(ages)
    fmt.Println(ages)

    index := sort.SearchInts(ages, 30)
    fmt.Println(index)

    names := []string{"priyanshu", "rounak", "rohit", "krish", "aryan"}
    sort.Strings(names)

    fmt.Println(names)

    fmt.Println(sort.SearchStrings(names, "priyanshu"))

}
