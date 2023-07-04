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

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		name   string
		input  int
		output int
	}{
		{"test_1", 50, 44},
		{"test_2", 4_000_000, 4613732},
	}
	for _, testCase := range testCases {
		multiples := Fibonacci(testCase.input)
		if multiples != testCase.output {
			t.Errorf("expected %d, got %d", testCase.output, multiples)
		}
	}
}

func TestPrimeFactors(t *testing.T) {
	testCases := []struct {
		name   string
		input  int
		output int
	}{
		{"test_1",13195 , 29},
		{"test_2", 600851475143, 6857},
	}
	for _, testCase := range testCases {
		multiples := PrimeFactor(testCase.input)
		if multiples != testCase.output {
			t.Errorf("expected %d, got %d", testCase.output, multiples)
		}
	}
}


func TestPalindromeProduct(t *testing.T) {
	testCases := []struct {
		name   string
		input  int
		output int
	}{
		{"test_1", 2, 9009},
		{"test_2", 3, 906609},
	}
	for _, testCase := range testCases {
		multiples := PalindromeProduct(testCase.input)
		if multiples != testCase.output {
			t.Errorf("expected %d, got %d", testCase.output, multiples)
		}
	}
}