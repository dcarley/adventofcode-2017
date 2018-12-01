package day02_test

import (
	"bytes"
	"io"
	"os"

	. "github.com/dcarley/adventofcode/2017/day02"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day02", func() {
	var exampleInput io.Reader

	BeforeEach(func() {
		exampleInput = bytes.NewBuffer([]byte(`5 9 2 8
9 4 7 3
3 8 6 5`))
	})

	Describe("Part1", func() {
		DescribeTable("example lines",
			func(input []byte, expected int) {
				buf := bytes.NewBuffer(input)
				checksum, err := Part1(buf)
				Expect(err).ToNot(HaveOccurred())
				Expect(expected).To(Equal(checksum))
			},
			Entry("5 1 9 5", []byte("5 1 9 5"), 8),
			Entry("7 5 3", []byte("7 5 3"), 4),
			Entry("2 4 6 8", []byte("2 4 6 8"), 6),
		)

		It("should solve example", func() {
			checksum, err := Part1(exampleInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(checksum).To(Equal(18))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day02.input")
			Expect(err).ToNot(HaveOccurred())

			checksum, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(checksum).To(Equal(43074))
		})
	})

	Describe("Part2", func() {
		DescribeTable("example lines",
			func(input []byte, expected int) {
				buf := bytes.NewBuffer(input)
				checksum, err := Part2(buf)
				Expect(err).ToNot(HaveOccurred())
				Expect(checksum).To(Equal(expected))
			},
			Entry("5 9 2 8", []byte("5 9 2 8"), 4),
			Entry("9 4 7 3", []byte("9 4 7 3"), 3),
			Entry("3 8 6 5", []byte("3 8 6 5"), 2),
		)

		It("should solve example", func() {
			checksum, err := Part2(exampleInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(checksum).To(Equal(9))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day02.input")
			Expect(err).ToNot(HaveOccurred())

			checksum, err := Part2(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(checksum).To(Equal(280))
		})
	})
})
