package calculator_test

import (
	"testing"
	"testing/calculator"
)

func TestAddition(t *testing.T) {

	a, b := 1, 2
	want := 3

	got := calculator.Addition(a, b)

	if got != want {

		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSubtraction(t *testing.T) {

	a, b := 5, 3
	want := 2

	got := calculator.Subtraction(a, b)

	if got != want {

		t.Errorf("got %d, want %d", got, want)
	}
}

func TestDivision(t *testing.T) {

	a, b := 4, 2
	want := 2

	got, _ := calculator.Division(a, b)

	if got != want {

		t.Errorf("got %d, want %d", got, want)
	}
}

func TestMultiply(t *testing.T) {

	a, b := 6, 5
	want := 30

	got := calculator.Multiply(a, b)

	if got != want {

		t.Errorf("got %d, want %d", got, want)
	}
}
