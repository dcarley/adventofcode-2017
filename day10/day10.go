package day10

import (
	"encoding/hex"
)

func Part1(list, lengths []int) int {
	var position int

	for skip, length := range lengths {
		ReverseSublistWrapped(list, position, length)
		position += length + skip
		position %= len(list) // wrap if it exceeds length of list
	}

	return list[0] * list[1]
}

func Part2(lengths []byte) []byte {
	list := BuildList(256)
	suffix := []byte{17, 31, 73, 47, 23}
	lengths = append(lengths, suffix...)

	var position, skip int
	for i := 0; i < 64; i++ {
		for _, length := range lengths {
			ReverseSublistWrapped(list, position, int(length))

			position += int(length) + skip
			position %= len(list)
			skip++
		}
	}

	denseHash := make([]byte, 16)
	for denseBlock := 0; denseBlock < len(denseHash); denseBlock++ {
		for i := 0; i < len(denseHash); i++ {
			denseHash[denseBlock] ^= byte(list[denseBlock*16+i])
		}
	}

	hash := make([]byte, hex.EncodedLen(len(denseHash)))
	hex.Encode(hash, denseHash)

	return hash
}

func BuildList(size int) []int {
	list := make([]int, size)
	for i := 0; i < len(list); i++ {
		list[i] = i
	}

	return list
}

func ReverseSublistWrapped(list []int, position, length int) {
	for i := 0; i < length/2; i++ {
		// mod wraps the index if it exceeds the length of the list
		left := (position + i) % len(list)
		right := (position + length - 1 - i) % len(list)

		list[left], list[right] = list[right], list[left]
	}
}
