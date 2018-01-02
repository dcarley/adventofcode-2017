package day05

import (
	"bufio"
	"io"
	"strconv"
)

func Part1(input io.Reader) (int, error) {
	return bothParts(input, false)
}

func Part2(input io.Reader) (int, error) {
	return bothParts(input, true)
}

func bothParts(input io.Reader, part2 bool) (int, error) {
	var steps, position int

	instructions, err := ParseInstructions(input)
	if err != nil {
		return 0, err
	}

	for position >= 0 && position < len(instructions) {
		offset := instructions[position]
		if part2 && offset >= 3 {
			instructions[position] -= 1
		} else {
			instructions[position] += 1
		}

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
