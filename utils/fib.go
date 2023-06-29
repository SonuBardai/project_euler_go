package utils

import "fmt"

func fibRecursion(input int) int {
	if input <= 1 {
		return input
	}
	return fibRecursion(input-1) + fibRecursion(input-2)
}

func FibSeriesRecursive(input int) []int {
	var i int
	var fib_series []int
	for {
		cur_fib := fibRecursion(i)
		if cur_fib >= input {
			break
		}
		fib_series = append(fib_series, cur_fib)
		i += 1
	}
	return fib_series
}

func FibSeries(input int) []int {
	var fib_series = []int{1, 1}
	for fib_series[len(fib_series) - 1] < input {
		temp := fib_series[len(fib_series) - 1] + fib_series[len(fib_series) - 2]
		fib_series = append(fib_series, temp)
	}
	return fib_series
}

func FibSum(fibSeries []int) int {
	var sum int
	fmt.Println(fibSeries)
	for _, i := range fibSeries {
		if i%2 == 0 {
			sum += i
		}
	}
	return sum
}
