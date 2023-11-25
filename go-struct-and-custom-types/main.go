package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createBill() bill {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("create a new bill name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    b := newBill(name)
    fmt.Println("created the bill : ", b.name)
    return b
}

func main() {
   // mybill := newBill("mario's bill")
   // mybill.addItems("onion soup", 4.50)
   // mybill.addItems("veg soup", 5.50)
   // mybill.addItems("chicken soup", 9.50)
   // mybill.updateTip(10)
   // fmt.Println(mybill.format())

   mybill := createBill()
   fmt.Println(mybill)
}
