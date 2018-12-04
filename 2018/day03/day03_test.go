package day03_test

import (
	"bytes"
	"io"
	"os"

	. "github.com/dcarley/adventofcode/2018/day03"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day03", func() {
	var exampleInput io.Reader

	BeforeEach(func() {
		exampleInput = bytes.NewBuffer([]byte(`#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2
`))
	})

	Describe("Part1", func() {
		It("should solve example", func() {
			res, err := Part1(exampleInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(4))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day03.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			res, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(121163))
		})
	})

	Describe("Part2", func() {
		It("should solve example", func() {
			res, err := Part2(exampleInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(3))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day03.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			res, err := Part2(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(943))
		})
	})
})
