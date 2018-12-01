package day05_test

import (
	"bytes"
	"io"
	"os"

	. "github.com/dcarley/adventofcode-2017/day05"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day05", func() {
	var exampleInput io.Reader

	BeforeEach(func() {
		exampleInput = bytes.NewBuffer([]byte(`0
3
0
1
-3`))
	})

	Describe("Part1", func() {
		It("should solve the example", func() {
			steps, err := Part1(exampleInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(steps).To(Equal(5))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day05.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			steps, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(steps).To(Equal(354121))
		})
	})

	Describe("Part2", func() {
		It("should solve the example", func() {
			steps, err := Part2(exampleInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(steps).To(Equal(10))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day05.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			steps, err := Part2(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(steps).To(Equal(27283023))
		})
	})
})
