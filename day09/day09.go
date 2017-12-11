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
	var (
		depth, score int
		garbage      bool
	)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		token := scanner.Bytes()

		if bytes.ContainsRune(token, IgnoreNext) {
			scanner.Scan()
			continue
		}

		if garbage && !bytes.ContainsRune(token, GarbageClose) {
			continue
		}

		switch {
		case bytes.ContainsRune(token, GarbageOpen):
			garbage = true
		case bytes.ContainsRune(token, GarbageClose):
			garbage = false
		case bytes.ContainsRune(token, GroupOpen):
			depth++
		case bytes.ContainsRune(token, GroupClose):
			score += depth
			depth--
		}
	}

	return score, scanner.Err()
}
