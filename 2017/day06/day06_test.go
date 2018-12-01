package day06_test

import (
	"os"

	. "github.com/dcarley/adventofcode-2017/day06"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day06", func() {
	Describe("Part1", func() {
		DescribeTable("Rebalance examples",
			func(input, expected []int) {
				Expect(Rebalance(input)).To(Equal(expected))
			},
			Entry("0 2 7 0", []int{0, 2, 7, 0}, []int{2, 4, 1, 2}),
			Entry("2 4 1 2", []int{2, 4, 1, 2}, []int{3, 1, 2, 3}),
			Entry("3 1 2 3", []int{3, 1, 2, 3}, []int{0, 2, 3, 4}),
			Entry("0 2 3 4", []int{0, 2, 3, 4}, []int{1, 3, 4, 1}),
			Entry("1 3 4 1", []int{1, 3, 4, 1}, []int{2, 4, 1, 2}),
		)

		It("should solve example", func() {
			Expect(Part1([]int{0, 2, 7, 0})).To(Equal(5))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day06.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			bank, err := Parse(file)
			Expect(err).ToNot(HaveOccurred())

			Expect(Part1(bank)).To(Equal(6681))
		})
	})

	Describe("Part2", func() {
		It("should solve example", func() {
			Expect(Part2([]int{0, 2, 7, 0})).To(Equal(4))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day06.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			bank, err := Parse(file)
			Expect(err).ToNot(HaveOccurred())

			Expect(Part2(bank)).To(Equal(2392))
		})
	})
})
