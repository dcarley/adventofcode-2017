package day10_test

import (
	. "github.com/dcarley/adventofcode-2017/day10"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func buildList(size int) []int {
	list := make([]int, size)
	for i := 0; i < len(list); i++ {
		list[i] = i
	}

	return list
}

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
			list := buildList(5)
			inputs := []int{3, 4, 1, 5}
			Expect(Part1(list, inputs)).To(Equal(12))
		})

		It("should solve puzzle input", func() {
			list := buildList(256)
			inputs := []int{106, 118, 236, 1, 130, 0, 235, 254, 59, 205, 2, 87, 129, 25, 255, 118}
			Expect(Part1(list, inputs)).To(Equal(6909))
		})
	})

	Describe("Part2", func() {
		DescribeTable("examples",
			func(inputs []byte, expected string) {
				Expect(Part2(buildList(256), inputs)).To(Equal(expected))
			},
			Entry("empty string", []byte(""), "a2582a3a0e66e6e86e3812dcb672a272"),
			Entry("AoC 2017", []byte("AoC 2017"), "33efeb34ea91902bb2f59c9920caa6cd"),
			Entry("1,2,3", []byte("1,2,3"), "3efbe78a8d82f29979031a4aa0b16a9d"),
			Entry("1,2,4", []byte("1,2,4"), "63960835bcdc130f0b66d7ff4f6a5a8e"),
		)

		It("should solve puzzle input", func() {
			inputs := []byte("106,118,236,1,130,0,235,254,59,205,2,87,129,25,255,118")
			Expect(Part2(buildList(256), inputs)).To(Equal("9d5f4561367d379cfbf04f8c471c0095"))
		})
	})
})
