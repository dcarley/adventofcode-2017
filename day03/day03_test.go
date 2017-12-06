package day03_test

import (
	. "github.com/dcarley/adventofcode-2017/day03"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day03", func() {
	Describe("Part1", func() {
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
})
