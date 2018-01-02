package daySKELETON_test

import (
	"bytes"
	"io"
	"os"

	. "github.com/dcarley/adventofcode-2017/daySKELETON"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("DaySKELETON", func() {
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
			file, err := os.Open("daySKELETON.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			res, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(-1))
		})
	})
})
