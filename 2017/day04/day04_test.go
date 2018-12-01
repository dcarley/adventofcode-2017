package day04_test

import (
	"bufio"
	"os"

	. "github.com/dcarley/adventofcode/2017/day04"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day04", func() {
	Describe("Part1", func() {
		DescribeTable("examples",
			func(input []byte, expected bool) {
				Expect(Part1(input)).To(Equal(expected))
			},
			Entry("aa bb cc dd ee", []byte("aa bb cc dd ee"), true),
			Entry("aa bb cc dd aa", []byte("aa bb cc dd aa"), false),
			Entry("aa bb cc dd aaa", []byte("aa bb cc dd aaa"), true),
		)

		It("should solve puzzle input", func() {
			var valid int

			file, err := os.Open("day04.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			lineScanner := bufio.NewScanner(file)
			for lineScanner.Scan() {
				if Part1(lineScanner.Bytes()) {
					valid++
				}
			}

			Expect(lineScanner.Err()).ToNot(HaveOccurred())
			Expect(valid).To(Equal(325))
		})
	})
	Describe("Part2", func() {
		DescribeTable("examples",
			func(input []byte, expected bool) {
				Expect(Part2(input)).To(Equal(expected))
			},
			Entry("abcde fghij", []byte("abcde fghij"), true),
			Entry("abcde xyz ecdab", []byte("abcde xyz ecdab"), false),
			Entry("a ab abc abd abf abj", []byte("a ab abc abd abf abj"), true),
			Entry("iiii oiii ooii oooi oooo", []byte("iiii oiii ooii oooi oooo"), true),
			Entry("oiii ioii iioi iiio", []byte("oiii ioii iioi iiio"), false),
		)

		It("should solve puzzle input", func() {
			var valid int

			file, err := os.Open("day04.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			lineScanner := bufio.NewScanner(file)
			for lineScanner.Scan() {
				if Part2(lineScanner.Bytes()) {
					valid++
				}
			}

			Expect(lineScanner.Err()).ToNot(HaveOccurred())
			Expect(valid).To(Equal(119))
		})
	})
})
