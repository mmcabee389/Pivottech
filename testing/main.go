package main

import (
	"fmt"
	"testing/calculator"
)

func main() {
	fmt.Printf("adding(1, 2) = %d\n)", calculator.Addition(1, 2))
	fmt.Printf("subtracting (10,5) = %d\n", calculator.Subtraction(10, 5))
	fmt.Printf("Multiplying (3, 4) = %d\n", calculator.Multiply(3, 4))

	r, err := calculator.Division(30, 10)
	fmt.Printf("Dividing(30, 10) = %d, %v\n", r, err)
}
