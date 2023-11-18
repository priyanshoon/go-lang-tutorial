package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
    s := strings.ToUpper(n)
    names := strings.Split(s, " ")
}

func main()  {
    getInitials("tifa lockhart")

}
