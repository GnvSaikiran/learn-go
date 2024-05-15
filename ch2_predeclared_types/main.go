package main

import "fmt"

func main() {
	fmt.Println("Excercise-1:")
	exercise1()
	fmt.Println("Excercise-2:")
	exercise2()
	fmt.Println("Excercise-3:")
	exercise3()
}

func exercise1() {
	i := 20
	var f = float64(i)
	fmt.Println(i, f)
}

func exercise2() {
	const value = 10
	i := value
	var f float64 = value
	fmt.Println(i, f)
}

func exercise3() {
	var b byte = 255
	var smallI int32 = 2_147_483_647
	var bigI uint64 = 18_446_744_073_709_551_615

	fmt.Println(b, smallI, bigI)

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b, smallI, bigI)
}
