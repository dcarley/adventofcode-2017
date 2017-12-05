package day05_test

import (
	"io/ioutil"

	. "github.com/dcarley/adventofcode-2017/day05"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day05", func() {
	Describe("Part1", func() {
		It("should solve the example", func() {
			input := []byte(`0
3
0
1
-3`)
			steps, err := Part1(input)
			Expect(err).ToNot(HaveOccurred())
			Expect(steps).To(Equal(5))
		})

		It("should solve puzzle input", func() {
			input, err := ioutil.ReadFile("day05.input")
			Expect(err).ToNot(HaveOccurred())

			steps, err := Part1(input)
			Expect(err).ToNot(HaveOccurred())
			Expect(steps).To(Equal(354121))
		})
	})

	Describe("Part2", func() {
		It("should solve the example", func() {
			input := []byte(`0
3
0
1
-3`)
			steps, err := Part2(input)
			Expect(err).ToNot(HaveOccurred())
			Expect(steps).To(Equal(10))
		})

		It("should solve puzzle input", func() {
			input, err := ioutil.ReadFile("day05.input")
			Expect(err).ToNot(HaveOccurred())

			steps, err := Part2(input)
			Expect(err).ToNot(HaveOccurred())
			Expect(steps).To(Equal(27283023))
		})
	})
})
