package day04

import (
	"bufio"
	"bytes"
	"sort"
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

type ByteSlice []byte

func (p ByteSlice) Len() int           { return len(p) }
func (p ByteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p ByteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Part2(input []byte) bool {
	wordScanner := bufio.NewScanner(bytes.NewBuffer(input))
	wordScanner.Split(bufio.ScanWords)

	words := map[string]struct{}{}
	for wordScanner.Scan() {
		word := wordScanner.Bytes()
		sort.Sort(ByteSlice(wordScanner.Bytes()))
		if _, exists := words[string(word)]; exists {
			return false
		}

		words[string(word)] = struct{}{}
	}

	return true
}
