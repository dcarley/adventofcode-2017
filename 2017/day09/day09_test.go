package day09_test

import (
	"bytes"
	"os"

	. "github.com/dcarley/adventofcode/2017/day09"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day09", func() {
	Describe("Part1", func() {
		DescribeTable("examples",
			func(input []byte, expected int) {
				score, err := Part1(bytes.NewBuffer(input))
				Expect(err).ToNot(HaveOccurred())
				Expect(score).To(Equal(expected))
			},
			Entry("{}", []byte("{}"), 1),
			Entry("{{{}}}", []byte("{{{}}}"), 6),
			Entry("{{},{}}", []byte("{{},{}}"), 5),
			Entry("{{{},{},{{}}}}", []byte("{{{},{},{{}}}}"), 16),
			Entry("{<a>,<a>,<a>,<a>}", []byte("{<a>,<a>,<a>,<a>}"), 1),
			Entry("{{<ab>},{<ab>},{<ab>},{<ab>}}", []byte("{{<ab>},{<ab>},{<ab>},{<ab>}}"), 9),
			Entry("{{<!!>},{<!!>},{<!!>},{<!!>}}", []byte("{{<!!>},{<!!>},{<!!>},{<!!>}}"), 9),
			Entry("{{<a!>},{<a!>},{<a!>},{<ab>}}", []byte("{{<a!>},{<a!>},{<a!>},{<ab>}}"), 3),
		)

		It("should solve puzzle input", func() {
			file, err := os.Open("day09.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			score, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(score).To(Equal(10050))
		})
	})

	Describe("Part2", func() {
		DescribeTable("examples",
			func(input []byte, expected int) {
				garbage, err := Part2(bytes.NewBuffer(input))
				Expect(err).ToNot(HaveOccurred())
				Expect(garbage).To(Equal(expected))
			},
			Entry("<>", []byte("<>"), 0),
			Entry("<random characters>", []byte("<random characters>"), 17),
			Entry("<<<<>", []byte("<<<<>"), 3),
			Entry("<{!>}>", []byte("<{!>}>"), 2),
			Entry("<!!>", []byte("<!!>"), 0),
			Entry("<!!!>>", []byte("<!!!>>"), 0),
			Entry(`<{o"i!a,<{i<a>`, []byte(`<{o"i!a,<{i<a>`), 10),
		)

		It("should solve puzzle input", func() {
			file, err := os.Open("day09.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			garbage, err := Part2(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(garbage).To(Equal(4482))
		})
	})
})
