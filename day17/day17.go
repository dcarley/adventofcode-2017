package day17

const Iterations = 2017

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
	buffer, position := Spinlock(step, Iterations)

	return buffer[(position+1)%len(buffer)]
}
