package solutions

import (
	"testing"
)

func TestMultiples(t *testing.T) {
	testCases := []struct {
		name   string
		input  int
		output int
	}{
		{"test_1", 10, 23},
		{"test_2", 1000, 233168},
	}
	for _, testCase := range testCases {
		multiples := Multiples(testCase.input)
		if multiples != testCase.output {
			t.Errorf("expected %d, got %d", testCase.output, multiples)
		}
	}
}
