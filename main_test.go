package main

import "testing"

type testCase struct {
	input  string
	output interface{}
}

func TestInt(t *testing.T) {
	t.Skip()

	tests := []testCase{
		{"0", 0},
		{"42", 42},
	}

	runTests(t, tests)
}

func TestAddSub(t *testing.T) {
	t.Skip()

	tests := []testCase{
		{"5 + 20 - 4", 21},
		{"12 + 34 - 5", 41},
		{"  1  + 2 -3  +   4-  5  ", -1},
	}

	runTests(t, tests)
}

func TestMulDiv(t *testing.T) {
	t.Skip()

	tests := []testCase{
		{"5 + 6 * 7", 47},
		{"10 - 4 / 2", 8},
		{"6 + 5 * 4 - 9 / 3", 23},
	}

	runTests(t, tests)
}

func TestParen(t *testing.T) {
	t.Skip()

	tests := []testCase{
		{"5 * (9 - 6)", 15},
		{"(3 + 5) / 2", 4},
		{"1 + 9 * (13 - (8 + (6 - 3) * 4) / 5) / 3", 28},
	}

	runTests(t, tests)
}

func TestUnary(t *testing.T) {
	t.Skip()

	tests := []testCase{
		{"-10 + 20", 10},
		{"- -10", 10},
		{"- - +10", 10},
	}

	runTests(t, tests)
}

func TestComplexArithmetic(t *testing.T) {
	t.Skip()

	tests := []testCase{
		{"-5 * (-2) * (-3) + (-26) / (-2) - 12 * (-3)", 19},
		{"(8 - 5) * (-2) + - - 1 + (-9 + 8) * (5 - 7) - 12 / (22 - 16) * (-3) - - - 2", 1},
	}

	runTests(t, tests)
}

func runTests(t *testing.T, tests []testCase) {
	for _, tt := range tests {
		output, err := calc(tt.input)
		if err != nil {
			t.Errorf("%q => got unexpected error: %v", tt.input, err)
		}
		if output != tt.output {
			t.Errorf("%q => expected %v, but got %v", tt.input, tt.output, output)
		}
	}
}
