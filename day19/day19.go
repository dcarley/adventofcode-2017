package day19

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

const (
	CharStart    = '|'
	CharJunction = '+'
	CharEmpty    = ' '
)

type Bearing struct {
	X, Y int
}

var (
	North = Bearing{0, -1}
	East  = Bearing{1, 0}
	South = Bearing{0, 1}
	West  = Bearing{-1, 0}
)

type Position struct {
	X, Y    int
	Bearing Bearing
}

func (p *Position) Move() {
	p.X += p.Bearing.X
	p.Y += p.Bearing.Y
}

type Diagram [][]byte

func (d Diagram) At(position Position) byte {
	return d[position.Y][position.X]
}

func (d Diagram) InBounds(position Position) bool {
	return position.Y >= 0 &&
		position.X >= 0 &&
		position.Y < len(d) &&
		position.X < len(d[position.Y])
}

func Parse(input io.Reader) (Diagram, error) {
	var diagram Diagram
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		buf := make([]byte, len(scanner.Bytes()))
		copy(buf, scanner.Bytes())
		diagram = append(diagram, buf)
	}

	return diagram, scanner.Err()
}

// MoveToStart scans east from 0,0 until it finds the start character.
func MoveToStart(diagram Diagram) (Position, error) {
	position := Position{
		Bearing: East,
	}

	for diagram.InBounds(position) {
		if diagram.At(position) == CharStart {
			position.Bearing = South
			return position, nil
		}

		position.Move()
	}

	return position, errors.New("start position not found")
}

// MoveToEnd moves along the diagram, collecting letters, until it reaches
// the edge of the diagram or an empty cell.
func MoveToEnd(position Position, diagram Diagram) ([]byte, error) {
	var letters []byte

	for diagram.InBounds(position) && diagram.At(position) != CharEmpty {
		if char := diagram.At(position); char >= 'A' && char <= 'Z' {
			letters = append(letters, char)
		}

		if diagram.At(position) == CharJunction {
			bearing, err := FindNewBearing(position, diagram)
			if err != nil {
				return []byte{}, err
			}
			position.Bearing = bearing
		}

		position.Move()
	}

	return letters, nil
}

// FindNewBearing finds a new direction to travel after encountering a
// junction character.
func FindNewBearing(position Position, diagram Diagram) (Bearing, error) {
	var bearings []Bearing

	// Change direction at right angles because we've already exhausted
	// backwards and forwards.
	switch position.Bearing {
	case North, South:
		bearings = []Bearing{East, West}
	case East, West:
		bearings = []Bearing{North, South}
	}

	for _, bearing := range bearings {
		// Check if the next move in the new bearing has a valid character.
		if diagram.At(Position{X: position.X + bearing.X, Y: position.Y + bearing.Y}) != CharEmpty {
			return bearing, nil
		}
	}

	return Bearing{}, errors.New("unable to find new bearing")
}

func Part1(input io.Reader) (string, error) {
	diagram, err := Parse(input)
	if err != nil {
		return "", err
	}

	position, err := MoveToStart(diagram)
	if err != nil {
		return "", err
	}

	letters, err := MoveToEnd(position, diagram)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", letters), nil
}
