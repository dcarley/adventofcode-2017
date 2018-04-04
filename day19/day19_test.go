package day19_test

import (
	"bytes"
	"io"
	"os"

	. "github.com/dcarley/adventofcode-2017/day19"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day19", func() {
	var exampleInput io.Reader

	BeforeEach(func() {
		exampleInput = bytes.NewBuffer([]byte(`     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 
`))
	})

	Describe("Part1", func() {
		It("should solve example", func() {
			res, err := Part1(exampleInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal("ABCDEF"))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day19.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			res, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal("ITSZCJNMUO"))
		})
	})
})
