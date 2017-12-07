package day03_test

import (
	. "github.com/dcarley/adventofcode-2017/day03"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day03", func() {
	DescribeTable("Move",
		func(bearing int, expected Position) {
			position := Move(Position{0, 0}, bearing, 3)
			Expect(position).To(Equal(expected))
		},
		Entry("north", BearNorth, Position{0, -3}),
		Entry("east", BearEast, Position{3, 0}),
		Entry("south", BearSouth, Position{0, 3}),
		Entry("west", BearWest, Position{-3, 0}),
	)

	DescribeTable("Neighbours",
		func(position Position, expected []Position) {
			Expect(Neighbours(position)).To(Equal(expected))
		},
		Entry("0,0", Position{0, 0}, []Position{
			{-1, -1}, {-1, 0}, {-1, 1},
			{0, -1}, {0, 1},
			{1, -1}, {1, 0}, {1, 1},
		}),
		Entry("2,2", Position{2, 2}, []Position{
			{1, 1}, {1, 2}, {1, 3},
			{2, 1}, {2, 3},
			{3, 1}, {3, 2}, {3, 3},
		}),
		Entry("-2,-2", Position{-2, -2}, []Position{
			{-3, -3}, {-3, -2}, {-3, -1},
			{-2, -3}, {-2, -1},
			{-1, -3}, {-1, -2}, {-1, -1},
		}),
	)

	Describe("Part1", func() {
		DescribeTable("examples",
			func(input int, expected int) {
				Expect(Part1(input)).To(Equal(expected))
			},
			Entry("1", 1, 0),
			Entry("12", 12, 3),
			Entry("23", 23, 2),
			Entry("1024", 1024, 31),
		)

		It("should solve puzzle input", func() {
			Expect(Part1(277678)).To(Equal(475))
		})
	})

	Describe("Part2", func() {
		DescribeTable("examples",
			func(input int, expected int) {
				Expect(Part2(input)).To(Equal(expected))
			},
			Entry("value after 2", 2, 4),
			Entry("value after 4", 4, 5),
			Entry("value after 5", 5, 10),
			Entry("value after 10", 10, 11),
			Entry("value after 11", 11, 23),
			Entry("value after 23", 23, 25),
			Entry("value after 25", 25, 26),
		)

		It("should solve puzzle input", func() {
			Expect(Part2(277678)).To(Equal(279138))
		})
	})
})
