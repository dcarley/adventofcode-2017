package day03

import (
	"fmt"
	"io"
)

type Coord struct {
	X, Y int
}

type Claim struct {
	ID       int
	Position Coord
	Size     Coord
}

func Part1(input io.Reader) (int, error) {
	seen := map[Coord]int{}

	for {
		var claim Claim
		_, err := fmt.Fscanf(input, "#%d @ %d,%d: %dx%d\n",
			&claim.ID,
			&claim.Position.X,
			&claim.Position.Y,
			&claim.Size.X,
			&claim.Size.Y,
		)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			return 0, err
		}

		for y := claim.Position.Y; y < claim.Position.Y+claim.Size.Y; y++ {
			for x := claim.Position.X; x < claim.Position.X+claim.Size.X; x++ {
				seen[Coord{x, y}]++
			}
		}
	}

	var overlaps int
	for _, claims := range seen {
		if claims > 1 {
			overlaps++
		}
	}

	return overlaps, nil
}
