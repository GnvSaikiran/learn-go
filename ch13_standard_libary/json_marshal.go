package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func marshalFunction() {
	data := []byte(`{"name": "Mark", "age":20}`)
	var p Person
	err := json.Unmarshal(data, &p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(p)

	data, err = json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(data))
}
