package day13

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Firewall []int

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
			firewall = append(firewall, 0)
		}

		firewall = append(firewall, layerRange)
	}

	return firewall, scanner.Err()
}

func (f Firewall) Scan(delay int, earlyAbort bool) (severity int, caught bool) {
	for depth, layerRange := range f {
		picoSecond := delay + depth
		// https://en.wikipedia.org/wiki/Triangle_wave
		if layerRange > 0 && picoSecond%(layerRange*2-2) == 0 {
			if earlyAbort {
				return 0, true
			}

			caught = true
			severity += picoSecond * layerRange
		}
	}

	return
}

func Part1(input io.Reader) (int, error) {
	firewall, err := NewFirewall(input)
	if err != nil {
		return 0, err
	}

	severity, _ := firewall.Scan(0, false)

	return severity, nil
}

func Part2(input io.Reader) (int, error) {
	firewall, err := NewFirewall(input)
	if err != nil {
		return 0, err
	}

	var (
		delay  int
		caught bool = true
	)
	for caught {
		delay++
		_, caught = firewall.Scan(delay, true)
	}

	return delay, nil
}
