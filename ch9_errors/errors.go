package main

import (
	"errors"
)

var ErrInvalidID = errors.New("id is invalid")

var employees = []Employee{
	{id: 1, firstName: "Mark", lastName: "Rober", age: 20},
	{id: 2, lastName: "Rober", age: 20},
	{id: 3, firstName: "Mark", age: 20},
	{id: 4, firstName: "Mark", lastName: "Rober"},
}

func getUser(id int) error {
	for _, emp := range employees {
		if emp.id == id {
			switch {
			case emp.firstName == "":
				return EmptyFieldErr{"firstname"}
			case emp.lastName == "":
				return EmptyFieldErr{"lastname"}
			case emp.age == 0:
				return EmptyFieldErr{"age"}
			default:
				return nil
			}
		}
	}

	return ErrInvalidID
}

type EmptyFieldErr struct {
	name string
}

func (e EmptyFieldErr) Error() string {
	return e.name + " is empty"
}

type Employee struct {
	id                  int
	firstName, lastName string
	age                 int
}

func TestCustomErrorType() {

}
