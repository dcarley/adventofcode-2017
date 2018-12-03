package day02_test

import (
	"bytes"
	"io"
	"os"

	. "github.com/dcarley/adventofcode/2018/day02"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day02", func() {
	var exampleInput io.Reader

	BeforeEach(func() {
		exampleInput = bytes.NewBuffer([]byte(`abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab
`))
	})

	Describe("Part1", func() {
		It("should solve example", func() {
			res, err := Part1(exampleInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(12))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day02.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			res, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(6448))
		})
	})
})
