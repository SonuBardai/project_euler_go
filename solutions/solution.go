package solutions

import (
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
