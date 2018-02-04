package day18_test

import (
	"bytes"
	"io"
	"os"

	. "github.com/dcarley/adventofcode-2017/day18"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day18", func() {
	var exampleInput io.Reader

	Describe("Part1", func() {
		BeforeEach(func() {
			exampleInput = bytes.NewBuffer([]byte(`set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2
`))
		})

		It("should solve example", func() {
			res, err := Part1(exampleInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(4))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day18.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			res, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(8600))
		})
	})

	Describe("Part2", func() {
		BeforeEach(func() {
			exampleInput = bytes.NewBuffer([]byte(`snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d
`))
		})

		It("should solve example", func() {
			res, err := Part2(exampleInput)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(3))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day18.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			res, err := Part2(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(7239))
		})
	})
})
