package day01_test

import (
	"io/ioutil"
	"strings"

	. "github.com/dcarley/adventofcode/2017/day01"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day01", func() {
	Describe("Part1", func() {
		DescribeTable("examples",
			func(input string, expected int) {
				captcha, err := Part1(input)
				Expect(err).ToNot(HaveOccurred())
				Expect(captcha).To(Equal(expected))
			},
			Entry("1122", "1122", 3),
			Entry("1111", "1111", 4),
			Entry("1234", "1234", 0),
			Entry("91212129", "91212129", 9),
		)

		It("should solve puzzle input", func() {
			input, err := ioutil.ReadFile("day01.input")
			Expect(err).ToNot(HaveOccurred())

			captcha, err := Part1(strings.TrimSpace(string(input)))
			Expect(err).ToNot(HaveOccurred())
			Expect(captcha).To(Equal(1203))
		})
	})

	Describe("Part2", func() {
		DescribeTable("examples",
			func(input string, expected int) {
				captcha, err := Part2(input)
				Expect(err).ToNot(HaveOccurred())
				Expect(captcha).To(Equal(expected))
			},
			Entry("1212", "1212", 6),
			Entry("1221", "1221", 0),
			Entry("123425", "123425", 4),
			Entry("123123", "123123", 12),
			Entry("12131415", "12131415", 4),
		)

		It("should solve puzzle input", func() {
			input, err := ioutil.ReadFile("day01.input")
			Expect(err).ToNot(HaveOccurred())

			captcha, err := Part2(strings.TrimSpace(string(input)))
			Expect(err).ToNot(HaveOccurred())
			Expect(captcha).To(Equal(1146))
		})

		It("should error on odd length input", func() {
			_, err := Part2("123")
			Expect(err).To(MatchError("input length is not even"))
		})
	})
})
