package day17

func Spinlock(step, iterations int) ([]int, int) {
	var (
		position int
		buffer   = []int{position}
	)

	for i := 1; i <= iterations; i++ {
		position = (position+step)%len(buffer) + 1
		buffer = append(
			buffer[:position],
			append([]int{i}, buffer[position:]...)...,
		)
	}

	return buffer, position
}

func Part1(step int) int {
	const iterations = 2017
	buffer, position := Spinlock(step, iterations)

	return buffer[position+1]
}

func Part2(step int) int {
	const iterations = 50000000
	var position, secondItem int

	for i := 1; i <= iterations; i++ {
		position = (position+step)%i + 1
		if position == 1 { // 0 is always the first item
			secondItem = i
		}
	}

	return secondItem
}
