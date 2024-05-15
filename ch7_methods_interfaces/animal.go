package main

import "fmt"

type Animal interface {
	makeSound()
}

type Dog string

func (d Dog) makeSound() {
	fmt.Println("bark bark")
}

func (d Dog) description() {
	fmt.Println("It has four legs")
}

func temp(a Animal) {
	a.makeSound()
	a.(Dog).description()
}

func animalFunc() {
	d := Dog("tom")
	fmt.Println(d)
	temp(d)
}
