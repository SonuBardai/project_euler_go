package solutions

func Multiples(input int) int {
	var total int
	for i := 1; i < input; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return total
}
