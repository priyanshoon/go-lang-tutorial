package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func updateName(x *string) {
	*x = "sumo"
}

func main() {
	// name := "tifa"
	// //updateName(name)
	//
	// m := &name
	//
	// updateName(m)
	// fmt.Println(name)

	v := Vertex{100, 100}
	var p *Vertex = &v

	fmt.Println(&v.X, &v.Y)
	fmt.Printf("%p\n", &v)

	fmt.Println(p)

	//    fmt.Println(name)
}
