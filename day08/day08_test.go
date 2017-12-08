package day08_test

import (
	"bytes"
	"io"
	"os"

	. "github.com/dcarley/adventofcode-2017/day08"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day08", func() {
	var example io.Reader

	BeforeEach(func() {
		example = bytes.NewBuffer([]byte(`b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`))
	})

	DescribeTable("ParseInstruction",
		func(input string, expected Instruction) {
			instruction, err := ParseInstruction(input)
			Expect(err).ToNot(HaveOccurred())
			Expect(instruction).To(Equal(expected))
		},
		Entry("b inc 5 if a > 1", "b inc 5 if a > 1", Instruction{
			Register: "b",
			Action:   "inc",
			Value:    5,
			Condition: Condition{
				Register: "a",
				Operator: ">",
				Value:    1,
			},
		}),
		Entry("a inc 1 if b < 5", "a inc 1 if b < 5", Instruction{
			Register: "a",
			Action:   "inc",
			Value:    1,
			Condition: Condition{
				Register: "b",
				Operator: "<",
				Value:    5,
			},
		}),
		Entry("c dec -10 if a >= 1", "c dec -10 if a >= 1", Instruction{
			Register: "c",
			Action:   "dec",
			Value:    -10,
			Condition: Condition{
				Register: "a",
				Operator: ">=",
				Value:    1,
			},
		}),
		Entry("c inc -20 if c == 10", "c inc -20 if c == 10", Instruction{
			Register: "c",
			Action:   "inc",
			Value:    -20,
			Condition: Condition{
				Register: "c",
				Operator: "==",
				Value:    10,
			},
		}),
	)

	Describe("Part1", func() {
		It("should solve example", func() {
			highest, err := Part1(example)
			Expect(err).ToNot(HaveOccurred())
			Expect(highest).To(Equal(1))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day08.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			highest, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(highest).To(Equal(4888))
		})
	})

	Describe("Part2", func() {
		It("should solve example", func() {
			highest, err := Part2(example)
			Expect(err).ToNot(HaveOccurred())
			Expect(highest).To(Equal(10))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day08.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			highest, err := Part2(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(highest).To(Equal(-1))
		})
	})
})
