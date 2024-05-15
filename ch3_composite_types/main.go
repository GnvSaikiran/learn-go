package main

import "fmt"

func main() {
	fmt.Println("Exercise-1:")
	exercise1()
	fmt.Println("Exercise-2:")
	exercise2()
	fmt.Println("Exercise-3:")
	exercise3()
}

func exercise1() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	first := greetings[:2]
	second := greetings[1:4]
	third := greetings[3:]
	fmt.Println(greetings, first, second, third)
}

func exercise2() {
	message := "Hi 🧑 and 👧"
	fmt.Println(message[3:8])
}

func exercise3() {
	type Employee struct {
		firstName, lastName string
		id                  int
	}

	a := Employee{"Mark", "Rober", 1}
	b := Employee{firstName: "Jenna", lastName: "Fischer", id: 2}
	var c Employee
	c.firstName = "Jim"
	c.lastName = "Scott"
	c.id = 3

	fmt.Println(a, b, c)
}
