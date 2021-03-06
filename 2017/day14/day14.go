package day14

import (
	"fmt"
	"sync"

	"github.com/dcarley/adventofcode/2017/day10"
)

const Size = 128

func Part1(input []byte) int {
	grid := NewGrid(input)

	var used int
	for _, y := range grid {
		for _, x := range y {
			if x {
				used++
			}
		}
	}

	return used
}

func Part2(input []byte) int {
	grid := NewGrid(input)

	var regionID int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] {
				grid.UnsetNeighbours(Position{x, y})
				regionID++
			}
		}
	}

	return regionID
}

type Position struct {
	X, Y int
}

func (p Position) Neighbours() []Position {
	neighbours := []Position{}
	for _, offset := range []int{-1, +1} {
		if x := p.X + offset; x >= 0 && x < Size {
			neighbours = append(neighbours, Position{x, p.Y})
		}
		if y := p.Y + offset; y >= 0 && y < Size {
			neighbours = append(neighbours, Position{p.X, y})
		}
	}

	return neighbours
}

type Grid [Size][Size]bool

func NewGrid(input []byte) Grid {
	hashes := make([][]byte, Size)

	wg := sync.WaitGroup{}
	for i := 0; i < len(hashes); i++ {
		wg.Add(1)
		go func(index int) {
			suffix := fmt.Sprintf("-%d", index)
			hashes[index] = day10.Part2(append(input, []byte(suffix)...))
			wg.Done()
		}(i)
	}
	wg.Wait()

	grid := Grid{}
	for y, hash := range hashes {
		for x, char := range hash {
			var bit int
			for mask := byte(8); mask >= 1; mask /= 2 {
				if fromHexChar(char)&mask > 0 {
					grid[y][x*4+bit] = true
				}
				bit++
			}
		}
	}

	return grid
}

func (g *Grid) UnsetNeighbours(position Position) {
	g[position.Y][position.X] = false

	for _, neighbour := range position.Neighbours() {
		if g[neighbour.Y][neighbour.X] {
			g.UnsetNeighbours(neighbour)
		}
	}

	return
}

// https://golang.org/src/encoding/hex/hex.go?s=1552:1593#L73
func fromHexChar(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}

	return 0
}
