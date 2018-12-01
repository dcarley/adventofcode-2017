package day20_test

import (
	"bytes"
	"io"
	"os"

	. "github.com/dcarley/adventofcode-2017/day20"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day20", func() {
	var exampleInputPart1, exampleInputPart2 io.Reader

	BeforeEach(func() {
		exampleInputPart1 = bytes.NewBuffer([]byte(`p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>
p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>
`))
		exampleInputPart2 = bytes.NewBuffer([]byte(`p=<-6,0,0>, v=<3,0,0>, a=<0,0,0>
p=<-4,0,0>, v=<2,0,0>, a=<0,0,0>
p=<-2,0,0>, v=<1,0,0>, a=<0,0,0>
p=<3,0,0>, v=<-1,0,0>, a=<0,0,0>
`))
	})

	DescribeTable("Particle tick",
		func(ticks int, positions []Coord) {
			particles, err := Parse(exampleInputPart1)
			Expect(err).ToNot(HaveOccurred())
			Expect(particles).To(HaveLen(len(positions)))

			for i := 0; i < len(particles); i++ {
				for tick := 0; tick < ticks; tick++ {
					particles[i].Tick()
				}
			}

			for i := 0; i < len(particles); i++ {
				Expect(particles[i].Position).To(Equal(positions[i]), "incorrect position for particle %d", i)
			}
		},
		Entry("0 ticks", 0, []Coord{{3, 0, 0}, {4, 0, 0}}),
		Entry("1 ticks", 1, []Coord{{4, 0, 0}, {2, 0, 0}}),
		Entry("2 ticks", 2, []Coord{{4, 0, 0}, {-2, 0, 0}}),
		Entry("3 ticks", 3, []Coord{{3, 0, 0}, {-8, 0, 0}}),
	)

	Describe("Part1", func() {
		It("should solve example", func() {
			res, err := Part1(exampleInputPart1)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(0))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day20.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			res, err := Part1(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(344))
		})
	})

	Describe("Part2", func() {
		It("should solve example", func() {
			res, err := Part2(exampleInputPart2)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(1))
		})

		It("should solve puzzle input", func() {
			file, err := os.Open("day20.input")
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()

			res, err := Part2(file)
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(Equal(404))
		})
	})
})
