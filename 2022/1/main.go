package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func firstPuzzle() (int, error) {
	file, err := os.Open("input.txt")

	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0
	max := 0

	for scanner.Scan() {
		txt := scanner.Text()

		if txt == "" {
			if sum > max {
				max = sum
			}

			sum = 0
		} else {
			i, err := strconv.Atoi(txt)

			if err != nil {
				return 0, err
			}

			sum += i
		}
	}

	return max, nil
}

func secondPuzzle() (int, error) {
	file, err := os.Open("input.txt")

	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var elves []int
	sum := 0

	for scanner.Scan() {
		txt := scanner.Text()

		if txt == "" {
			elves = append(elves, sum)
			sum = 0
		} else {
			i, err := strconv.Atoi(txt)

			if err != nil {
				return 0, err
			}

			sum += i
		}
	}

	sort.Ints(elves)
	total := 0

	for _, i := range elves[len(elves)-3:] {
		total += i
	}

	return total, nil
}

func main() {
	first, err := firstPuzzle()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(first)

	second, err := secondPuzzle()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(second)
}
