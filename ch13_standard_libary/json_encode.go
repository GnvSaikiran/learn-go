package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func encodeFunction() {
	toFile := Person{
		Name: "mark",
		Age:  22,
	}

	file, err := os.Create("person.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(toFile)
	if err != nil {
		panic(err)
	}

	var p Person
	file2, err := os.Open("person.json")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	err = json.NewDecoder(file2).Decode(&p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)
}
