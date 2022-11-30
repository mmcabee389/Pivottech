package calculator_test

import (
	"github.com/mmcabee389/Pivottech/calculator"
	"testing"
)

func TestAddition(t *testing.T) {
	cases := []struct {
		problem string
		num1    int
		num2    int
		answer  int
	}{
		{
			problem: "1+2",
			num1:    1,
			num2:    2,
			answer:  3,
		},
		{
			problem: "4+2",
			num1:    4,
			num2:    2,
			answer:  6,
		},
	}
	for _, ta := range cases {
		t.Run(ta.problem, func(t *testing.T) {
			result := calculator.Addition(ta.num1, ta.num2)
			if result != ta.answer {
				t.Errorf("expected %d, but got %d", ta.answer, result)

			}
		})
	}
}

func TestSubtraction(t *testing.T) {
	cases := []struct {
		problem string
		num1    int
		num2    int
		answer  int
	}{
		{
			problem: "5-2",
			num1:    5,
			num2:    2,
			answer:  3,
		},
		{
			problem: "4-2",
			num1:    4,
			num2:    2,
			answer:  2,
		},
	}
	for _, ts := range cases {
		t.Run(ts.problem, func(t *testing.T) {
			result := calculator.Subtraction(ts.num1, ts.num2)
			if result != ts.answer {
				t.Errorf("expected %d, but got %d", ts.answer, result)
			}
		})
	}
}

func TestDivision(t *testing.T) {
	cases := []struct {
		problem string
		num1    int
		num2    int
		answer  int
	}{
		{
			problem: "8/4",
			num1:    8,
			num2:    4,
			answer:  2,
		},
		{
			problem: "4/2",
			num1:    4,
			num2:    2,
			answer:  2,
		},
	}
	for _, td := range cases {
		t.Run(td.problem, func(t *testing.T) {
			result, _ := calculator.Division(td.num1, td.num2)
			if result != td.answer {
				t.Errorf("expected %d, but got %d", td.answer, result)
			}
		})
	}
}
func TestMultiply(t *testing.T) {
	cases := []struct {
		problem string
		num1    int
		num2    int
		answer  int
	}{
		{
			problem: "1*2",
			num1:    1,
			num2:    2,
			answer:  2,
		},
		{
			problem: "4*2",
			num1:    4,
			num2:    2,
			answer:  8,
		},
	}
	for _, tm := range cases {
		t.Run(tm.problem, func(t *testing.T) {
			result := calculator.Multiply(tm.num1, tm.num2)
			if result != tm.answer {
				t.Errorf("expected %d, but got %d", tm.answer, result)

			}
		})
	}
}
