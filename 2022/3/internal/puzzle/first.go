package puzzle

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/eidsvold/advent-of-code/2022/3/pkg/scanner"
)

func missortedItemTypes(first string, second string) string {
	var itemTypes string

	for _, ch := range first {
		if strings.ContainsRune(second, ch) && !strings.ContainsRune(itemTypes, ch) {
			itemTypes += string(ch)
		}
	}

	return itemTypes
}

func itemTypePriority(item rune) int {
	if unicode.IsUpper(item) {
		return int((item + 26) - 64)
	} else {
		return int(unicode.ToUpper(item) - 64)
	}
}

func SolveFirstPuzzle() (int, error) {
	scanner, err := scanner.NewLineScanner("assets/input.txt")

	if err != nil {
		return 0, err
	}

	sum := 0

	for scanner.Scan() {
		items := scanner.Text()
		mid := len(items) / 2

		cmptOne, cmptTwo := items[0:mid], items[mid:]

		if len(cmptOne) != len(cmptTwo) {
			return 0, fmt.Errorf("compartments do not contain same amount of items")
		}

		missorted := missortedItemTypes(cmptOne, cmptTwo)

		for _, ch := range missorted {
			sum += itemTypePriority(ch)
		}
	}

	return sum, nil
}
