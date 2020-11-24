package main

import "testing"

type testCase struct {
	input  string
	output interface{}
}

func TestDummy(t *testing.T) {
	tests := []testCase{
		{"", nil},
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
