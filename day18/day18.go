package day18

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"time"
)

const (
	Buffer  = 256
	Timeout = 10 * time.Millisecond
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

func getRegisterOrRaw(value string, registers map[string]int) int {
	if fromRegister, ok := registers[value]; ok {
		return fromRegister
	}

	fromRaw, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("unable to parse register value: %s", err)
	}

	return fromRaw
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
		advance := 1

		var value int
		if inst.Value != "" {
			if v, err := strconv.Atoi(inst.Value); err == nil {
				value = v // raw value
			} else {
				value = registers[inst.Value] // value from another register
			}
		}

		switch inst.Operation {
		case "snd":
			if v := getRegisterOrRaw(inst.Register, registers); v > 0 {
				writer <- v
			}
		case "rcv":
			if reader != nil {
				select {
				case v := <-reader:
					registers[inst.Register] = v
				case <-time.After(Timeout):
					done = true
				}
			} else {
				if v := getRegisterOrRaw(inst.Register, registers); v > 0 {
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
			if v := getRegisterOrRaw(inst.Register, registers); v > 0 {
				advance = value
			}
		}

		position += advance
		if position >= len(instructions) {
			done = true
		}
	}

	return
}

func ProxyCount(in <-chan int, out chan<- int) int {
	defer close(out)

	var count int
	for val := range in {
		count++
		out <- val
	}

	return count
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
		chan0 = make(chan int, Buffer)
		chan1 = make(chan int, Buffer)
		proxy = make(chan int, Buffer)
	)

	go Run(instructions, 0, chan1, chan0)
	go Run(instructions, 1, proxy, chan1)

	return ProxyCount(proxy, chan0), nil
}
