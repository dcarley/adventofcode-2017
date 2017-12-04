package day04

import (
	"bufio"
	"bytes"
)

func Part1(input []byte) bool {
	words := map[string]struct{}{}

	wordScanner := bufio.NewScanner(bytes.NewBuffer(input))
	wordScanner.Split(bufio.ScanWords)

	for wordScanner.Scan() {
		if _, exists := words[wordScanner.Text()]; exists {
			return false
		}

		words[wordScanner.Text()] = struct{}{}
	}

	return true
}
