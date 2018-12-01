package day15_test

import (
	. "github.com/dcarley/adventofcode-2017/day15"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day15", func() {
	const (
		exampleA = 65
		exampleB = 8921
		inputA   = 618
		inputB   = 814
	)

	Describe("Part1", func() {
		It("should solve example", func() {
			Expect(Part1(exampleA, exampleB)).To(Equal(588))
		})

		It("should solve puzzle input", func() {
			Expect(Part1(inputA, inputB)).To(Equal(577))
		})
	})

	Describe("Part2", func() {
		It("should solve example", func() {
			Expect(Part2(exampleA, exampleB)).To(Equal(309))
		})

		It("should solve puzzle input", func() {
			Expect(Part2(inputA, inputB)).To(Equal(316))
		})
	})
})
