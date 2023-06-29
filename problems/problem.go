package problems

import (
	"fmt"

	"github.com/SonuBardai/project_euler_go/solutions"
)

func ProblemMapper(problem_number int) (func(input int) int, error) {
	switch problem_number {
	case 1:
		return solutions.Multiples, nil
	default:
		return nil, fmt.Errorf("no matching problem found")
	}
}
