package day14_test

import (
	. "github.com/dcarley/adventofcode/2017/day14"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day14", func() {
	var (
		input, exampleInput []byte
	)

	BeforeEach(func() {
		input = []byte(`stpzcrnm`)
		exampleInput = []byte(`flqrgnkx`)
	})

	Describe("Part1", func() {
		It("should solve example", func() {
			Expect(Part1(exampleInput)).To(Equal(8108))
		})

		It("should solve puzzle input", func() {
			Expect(Part1(input)).To(Equal(8250))
		})
	})

	Describe("Part2", func() {
		It("should solve example", func() {
			Expect(Part2(exampleInput)).To(Equal(1242))
		})

		It("should solve puzzle input", func() {
			Expect(Part2(input)).To(Equal(1113))
		})
	})
})
