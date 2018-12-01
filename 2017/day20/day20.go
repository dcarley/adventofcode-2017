package day20

import (
	"fmt"
	"io"
	"math"
)

const MaxTicks = 1000

type Coord struct {
	X, Y, Z int
}
type Particle struct {
	Position, Velocity, Acceleration Coord
	Destroyed                        bool
}

func (p *Particle) Tick() {
	p.Velocity.X += p.Acceleration.X
	p.Velocity.Y += p.Acceleration.Y
	p.Velocity.Z += p.Acceleration.Z
	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y
	p.Position.Z += p.Velocity.Z
}
func (p Particle) Distance() float64 {
	return math.Abs(float64(p.Position.X)) +
		math.Abs(float64(p.Position.Y)) +
		math.Abs(float64(p.Position.Z))
}

func Parse(input io.Reader) ([]Particle, error) {
	var particles []Particle
	for {
		var particle Particle
		_, err := fmt.Fscanf(input, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>\n",
			&particle.Position.X,
			&particle.Position.Y,
			&particle.Position.Z,
			&particle.Velocity.X,
			&particle.Velocity.Y,
			&particle.Velocity.Z,
			&particle.Acceleration.X,
			&particle.Acceleration.Y,
			&particle.Acceleration.Z,
		)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			return []Particle{}, err
		}
		particles = append(particles, particle)
	}

	return particles, nil
}

type Score struct {
	Index int
	Value float64
	Found bool
}

func Nearest(particles []Particle) int {
	score := Score{}
	for i := 0; i < len(particles); i++ {
		if particles[i].Destroyed {
			continue
		}
		if dist := particles[i].Distance(); dist < score.Value || !score.Found {
			score.Index = i
			score.Value = dist
			score.Found = true
		}
	}

	return score.Index
}

func Part1(input io.Reader) (int, error) {
	particles, err := Parse(input)
	if err != nil {
		return -1, err
	}

	for tick := 0; tick < MaxTicks; tick++ {
		for i := 0; i < len(particles); i++ {
			particles[i].Tick()
		}
	}

	return Nearest(particles), nil
}

func Part2(input io.Reader) (int, error) {
	particles, err := Parse(input)
	if err != nil {
		return -1, err
	}

	for tick := 0; tick < MaxTicks; tick++ {
		positions := map[Coord]*Particle{}
		for i := 0; i < len(particles); i++ {
			if particles[i].Destroyed {
				continue
			}

			particles[i].Tick()

			position := particles[i].Position
			if existing, collided := positions[position]; collided {
				existing.Destroyed = true
				particles[i].Destroyed = true
			} else {
				positions[position] = &particles[i]
			}
		}
	}

	var notDestroyed int
	for i := 0; i < len(particles); i++ {
		if !particles[i].Destroyed {
			notDestroyed++
		}
	}

	return notDestroyed, nil
}
