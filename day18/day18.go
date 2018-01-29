package day18

import (
	"fmt"
	"io"
	"strconv"
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

func Part1(input io.Reader) (int, error) {
	instructions, err := Parse(input)
	if err != nil {
		return 0, err
	}

	var (
		position  int
		last      int
		registers = map[string]int{}
	)

	for {
		inst := instructions[position]

		var value int
		if inst.Value != "" {
			var err error
			value, err = strconv.Atoi(inst.Value)
			if err != nil {
				if v, ok := registers[inst.Value]; ok {
					value = v
				} else {
					return 0, err
				}
			}
		}

		var jumped bool
		switch inst.Operation {
		case "snd":
			last = registers[inst.Register]
		case "rcv":
			if v, ok := registers[inst.Register]; ok && v > 0 {
				return last, nil
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
			if registers[inst.Register] > 0 {
				position += value
				jumped = true
			}
		}

		if !jumped {
			position++
		}

		if position >= len(instructions) {
			break
		}
	}

	return last, nil
}
