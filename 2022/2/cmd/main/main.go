package main

import (
	"advent-of-code/2022/2/internal/puzzle"

	"fmt"
	"os"
)

func main() {
	first, err := puzzle.SolveFirstPuzzle()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(first)

	second, err := puzzle.SolveSecondPuzzle()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(second)
}
