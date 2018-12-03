package day02

import (
	"bufio"
	"fmt"
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

func Part2(input io.Reader) (string, error) {
	var ids []string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	for a := range ids {
		for b := range ids {
			if a == b {
				continue
			}

			if len(ids[a]) != len(ids[b]) {
				return "", fmt.Errorf("different length ids: %s, %s", ids[a], ids[b])
			}

			var diff []int
			for i := 0; i < len(ids[a]); i++ {
				if ids[a][i] != ids[b][i] {
					diff = append(diff, i)
				}
			}

			if len(diff) == 1 {
				i := diff[0]
				return ids[a][:i] + ids[a][i+1:], nil
			}
		}
	}

	return "", fmt.Errorf("no matching ids found")
}
