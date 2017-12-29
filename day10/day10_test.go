package day10_test

import (
	. "github.com/dcarley/adventofcode-2017/day10"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day10", func() {
	Describe("Part1", func() {
		DescribeTable("twist examples",
			func(in, out []int, position, length int) {
				ReverseSublistWrapped(in, position, length)
				Expect(in).To(Equal(out))
			},
			Entry("twist 1",
				[]int{0, 1, 2, 3, 4},
				[]int{2, 1, 0, 3, 4},
				0, 3,
			),
			Entry("twist 2",
				[]int{2, 1, 0, 3, 4},
				[]int{4, 3, 0, 1, 2},
				3, 4,
			),
			Entry("twist 3",
				[]int{4, 3, 0, 1, 2},
				[]int{4, 3, 0, 1, 2},
				3, 1,
			),
			Entry("twist 4",
				[]int{4, 3, 0, 1, 2},
				[]int{3, 4, 2, 1, 0},
				1, 5,
			),
		)

		It("should solve example", func() {
			list := make([]int, 5)
			for i := 0; i < len(list); i++ {
				list[i] = i
			}
			inputs := []int{3, 4, 1, 5}
			Expect(Part1(list, inputs)).To(Equal(12))
		})

		It("should solve puzzle input", func() {
			list := make([]int, 256)
			for i := 0; i < len(list); i++ {
				list[i] = i
			}
			inputs := []int{106, 118, 236, 1, 130, 0, 235, 254, 59, 205, 2, 87, 129, 25, 255, 118}
			Expect(Part1(list, inputs)).To(Equal(6909))
		})
	})
})
