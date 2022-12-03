package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var movePtsMap = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var ptsWinMap = map[int]int{
	1: 3,
	2: 1,
	3: 2,
}

func playRound(opponent string, you string) int {
	// TODO: return error for unknown cmd

	opponentPts := movePtsMap[opponent]
	yourPts := movePtsMap[you]

	if ptsWinMap[yourPts] == opponentPts {
		return 6 + yourPts
	}

	if ptsWinMap[opponentPts] == yourPts {
		return 0 + yourPts
	}

	return 3 + yourPts
}

func firstPuzzle() (int, error) {
	file, err := os.Open("input.txt")

	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan() {
		slice := strings.Split(scanner.Text(), " ")

		if len(slice) != 2 {
			return 0, errors.New("Line is not of length 2")
		}

		pts := playRound(slice[0], slice[1])
		sum += pts
	}

	return sum, nil
}

func main() {
	first, err := firstPuzzle()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(first)
}
