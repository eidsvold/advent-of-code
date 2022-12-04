package puzzle

import (
	"advent-of-code/2022/2/pkg/utils"

	"fmt"
	"strings"
)

type Hand struct {
	points int
	beats  string
}

var handMap = map[string]Hand{
	"A": {points: 1, beats: "Z"},
	"B": {points: 2, beats: "X"},
	"C": {points: 3, beats: "Y"},
	"X": {points: 1, beats: "C"},
	"Y": {points: 2, beats: "A"},
	"Z": {points: 3, beats: "B"},
}

func playHand(opponent string, player string) (int, error) {
	opponentHand, found := handMap[opponent]

	if !found {
		return 0, fmt.Errorf("unknown opponent hand: %s", opponent)
	}

	playerHand, found := handMap[player]

	if !found {
		return 0, fmt.Errorf("unknown player hand: %s", player)
	}

	if playerHand.beats == opponent {
		return (playerHand.points + 6), nil
	}

	if opponentHand.beats == player {
		return playerHand.points, nil
	}

	return (playerHand.points + 3), nil
}

func SolveFirstPuzzle() (int, error) {
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

		pts, err := playHand(slice[0], slice[1])

		if err != nil {
			return 0, err
		}

		sum += pts
	}

	return sum, nil
}
