package main

import (
	"fmt"
	"go_tooling/data"

	"google.golang.org/protobuf/proto"
)

type Direction int

const (
	_ Direction = iota
	North
	South
	East
	West
)

//go:generate stringer -type=Direction

//go:generate protoc -I=. --go_out=. --go_opt=module=go_tooling --go_opt=Mperson.proto=go_tooling/data person.proto

func main() {
	fmt.Println(North.String())

	p := &data.Person{
		Name:  "Mark",
		Id:    1,
		Email: "mark@mark.com",
	}

	fmt.Println(p)
	protoBytes, _ := proto.Marshal(p)
	fmt.Println(protoBytes)
	var p2 data.Person
	proto.Unmarshal(protoBytes, &p2)
	fmt.Println(&p2)

	fileEmbed()
	embedDir()
	printRights()
}
