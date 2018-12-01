package day09

import (
	"bufio"
	"bytes"
	"io"
)

const (
	GroupOpen    = '{'
	GroupClose   = '}'
	GarbageOpen  = '<'
	GarbageClose = '>'
	IgnoreNext   = '!'
)

func Part1(input io.Reader) (int, error) {
	score, _, err := bothParts(input)

	return score, err
}

func Part2(input io.Reader) (int, error) {
	_, garbage, err := bothParts(input)

	return garbage, err
}

func bothParts(input io.Reader) (score int, garbage int, err error) {
	var (
		depth     int
		inGarbage bool
	)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		token := scanner.Bytes()

		if bytes.ContainsRune(token, IgnoreNext) {
			scanner.Scan()
			continue
		}

		if inGarbage && !bytes.ContainsRune(token, GarbageClose) {
			garbage++
			continue
		}

		switch {
		case bytes.ContainsRune(token, GarbageOpen):
			inGarbage = true
		case bytes.ContainsRune(token, GarbageClose):
			inGarbage = false
		case bytes.ContainsRune(token, GroupOpen):
			depth++
		case bytes.ContainsRune(token, GroupClose):
			score += depth
			depth--
		}
	}

	return score, garbage, scanner.Err()
}
