package main

import (
	"errors"
	"fmt"
)

func main() {
	userIds := []int{1, 2, 5}
	for _, userId := range userIds {
		var emptyFieldErr EmptyFieldErr
		err := getUser(userId)
		if err != nil {
			if errors.Is(err, ErrInvalidID) {
				fmt.Println(err)
			} else if errors.As(err, &emptyFieldErr) {
				fmt.Println(emptyFieldErr)
			} else {
				fmt.Println(err)
			}
		}

	}

}
