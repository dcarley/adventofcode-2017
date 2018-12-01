package day01

import (
	"fmt"
	"io"
)

func Part1(input io.Reader) (int, error) {
	var freq, change int
	for {
		_, err := fmt.Fscanf(input, "%d\n", &change)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

		freq += change
	}

	return freq, nil
}
