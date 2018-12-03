package day01

import (
	"fmt"
	"io"
)

func Part1(input io.Reader) (int, error) {
	changes, err := parseChanges(input)
	if err != nil {
		return 0, err
	}

	var freq int
	for i := 0; i < len(changes); i++ {
		freq += changes[i]
	}

	return freq, nil
}

func Part2(input io.Reader) (int, error) {
	changes, err := parseChanges(input)
	if err != nil {
		return 0, err
	}

	var (
		freq int
		dupe bool
	)
	seen := map[int]struct{}{
		freq: struct{}{},
	}
	for i := 0; !dupe; i++ {
		if i == len(changes) {
			i = 0
		}

		freq += changes[i]
		if _, ok := seen[freq]; ok {
			dupe = true
		} else {
			seen[freq] = struct{}{}
		}
	}

	return freq, nil
}

func parseChanges(input io.Reader) ([]int, error) {
	var changes []int
	for {
		var change int
		_, err := fmt.Fscanf(input, "%d\n", &change)
		if err == io.EOF {
			break
		}
		if err != nil {
			return []int{}, err
		}

		changes = append(changes, change)
	}

	return changes, nil
}
