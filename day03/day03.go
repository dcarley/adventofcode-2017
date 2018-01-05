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

func (p *Position) Move(bearing, moves int) {
	switch bearing {
	case BearNorth:
		p.Y -= moves
	case BearEast:
		p.X += moves
	case BearSouth:
		p.Y += moves
	case BearWest:
		p.X -= moves
	}
}

func (p Position) Neighbours() []Position {
	neighbours := []int{-1, 0, 1}
	positions := make([]Position, 0, len(neighbours)^2-1)

	for _, xOffset := range neighbours {
		for _, yOffset := range neighbours {
			// ignore centre (self)
			if xOffset == 0 && yOffset == 0 {
				continue
			}

			positions = append(positions, Position{
				X: p.X + xOffset,
				Y: p.Y + yOffset,
			})
		}
	}

	return positions
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
		position.Move(bearing, moves)
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
			position.Move(bearing, 1)

			// sum neighbours
			cellValue = 0
			for _, neighbour := range position.Neighbours() {
				if val, ok := cells[neighbour]; ok {
					cellValue += val
				}
			}

			cells[position] = cellValue
		}
	}

	return cellValue
}
