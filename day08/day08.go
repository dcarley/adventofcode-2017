package day08

import (
	"fmt"
	"io"
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
	return bothParts(input, false)
}

func Part2(input io.Reader) (int, error) {
	return bothParts(input, true)
}

func bothParts(input io.Reader, returnHighestEver bool) (int, error) {
	var highestEver int
	registers := map[string]int{}

	for {
		instruction := Instruction{}
		_, err := fmt.Fscanf(input, "%s %s %d if %s %s %d\n",
			&instruction.Register,
			&instruction.Action,
			&instruction.Value,
			&instruction.Condition.Register,
			&instruction.Condition.Operator,
			&instruction.Condition.Value,
		)
		if err == io.EOF {
			break
		}
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

			if val := registers[instruction.Register]; val > highestEver {
				highestEver = val
			}
		}
	}

	if returnHighestEver {
		return highestEver, nil
	}

	var highest int
	for _, val := range registers {
		if val > highest {
			highest = val
		}
	}

	return highest, nil
}
