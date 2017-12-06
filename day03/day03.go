package day03

import (
	"math"
)

const (
	BearEast = iota
	BearSouth
	BearWest
	BearNorth
)

type Position struct {
	X, Y int
}

func Part1(targetValue int) int {
	var (
		side, sideLength int
		position         Position
	)

	cellValue := 1
	for cellValue < targetValue {
		side++

		// the length increments every other side
		if side%2 == 0 {
			sideLength++
		}

		bearing := side % 4
		for moves := 0; moves < sideLength && cellValue < targetValue; moves++ {
			cellValue++
			position = Move(position, bearing, 1)
		}
	}

	// calculate manhatten distance
	return int(math.Abs(float64(position.X)) + math.Abs(float64(position.Y)))
}

func Move(position Position, bearing, moves int) Position {
	switch bearing {
	case BearNorth:
		position.Y -= moves
	case BearEast:
		position.X += moves
	case BearSouth:
		position.Y += moves
	case BearWest:
		position.X -= moves
	}

	return position
}
