package day17_test

import (
	. "github.com/dcarley/adventofcode-2017/day17"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day17", func() {
	const (
		exampleInput = 3
		input        = 371
	)

	Describe("Part1", func() {
		DescribeTable("examples",
			func(iterations int, expectedBuffer []int, expectedPosition int) {
				buffer, position := Spinlock(exampleInput, iterations)
				Expect(buffer).To(Equal(expectedBuffer))
				Expect(position).To(Equal(expectedPosition))
			},
			Entry("iteration 1", 1, []int{0, 1}, 1),
			Entry("iteration 2", 2, []int{0, 2, 1}, 1),
			Entry("iteration 3", 3, []int{0, 2, 3, 1}, 2),
			Entry("iteration 4", 4, []int{0, 2, 4, 3, 1}, 2),
			Entry("iteration 5", 5, []int{0, 5, 2, 4, 3, 1}, 1),
			Entry("iteration 6", 6, []int{0, 5, 2, 4, 3, 6, 1}, 5),
			Entry("iteration 7", 7, []int{0, 5, 7, 2, 4, 3, 6, 1}, 2),
			Entry("iteration 8", 8, []int{0, 5, 7, 2, 4, 3, 8, 6, 1}, 6),
			Entry("iteration 9", 9, []int{0, 9, 5, 7, 2, 4, 3, 8, 6, 1}, 1),
		)

		It("should solve example", func() {
			Expect(Part1(exampleInput)).To(Equal(638))
		})

		It("should solve puzzle input", func() {
			Expect(Part1(input)).To(Equal(1311))
		})
	})

	Describe("Part2", func() {
		It("should solve puzzle input", func() {
			Expect(Part2(input)).To(Equal(39170601))
		})
	})
})
