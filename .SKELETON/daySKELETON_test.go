package day{{day}}_test

import (
	"bytes"
	"io"
	"os"

	. "github.com/dcarley/adventofcode/{{year}}/day{{day}}"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day{{day}}", func() {
	var exampleInput io.Reader

	BeforeEach(func() {
		exampleInput = bytes.NewBuffer([]byte{})
	})

	Describe("Part1", func() {
		DescribeTable("examples",
			func(expected int) {
				res, err := Part1(exampleInput)
				Expect(err).ToNot(HaveOccurred())
				Expect(res).To(Equal(expected))
			},
			Entry("XXX", 0),
		)

		It("should solve example", func() {
			res, err := Part1(exampleInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(-1))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day{{day}}.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			res, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(-1))
		})
	})
})
