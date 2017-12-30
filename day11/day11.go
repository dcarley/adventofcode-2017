package day11

import (
	"fmt"
	"math"
	"strings"
)

func Part1(path string) (int, error) {
	var (
		position Position
		err      error
	)

	path = strings.TrimSpace(path)
	for _, direction := range strings.Split(path, ",") {
		position, err = Move(position, direction)
		if err != nil {
			return 0, err
		}
	}

	return Distance(Position{}, position), nil
}

type Position struct {
	X, Y, Z int
}

func abs(v int) int {
	return int(math.Abs(float64(v)))
}

// https://www.redblobgames.com/grids/hexagons/#distances-cube
func Distance(a, b Position) int {
	return (abs(a.X-b.X) + abs(a.Y-b.Y) + abs(a.Z-b.Z)) / 2
}

// https://www.redblobgames.com/grids/hexagons/#coordinates-cube
func Move(position Position, direction string) (Position, error) {
	var err error

	switch direction {
	case "n":
		position.Y++
		position.Z--
	case "ne":
		position.X++
		position.Z--
	case "se":
		position.Y--
		position.X++
	case "s":
		position.Y--
		position.Z++
	case "sw":
		position.X--
		position.Z++
	case "nw":
		position.Y++
		position.X--
	default:
		err = fmt.Errorf("invalid direction: %s", direction)
	}

	return position, err
}
