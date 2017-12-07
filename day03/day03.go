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

		var moves int
		if cellValue+sideLength <= targetValue {
			// jump to next corner
			moves = sideLength
		} else {
			// jump to target value
			moves = targetValue - cellValue
		}

		bearing := side % 4
		position = Move(position, bearing, moves)
		cellValue += moves
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
