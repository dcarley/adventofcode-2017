package day08_test

import (
	"bytes"
	"io"
	"os"

	. "github.com/dcarley/adventofcode-2017/day08"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day08", func() {
	var example io.Reader

	BeforeEach(func() {
		example = bytes.NewBuffer([]byte(`b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`))
	})

	Describe("Part1", func() {
		It("should solve example", func() {
			highest, err := Part1(example)
			Expect(err).ToNot(HaveOccurred())
			Expect(highest).To(Equal(1))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day08.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			highest, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(highest).To(Equal(4888))
		})
	})

	Describe("Part2", func() {
		It("should solve example", func() {
			highest, err := Part2(example)
			Expect(err).ToNot(HaveOccurred())
			Expect(highest).To(Equal(10))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day08.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			highest, err := Part2(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(highest).To(Equal(7774))
		})
	})
})
