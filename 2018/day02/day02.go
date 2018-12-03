package day02

import (
	"bufio"
	"io"
)

func Part1(input io.Reader) (int, error) {
	var count struct {
		two, three int
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		seen := map[byte]int{}
		for _, char := range scanner.Bytes() {
			seen[char]++
		}

		var counted struct {
			two, three bool
		}
		for _, val := range seen {
			if val == 2 && !counted.two {
				count.two++
				counted.two = true
			}
			if val == 3 && !counted.three {
				count.three++
				counted.three = true
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err

	}

	return count.two * count.three, nil
}
