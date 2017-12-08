package day08

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Condition struct {
	Register string
	Operator string
	Value    int
}

type Instruction struct {
	Register  string
	Action    string
	Value     int
	Condition Condition
}

func ParseInstruction(line string) (Instruction, error) {
	raw := strings.Split(line, " ")
	if len(raw) != 7 || raw[3] != "if" {
		return Instruction{}, fmt.Errorf("invalid instruction: %s", raw)
	}

	// <register> <action> <value> if <register> <operator> <value>
	instruction := Instruction{
		Register: raw[0],
		Action:   raw[1],
		Condition: Condition{
			Register: raw[4],
			Operator: raw[5],
		},
	}

	instVal, err := strconv.Atoi(raw[2])
	if err != nil {
		return Instruction{}, err
	}
	instruction.Value = instVal

	condVal, err := strconv.Atoi(raw[6])
	if err != nil {
		return Instruction{}, err
	}
	instruction.Condition.Value = condVal

	return instruction, nil
}

func CheckCondition(a, b int, operator string) bool {
	switch operator {
	case "<":
		return a < b
	case "<=":
		return a <= b
	case ">":
		return a > b
	case ">=":
		return a >= b
	case "==":
		return a == b
	case "!=":
		return a != b
	}

	return false
}

func Part1(input io.Reader) (int, error) {
	registers := map[string]int{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		instruction, err := ParseInstruction(scanner.Text())
		if err != nil {
			return 0, err
		}

		if CheckCondition(
			registers[instruction.Condition.Register],
			instruction.Condition.Value,
			instruction.Condition.Operator,
		) {
			switch instruction.Action {
			case "inc":
				registers[instruction.Register] += instruction.Value
			case "dec":
				registers[instruction.Register] -= instruction.Value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	var highest int
	for _, val := range registers {
		if val > highest {
			highest = val
		}
	}

	return highest, nil
}
