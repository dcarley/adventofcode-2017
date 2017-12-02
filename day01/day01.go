package day01

import (
	"strconv"
)

func Part1(input string) (int, error) {
	var sum, next int

	for i := 0; i < len(input); i++ {
		next = i + 1
		if next >= len(input) {
			next = 0
		}

		if input[i] == input[next] {
			digit, err := strconv.Atoi(string(input[i]))
			if err != nil {
				return 0, err
			}

			sum += digit
		}
	}

	return sum, nil
}
