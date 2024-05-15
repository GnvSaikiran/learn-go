package main

import (
	"fmt"

	"github.com/GnvSaikiran/hello"
	h "github.com/GnvSaikiran/hello/v2"
)

func main() {
	fmt.Println("Packages and Modules")
	h.GreetUser("Mark")
	hello.GreetUser()
}
