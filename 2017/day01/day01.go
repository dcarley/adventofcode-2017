package day01

import (
	"errors"
	"strconv"
)

func Part1(input string) (int, error) {
	return bothParts(input, 1)
}

func Part2(input string) (int, error) {
	if len(input)%2 != 0 {
		return 0, errors.New("input length is not even")
	}

	return bothParts(input, len(input)/2)
}

func bothParts(input string, step int) (int, error) {
	var sum, compareIndex int

	for i := 0; i < len(input); i++ {
		compareIndex = i + step
		if compareIndex >= len(input) {
			compareIndex %= len(input)
		}

		if input[i] == input[compareIndex] {
			digit, err := strconv.Atoi(string(input[i]))
			if err != nil {
				return 0, err
			}

			sum += digit
		}
	}

	return sum, nil
}
