package main

import "fmt"

type Person struct {
	firstName, lastName string
	age                 int
}

func main() {
	// exercise-1
	p1 := MakePerson("Mark", "Robert", 32)
	p2 := MakePersonPointer("John", "Targer", 20)
	fmt.Println(p1)
	fmt.Println(p2)

	// exercise-2
	s := []string{"Apple", "Guava", "Pineapple"}
	fmt.Printf("Before UpdateSlice: %v\n", s)
	UpdateSlice(s, "Orange")
	fmt.Printf("After UpdateSlice: %v\n", s)

	fmt.Printf("\nBefore GrowSlice: %v\n", s)
	GrowSlice(s, "Papaya")
	fmt.Printf("After Growslice: %v\n", s)
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func UpdateSlice(s []string, v string) {
	s[len(s)-1] = v
	fmt.Printf("In UpdateSlice: %v\n", s)
}

func GrowSlice(s []string, v string) {
	s = append(s, v)
	fmt.Printf("In GrowSlice: %v\n", s)
}
