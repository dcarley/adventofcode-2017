package day12_test

import (
	"bytes"
	"io"
	"os"

	. "github.com/dcarley/adventofcode-2017/day12"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day12", func() {
	var exampleInput io.Reader

	BeforeEach(func() {
		exampleInput = bytes.NewBuffer([]byte(`0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5
`))
	})

	Describe("Part1", func() {
		const group = 0

		It("should solve example", func() {
			members, err := Part1(exampleInput, group)
			Expect(err).ToNot(HaveOccurred())
			Expect(members).To(Equal(6))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day12.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			members, err := Part1(file, group)
			Expect(err).ToNot(HaveOccurred())
			Expect(members).To(Equal(141))
		})
	})

	Describe("Part2", func() {
		It("should solve example", func() {
			members, err := Part2(exampleInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(members).To(Equal(2))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day12.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			members, err := Part2(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(members).To(Equal(171))
		})
	})
})
