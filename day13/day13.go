package day13

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Layer struct {
	Range     int
	Position  int
	Direction int
}

func (l *Layer) Advance() {
	if l.Range == 0 {
		return
	}

	switch l.Position {
	case 0:
		l.Direction = +1
	case l.Range - 1:
		l.Direction = -1
	}

	l.Position += l.Direction
}

func (l Layer) Caught() bool {
	return l.Range > 0 && l.Position == 0
}

func (l *Layer) Reset() {
	l.Position = 0
	l.Direction = 0
}

type Firewall []*Layer

func NewFirewall(input io.Reader) (Firewall, error) {
	var firewall Firewall

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		if len(line) != 2 {
			return firewall, fmt.Errorf("unable to parse: %s", scanner.Text())
		}

		layerDepth, err := strconv.Atoi(line[0])
		if err != nil {
			return firewall, err
		}
		layerRange, err := strconv.Atoi(line[1])
		if err != nil {
			return firewall, err
		}

		// pad for missing layers
		for i := len(firewall); i < layerDepth; i++ {
			firewall = append(firewall, nil)
		}

		firewall = append(firewall, &Layer{Range: layerRange})
	}

	return firewall, scanner.Err()
}

func (f Firewall) Scan(earlyAbort bool) (severity int, caught bool) {
	for position := 0; position < len(f); position++ {
		for depth, layer := range f {
			if layer == nil {
				continue
			}

			if position == depth && layer.Caught() {
				if earlyAbort {
					return 0, true
				}
				caught = true
				severity += depth * layer.Range
			}

			layer.Advance()
		}
	}

	return
}

func (f Firewall) Reset() {
	for _, layer := range f {
		if layer != nil {
			layer.Reset()
		}
	}
}

func Part1(input io.Reader) (int, error) {
	firewall, err := NewFirewall(input)
	if err != nil {
		return 0, err
	}

	severity, _ := firewall.Scan(false)

	return severity, nil
}

func Part2(input io.Reader) (int, error) {
	firewall, err := NewFirewall(input)
	if err != nil {
		return 0, err
	}

	var delay int
	for {
		delay++
		firewall.Reset()

		fmt.Println("delay:", delay)
		for i := 0; i < delay; i++ {
			firewall.Scan(false)
		}

		_, caught := firewall.Scan(true)
		if !caught {
			break
		}
	}

	return delay, nil
}
