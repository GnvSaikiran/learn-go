package main

import (
	"fmt"
	"strconv"
)

type Number interface {
	int | float64
}

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type intPrintable int

func (i intPrintable) String() string {
	return "Value is: " + strconv.FormatInt(int64(i), 10)
}

func printPrintable[T Printable](p T) {
	fmt.Println(p)
}

func double[T Number](value T) T {
	return value * 2
}

func main() {
	fmt.Println(double(2))
	fmt.Println(double(2.5))

	var i intPrintable = 3
	printPrintable(i)

	head := Node[int]{}
	head.Add(3)
	head.Add(4)
	head.PrintValues()
	head.Insert(6, 1)
	head.PrintValues()
	fmt.Println(head.Index(4))
	fmt.Println(head.Index(5))
}
