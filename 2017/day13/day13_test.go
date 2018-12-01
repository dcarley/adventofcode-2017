package day13_test

import (
	"bytes"
	"io"
	"os"

	. "github.com/dcarley/adventofcode/2017/day13"

	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day13", func() {
	var exampleInput io.Reader

	BeforeEach(func() {
		exampleInput = bytes.NewBuffer([]byte(`0: 3
1: 2
4: 4
6: 4`))
	})

	Describe("Part1", func() {
		It("should solve example", func() {
			res, err := Part1(exampleInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(24))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day13.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			res, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(1528))
		})
	})

	Describe("Part2", func() {
		It("should solve example", func() {
			res, err := Part2(exampleInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(10))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day13.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			res, err := Part2(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(3896406))
		})
	})
})
