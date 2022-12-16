package scanner

import (
	"bufio"
	"os"
)

func NewLineScanner(fileName string) (*bufio.Scanner, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return scanner, nil
}
