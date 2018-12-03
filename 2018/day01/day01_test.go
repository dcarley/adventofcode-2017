package day01_test

import (
	"bytes"
	"os"

	. "github.com/dcarley/adventofcode/2018/day01"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day01", func() {
	Describe("Part1", func() {
		DescribeTable("examples",
			func(input []byte, expected int) {
				res, err := Part1(bytes.NewBuffer(input))
				Expect(err).ToNot(HaveOccurred())
				Expect(res).To(Equal(expected))
			},
			Entry("+1, -2, +3, +1", []byte("+1\n-2\n+3\n+1\n"), 3),
			Entry("+1, +1, +1", []byte("+1\n+1\n+1\n"), 3),
			Entry("+1, +1, -2", []byte("+1\n+1\n-2\n"), 0),
			Entry("-1, -2, -3", []byte("-1\n-2\n-3\n"), -6),
		)

		It("should solve puzzle input", func() {
			file, err := os.Open("day01.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			res, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(580))
		})
	})

	Describe("Part2", func() {
		DescribeTable("examples",
			func(input []byte, expected int) {
				res, err := Part2(bytes.NewBuffer(input))
				Expect(err).ToNot(HaveOccurred())
				Expect(res).To(Equal(expected))
			},
			Entry("+1, -1", []byte("+1\n-1\n"), 0),
			Entry("+3, +3, +4, -2, -4", []byte("+3\n+3\n+4\n-2\n-4"), 10),
			Entry("-6, +3, +8, +5, -6", []byte("-6\n+3\n+8\n+5\n-6"), 5),
			Entry("+7, +7, -2, -7, -4", []byte("+7\n+7\n-2\n-7\n-4"), 14),
		)

		It("should solve puzzle input", func() {
			file, err := os.Open("day01.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			res, err := Part2(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(81972))
		})
	})
})
