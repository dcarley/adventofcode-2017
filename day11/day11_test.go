package day11_test

import (
	"io/ioutil"

	. "github.com/dcarley/adventofcode-2017/day11"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day11", func() {
	Describe("Part1", func() {
		DescribeTable("examples",
			func(path string, expected int) {
				distance, err := Part1(path)
				Expect(err).ToNot(HaveOccurred())
				Expect(distance).To(Equal(expected))
			},
			Entry("ne,ne,ne", "ne,ne,ne", 3),
			Entry("ne,ne,sw,sw", "ne,ne,sw,sw", 0),
			Entry("ne,ne,s,s", "ne,ne,s,s", 2),
			Entry("se,sw,se,sw,sw", "se,sw,se,sw,sw", 3),
		)

		It("should solve puzzle input", func() {
			path, err := ioutil.ReadFile("day11.input")
			Expect(err).ToNot(HaveOccurred())

			distance, err := Part1(string(path))
			Expect(err).ToNot(HaveOccurred())
			Expect(distance).To(Equal(643))
		})
	})
})
