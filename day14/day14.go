package day14

import (
	"fmt"

	"github.com/dcarley/adventofcode-2017/day10"
)

const Size = 128

func Part1(input []byte) int {
	hashes := make([][]byte, Size)
	for i := 0; i < len(hashes); i++ {
		suffix := fmt.Sprintf("-%d", i)
		hashes[i] = day10.Part2(append(input, []byte(suffix)...))
	}

	var used int
	for _, hash := range hashes {
		for _, char := range hash {
			for mask := byte(8); mask >= 1; mask /= 2 {
				if fromHexChar(char)&mask > 0 {
					used++
				}
			}
		}
	}

	return used
}

// https://golang.org/src/encoding/hex/hex.go?s=1552:1593#L73
func fromHexChar(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}

	return 0
}
