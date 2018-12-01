package day15

const (
	Product = 2147483647
	FactorA = 16807
	FactorB = 48271
	Mask    = 0xFFFF
	Buffer  = 4096
)

func Generator(seed, factor, iterations, modFilter int) <-chan int {
	results := make(chan int, Buffer)

	go func() {
		defer close(results)
		var (
			i   int
			val = seed
		)

		for i < iterations {
			val = (val * factor) % Product
			if val%modFilter == 0 {
				results <- val
				i++
			}
		}
	}()

	return results
}

func Judge(chanA, chanB <-chan int) int {
	var count int
	for {
		a, ok := <-chanA
		if !ok {
			break
		}
		b, ok := <-chanB
		if !ok {
			break
		}
		if a&Mask == b&Mask {
			count++
		}
	}

	return count
}

func Part1(a, b int) int {
	const iterations = 40000000

	var count int
	for i := 0; i < iterations; i++ {
		a = (a * FactorA) % Product
		b = (b * FactorB) % Product
		if a&Mask == b&Mask {
			count++
		}
	}

	return count
}

func Part2(a, b int) int {
	const iterations = 5000000

	chanA := Generator(a, FactorA, iterations, 4)
	chanB := Generator(b, FactorB, iterations, 8)

	return Judge(chanA, chanB)
}
