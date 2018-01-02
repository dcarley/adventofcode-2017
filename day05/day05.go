package day05

import (
	"bufio"
	"io"
	"strconv"
)

func Part1(input io.Reader) (int, error) {
	offsetFunc := func(offset int) int {
		return 1
	}

	return bothParts(input, offsetFunc)
}

func Part2(input io.Reader) (int, error) {
	offsetFunc := func(offset int) int {
		if offset >= 3 {
			return -1
		}

		return 1
	}

	return bothParts(input, offsetFunc)
}

func bothParts(input io.Reader, offsetFunc func(int) int) (int, error) {
	var steps, position int

	instructions, err := ParseInstructions(input)
	if err != nil {
		return 0, err
	}

	for position >= 0 && position < len(instructions) {
		offset := instructions[position]
		instructions[position] += offsetFunc(offset)
		position = position + offset
		steps++
	}

	return steps, nil
}

func ParseInstructions(input io.Reader) ([]int, error) {
	instructions := []int{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		offset, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return []int{}, err
		}

		instructions = append(instructions, offset)
	}

	if err := scanner.Err(); err != nil {
		return []int{}, err
	}

	return instructions, nil
}
