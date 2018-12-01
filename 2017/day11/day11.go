package day11

import (
	"fmt"
	"math"
	"strings"
)

func Part1(path string) (int, error) {
	final, _, err := bothParts(path)
	return final, err
}

func Part2(path string) (int, error) {
	_, furthest, err := bothParts(path)
	return furthest, err
}

func bothParts(path string) (final, furthest int, err error) {
	var position Position
	path = strings.TrimSpace(path)

	for _, direction := range strings.Split(path, ",") {
		err = position.Move(direction)
		if err != nil {
			return
		}
		if dist := position.Distance(Position{0, 0, 0}); dist > furthest {
			furthest = dist
		}
	}

	final = position.Distance(Position{0, 0, 0})
	return
}

func abs(v int) int {
	return int(math.Abs(float64(v)))
}

type Position struct {
	X, Y, Z int
}

// https://www.redblobgames.com/grids/hexagons/#distances-cube
func (p Position) Distance(to Position) int {
	return (abs(p.X-to.X) + abs(p.Y-to.Y) + abs(p.Z-to.Z)) / 2
}

// https://www.redblobgames.com/grids/hexagons/#coordinates-cube
func (p *Position) Move(direction string) error {
	var err error

	switch direction {
	case "n":
		p.Y++
		p.Z--
	case "ne":
		p.X++
		p.Z--
	case "se":
		p.Y--
		p.X++
	case "s":
		p.Y--
		p.Z++
	case "sw":
		p.X--
		p.Z++
	case "nw":
		p.Y++
		p.X--
	default:
		err = fmt.Errorf("invalid direction: %s", direction)
	}

	return err
}
