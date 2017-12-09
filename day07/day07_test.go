package day07_test

import (
	"bytes"
	"os"

	. "github.com/dcarley/adventofcode-2017/day07"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day07", func() {
	Describe("Part1", func() {
		It("should solve example", func() {
			input := bytes.NewBuffer([]byte(`pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`))

			root, err := Part1(input)
			Expect(err).ToNot(HaveOccurred())
			Expect(root).To(Equal("tknk"))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day07.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			root, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(root).To(Equal("wiapj"))
		})
	})
})
