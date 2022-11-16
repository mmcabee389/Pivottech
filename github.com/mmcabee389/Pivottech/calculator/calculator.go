package calculator

import (
	"errors"
)

var ErrDividedByZero = errors.New("divide by Zero")

func Addition(a, b int) int {
	return a + b
}

func Subtraction(a, b int) int {

	return a - b
}

func Division(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDividedByZero
	}
	return a / b, nil
}

func Multiply(a, b int) int {

	return a * b
}
