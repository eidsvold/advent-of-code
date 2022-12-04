package puzzle

import (
	"advent-of-code/2022/2/pkg/utils"

	"fmt"
	"strings"
)

type Yo struct {
}

func SolveSecondPuzzle() (int, error) {
	scanner, err := utils.NewScanner("assets/input.txt")

	if err != nil {
		return 0, err
	}

	sum := 0

	for scanner.Scan() {
		slice := strings.Split(scanner.Text(), " ")

		if len(slice) != 2 {
			return 0, fmt.Errorf("slice is not of length 2: length %d", len(slice))
		}

		opponent, expected := slice[0], slice[1]

		if expected == "X" {
			sum += 0
		}

		pts, err := playHand(slice[0], slice[1])

		if err != nil {
			return 0, err
		}

		sum += pts
	}
}
