package day16_test

import (
	"bytes"
	"os"

	. "github.com/dcarley/adventofcode/2017/day16"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day16", func() {
	Describe("Part1", func() {
		DescribeTable("examples",
			func(input, expected []byte) {
				result, err := Part1(bytes.NewBuffer(input))
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal(expected))
			},
			Entry("s3", []byte("s3"), []byte("nopabcdefghijklm")),
			Entry("x3/4", []byte("x3/4"), []byte("abcedfghijklmnop")),
			Entry("pe/b", []byte("pe/b"), []byte("aecdbfghijklmnop")),
			Entry("s3,x3/4,pe/b", []byte("s3,x3/4,pe/b"), []byte("nopeacdbfghijklm")),
		)

		It("should solve puzzle input", func() {
			file, err := os.Open("day16.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			result, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal([]byte("pkgnhomelfdibjac")))
		})
	})

	Describe("Part2", func() {
		It("should solve puzzle input", func() {
			file, err := os.Open("day16.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			result, err := Part2(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal([]byte("pogbjfihclkemadn")))
		})
	})
})
