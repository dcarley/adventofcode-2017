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

func Part2(targetAfterValue int) int {
	var (
		side, sideLength int
		cellValue        int
		position         Position
	)

	// pre-populate first cell because it has no neighbours yet
	cells := map[Position]int{
		position: 1,
	}

	for cellValue <= targetAfterValue {
		side++

		// the length increments every other side
		if side%2 == 0 {
			sideLength++
		}

		bearing := side % 4
		for moves := 0; moves < sideLength && cellValue <= targetAfterValue; moves++ {
			position = Move(position, bearing, 1)

			// sum neighbours
			cellValue = 0
			for _, neighbour := range Neighbours(position) {
				if val, ok := cells[neighbour]; ok {
					cellValue += val
				}
			}

			cells[position] = cellValue
		}
	}

	return cellValue
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

func Neighbours(position Position) []Position {
	neighbours := []int{-1, 0, 1}
	positions := make([]Position, 0, len(neighbours)^2-1)

	for _, xOffset := range neighbours {
		for _, yOffset := range neighbours {
			// ignore centre (self)
			if xOffset == 0 && yOffset == 0 {
				continue
			}

			positions = append(positions, Position{
				X: position.X + xOffset,
				Y: position.Y + yOffset,
			})
		}
	}

	return positions
}
