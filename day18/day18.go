package day18

import (
	"fmt"
	"io"
	"strconv"
	"sync"
	"time"
)

const (
	Buffer  = 64
	Timeout = 100 * time.Millisecond
)

type Instruction struct {
	Operation string
	Register  string
	Value     string
}

func Parse(input io.Reader) ([]Instruction, error) {
	var instructions []Instruction

	for {
		var inst Instruction
		args, err := fmt.Fscanf(input, "%s %s %s\n",
			&inst.Operation,
			&inst.Register,
			&inst.Value,
		)
		if err == io.EOF {
			break
		}
		if err != nil && args < 2 {
			return instructions, err
		}

		instructions = append(instructions, inst)
	}

	return instructions, nil
}

func Run(instructions []Instruction, programID int, writer chan<- int, reader <-chan int) {
	defer close(writer)

	var (
		position  int
		done      bool
		registers = map[string]int{
			"p": programID,
		}
	)

	for !done {
		inst := instructions[position]

		var value int
		if inst.Value != "" {
			var err error
			value, err = strconv.Atoi(inst.Value) // raw value
			if err != nil {
				value = registers[inst.Value] // value from another register
			}
		}

		var jump bool
		switch inst.Operation {
		case "snd":
			writer <- registers[inst.Register]
		case "rcv":
			if reader != nil {
				select {
				case v := <-reader:
					registers[inst.Register] = v
				case <-time.After(Timeout):
					done = true
				}
			} else {
				if v, ok := registers[inst.Register]; ok && v > 0 {
					done = true
				}
			}
		case "set":
			registers[inst.Register] = value
		case "add":
			registers[inst.Register] += value
		case "mul":
			registers[inst.Register] *= value
		case "mod":
			registers[inst.Register] %= value
		case "jgz":
			if v, ok := registers[inst.Register]; ok {
				jump = v > 0 // raw value
			} else if v, err := strconv.Atoi(inst.Register); err == nil {
				jump = v > 0 // value from another register
			}

			if jump {
				position += value
			}
		}

		if !jump {
			position++
		}
		if position >= len(instructions) {
			break
		}
	}

	return
}

func Part1(input io.Reader) (int, error) {
	instructions, err := Parse(input)
	if err != nil {
		return 0, err
	}

	sent := make(chan int, Buffer)
	go Run(instructions, 0, sent, nil)

	var last int
	for val := range sent {
		last = val
	}

	return last, nil
}

func Part2(input io.Reader) (int, error) {
	instructions, err := Parse(input)
	if err != nil {
		return 0, err
	}

	var (
		count int
		wg    sync.WaitGroup
		chan0 = make(chan int, Buffer)
		chan1 = make(chan int, Buffer)
		proxy = make(chan int, Buffer)
	)

	wg.Add(1)
	go func(in <-chan int, out chan<- int) {
		for val := range in {
			count++
			out <- val
		}
		close(out)
		wg.Done()
	}(proxy, chan0)

	go Run(instructions, 0, chan1, chan0)
	go Run(instructions, 1, proxy, chan1)
	wg.Wait()

	return count, nil
}
