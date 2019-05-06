package main

import (
	"fmt"
)

// type error interface {
// 	Error() string
// }

type ErrZeroDivision struct {
	message string
}

func NewErrZeroDivision(message string) *ErrZeroDivision {
	return &ErrZeroDivision{
		message: message,
	}
}

func (e *ErrZeroDivision) Error() string {
	return e.message
}

func main() {

	result, err := divide(1.0, 0.0)
	if err != nil {
		switch err.(type) {
		case *ErrZeroDivision:
			fmt.Println("here: ", err.Error())
		default:
			fmt.Println("What the h* just happened?")
		}
	}
	fmt.Println(result)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, NewErrZeroDivision("Cannot divide by zero")
	}
	return a / b, nil
}
