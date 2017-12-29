package day10

func Part1(list, lengths []int) int {
	var position int

	for skip, length := range lengths {
		ReverseSublistWrapped(list, position, length)
		position += length + skip
		position %= len(list) // wrap if it exceeds length of list
	}

	return list[0] * list[1]
}

func ReverseSublistWrapped(list []int, position, length int) {
	for i := 0; i < length/2; i++ {
		// mod wraps the index if it exceeds the length of the list
		left := (position + i) % len(list)
		right := (position + length - 1 - i) % len(list)

		list[left], list[right] = list[right], list[left]
	}
}
