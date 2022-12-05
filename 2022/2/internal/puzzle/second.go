package puzzle

import (
	"advent-of-code/2022/2/pkg/utils"

	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

var hands = []string{"A", "B", "C"}
var handCount = len(hands)

func calculateExpectedPts(opponent string, outcome string) (int, error) {
	handIdx := slices.IndexFunc(hands, func(hand string) bool { return opponent == hand })

	if handIdx == -1 {
		return 0, fmt.Errorf("unknown opponent hand: %s", opponent)
	}

	switch outcome {
	case "X":
		pts := (handIdx+handCount-1)%handCount + 1
		return pts, nil
	case "Y":
		pts := handIdx + 1
		return pts + 3, nil
	case "Z":
		pts := (handIdx+1)%handCount + 1
		return pts + 6, nil
	}

	return 0, fmt.Errorf("unknown expected outcome: %s", outcome)
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

		opponent, outcome := slice[0], slice[1]

		pts, err := calculateExpectedPts(opponent, outcome)

		if err != nil {
			return 0, err
		}

		sum += pts

	}

	return sum, nil
}
