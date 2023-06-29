package main

import (
	"flag"
	"fmt"

	"github.com/SonuBardai/project_euler_go/problems"
)

func main() {
	var problem int
	var input int
	flag.IntVar(&problem, "problem", 1, "Problem number")
	flag.IntVar(&input, "input", 1, "Input value for the problem")
	flag.Parse()
	solution, error := problems.ProblemMapper(problem)
	if error != nil {
		panic(error)
	}
	final := solution(input)
	fmt.Println(final)
}
