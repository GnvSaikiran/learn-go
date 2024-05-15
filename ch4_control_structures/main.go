package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Exercise-1:")
	s := exercise1()
	fmt.Println(s)
	fmt.Println("Exercise-2:")
	exercise2(s)
	fmt.Println("Exercise-3:")
	exercise3()
}

func exercise1() []int {
	s := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		s = append(s, rand.Intn(100))
	}

	return s
}

func exercise2(s []int) {
	for _, v := range s {
		switch {
		case v%6 == 0:
			fmt.Println(v, "Six!")
		case v%2 == 0:
			fmt.Println(v, "Two!")
		case v%3 == 0:
			fmt.Println(v, "Three!")
		default:
			fmt.Println(v, "Never mind")
		}
	}
}

func exercise3() {
	var total int
	for i := 0; i < 10; i++ {
		total = total + i
		fmt.Println(total)
	}
}
