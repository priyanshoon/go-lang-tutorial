package main

import "fmt"

type Shape interface {
	Area() float64
	Primeter() float64
}

type Car struct {
	Make   string
	Model  string
	Height int
	Width  int
}

type person struct {
	name  string
	hobby map[string]int
	age   int
}

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func concatInDiffWay(s1, s2 string) string {
	return s1 + s2
}

// function can have multipel return value
func getName() (string, string) {
	return "Priyanshu", "Yagini"
}

func main() {
	var smsSendingLimit int
	var costPerSms float64
	var hasPermission bool
	var username string

	fmt.Println("Variables Tutorial")

	// %v prints go syntax representation value
	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSms, hasPermission, username)

	congrats := "Happy Birthday"
	fmt.Println(congrats)

	pennisPerText := "Hello"
	// %T prints type of a passed variable
	fmt.Printf("The type of variable pennisPerText is %T type\n", pennisPerText)

	averageOpenRate, displayMessage := .23, "The rate has been opened"
	fmt.Printf("%v and %v\n", averageOpenRate, displayMessage)

	accountAge := 2.6
	accountAgeInt := int(accountAge)
	fmt.Println(accountAgeInt) // 2

	const premiumPlanName = "Premium Plan"
	fmt.Println(premiumPlanName)

	const firstName = "Priyanshu"
	const lastName = "Prasad Gupta"
	const fullName = firstName + " " + lastName
	fmt.Println(fullName)

	messageLen := 10
	maxMessageLen := 20
	fmt.Printf("Trying to send message of length: %v\n", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message Sent")
	} else {
		fmt.Println("Message not sent")
	}

	// this method is to keep the code consise
	// if INIT_VAR; CONDITION {
	//	something...
	// }
	if lenght := 10; lenght < 20 {
		fmt.Println("Email is valid")
	}

	fmt.Println("---------------- Functions starts here ----------------------")

	fmt.Println(concat("Hello", ", World!"))
	// _ means treating value as null (in short compiler remove that shit for us)
	x, _ := getName()
	fmt.Println(x)

	fmt.Println("---------------- Structs Area ------------------------------")

	p := person{
		name:  "Priyanshu",
		hobby: map[string]int{"Cricket": 3, "Football": 5},
		age:   30,
	}

	fmt.Println(p)
	fmt.Println(p.hobby["Cricket"])

	fmt.Println("------------------ Pointer Area ----------------------------")

	i := 42
	fmt.Println(&i)
}