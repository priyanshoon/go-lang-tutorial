package main

import "fmt"

func main() {
    //    x := 0

    //    for x < 5 {
    //        fmt.Println("the value of x is : ", x);
    //        x++
    //    }

    //for i := 0; i < 5; i++ {
    //    fmt.Println("the value of i is : ", i);
    //}

    names := []string{"priyanshu", "rounak", "krish", "rohit", "aryan"}

    for i := 0; i < len(names); i++ {
        fmt.Println(names[i])
    }

}
