package solutions

import (
	"math"

	"github.com/SonuBardai/project_euler_go/utils"
)

func Multiples(input int) int {
	var total int
	for i := 1; i < input; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return total
}

func Fibonacci(input int) int {
	// fibSeries := utils.FibSeriesRecursive(input) // 1, 2, ... n; n < input
	fibSeries := utils.FibSeries(input)
	sum := utils.FibSum(fibSeries)
	return sum
}

func primeNumber(input int) bool {
	for i := 2; i < input; i++ {
		if input%i == 0 {
			return false
		}
	}
	return true
}

func PrimeFactor(input int) int {
	x := input
	for x%2 == 0 {
		x /= 2
	}
	for i := int(math.Sqrt(float64(x))); i > 2; i-- {
		if i%2 == 0 {
			continue
		}
		if input%i == 0 && primeNumber(i) {
			return i
		}
	}
	return 1
}
