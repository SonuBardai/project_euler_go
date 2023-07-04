package utils

import (
	"math"
	"strconv"
)

func reverse(strInt string) string {
	runeString := []rune(strInt)
	for i, j := 0, len(runeString) - 1; i < j; i, j = i+1, j-1 {
		runeString[i], runeString[j] = runeString[j], runeString[i]
	}
	return string(runeString)
}

func IsPalindrome(input int) bool {
	strInt := strconv.Itoa(input)
	return strInt == reverse(strInt)
}

func Niner(input int) int {
	var nines int
	for i := 0; i < input; i ++ {
		nines += int(math.Pow10(i)) * 9
	}
	return nines
}
