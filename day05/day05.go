package day05

import (
	"bufio"
	"bytes"
	"strconv"
)

func Part1(input []byte) (int, error) {
	var steps, position int

	instructions, err := ParseInstructions(input)
	if err != nil {
		return 0, err
	}

	for position >= 0 && position < len(instructions) {
		offset := instructions[position]
		instructions[position]++
		position = position + offset
		steps++
	}

	return steps, nil
}

func ParseInstructions(input []byte) ([]int, error) {
	instructions := make([]int, 0, bytes.Count(input, []byte{'\n'}))

	scanner := bufio.NewScanner(bytes.NewBuffer(input))
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
