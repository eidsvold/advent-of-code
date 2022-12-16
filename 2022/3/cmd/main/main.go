package main

import (
	"fmt"
	"os"

	"github.com/eidsvold/advent-of-code/2022/3/internal/puzzle"
)

func main() {
	first, err := puzzle.SolveFirstPuzzle()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(first)
}
