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

type Seen map[Coord][]int

func Part1(input io.Reader) (int, error) {
	seen, err := bothParts(input)
	if err != nil {
		return 0, err
	}

	var overlaps int
	for _, claims := range seen {
		if len(claims) > 1 {
			overlaps++
		}
	}

	return overlaps, nil
}

func Part2(input io.Reader) (int, error) {
	seen, err := bothParts(input)
	if err != nil {
		return 0, err
	}

	overlaps := map[int]bool{}
	for _, claims := range seen {
		if len(claims) == 1 {
			if _, ok := overlaps[claims[0]]; !ok {
				overlaps[claims[0]] = false
			}
		} else {
			for _, claim := range claims {
				overlaps[claim] = true
			}
		}
	}

	for claim, overlap := range overlaps {
		if overlap == false {
			return claim, nil
		}
	}

	return 0, fmt.Errorf("no non-overlapping claim found")
}

func bothParts(input io.Reader) (Seen, error) {
	seen := Seen{}

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
			return Seen{}, err
		}

		for y := claim.Position.Y; y < claim.Position.Y+claim.Size.Y; y++ {
			for x := claim.Position.X; x < claim.Position.X+claim.Size.X; x++ {
				seen[Coord{x, y}] = append(seen[Coord{x, y}], claim.ID)
			}
		}
	}

	return seen, nil
}
