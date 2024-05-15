package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

type Employee struct {
	FristName string `minStrlen:"3"`
	LastName  string
	Title     string `minStrlen:"5"`
	Age       int
}

func ValidateStringLength(v any) error {
	structType := reflect.TypeOf(v)
	if structType.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct, got %s", structType.Name())
	}

	var err error
	structValue := reflect.ValueOf(v)
	for i := range structType.NumField() {
		field := structType.Field(i)
		minLenStr := field.Tag.Get("minStrlen")
		if minLenStr == "" {
			continue
		}
		minLen, newErr := strconv.Atoi(minLenStr)
		if newErr != nil {
			err = errors.Join(err, fmt.Errorf("minStrlen should be a number %w", newErr))
		}
		value := structValue.Field(i).String()
		len := len(value)
		if len < minLen {
			err = errors.Join(err, fmt.Errorf(`length of "%s" field is less than %d`, field.Name, minLen))
		}
	}

	return err
}

type OrderInfo struct {
	IsReady     bool
	OrderNumber uint16
	OrderCode   rune
	Amount      int
	Items       []string
}

func OrderInfoSizeOffset(order OrderInfo) {
	fmt.Println(unsafe.Sizeof(order), unsafe.Offsetof(order.IsReady),
		unsafe.Offsetof(order.OrderNumber), unsafe.Offsetof(order.OrderCode),
		unsafe.Offsetof(order.Amount), unsafe.Offsetof(order.Items))
}

func main() {
	err := ValidateStringLength(Employee{"Mark", "Johnson", "Engineer", 22})
	if err != nil {
		fmt.Println(err)
	}

	order := OrderInfo{
		OrderCode:   'B',
		Amount:      2000,
		Items:       []string{"book", "pen"},
		OrderNumber: 123,
		IsReady:     true,
	}
	OrderInfoSizeOffset(order)
}
