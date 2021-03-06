package day16

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"unicode/utf8"
)

func Part1(input io.Reader) ([]byte, error) {
	programs := NewPrograms()
	err := bothParts(programs, input)

	return *programs, err
}

func Part2(input io.Reader) ([]byte, error) {
	const iterations = 1000000000

	buf, _ := ioutil.ReadAll(input)
	reader := bytes.NewReader(buf)

	cycle := cycleSize(reader)
	programs := NewPrograms()
	for i := 0; i < iterations%cycle; i++ {
		reader.Seek(0, io.SeekStart)
		err := bothParts(programs, reader)
		if err != nil {
			return Programs{}, err
		}
	}

	return *programs, nil
}

func cycleSize(input io.ReadSeeker) int {
	programs := NewPrograms()
	seen := map[string]struct{}{}
	for {
		if _, ok := seen[string(*programs)]; ok {
			break
		}

		seen[string(*programs)] = struct{}{}
		input.Seek(0, io.SeekStart)
		bothParts(programs, input)
	}

	return len(seen)
}

func bothParts(programs *Programs, input io.Reader) error {
	scanner := bufio.NewScanner(input)
	scanner.Split(ScanCommas)

	for scanner.Scan() {
		instruction := scanner.Bytes()
		switch instruction[0] {
		case 's':
			err := programs.Spin(instruction[1:])
			if err != nil {
				return err
			}
		case 'x': // Exchange
			err := programs.Exchange(instruction[1:])
			if err != nil {
				return err
			}
		case 'p': // Partner
			err := programs.Partner(instruction[1:])
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unable to parse instruction: %s", instruction)
		}
	}

	return scanner.Err()
}

// https://golang.org/src/bufio/scan.go?s=12782:12860#L374
func ScanCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i, width := 0, 0; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == ',' {
			return i + width, data[0:i], nil
		}
	}

	if atEOF && len(data) > 0 {
		return len(data), data[0:], nil
	}

	return 0, nil, nil
}

type Programs []byte

func NewPrograms() *Programs {
	programs := make(Programs, 16)
	for i := 0; i < len(programs); i++ {
		programs[i] = 'a' + byte(i)
	}

	return &programs
}

func (p *Programs) Spin(values []byte) error {
	size, err := strconv.Atoi(string(values))
	if err != nil {
		return err
	}

	*p = append(
		(*p)[len(*p)-int(size):len(*p)],
		(*p)[:len(*p)-int(size)]...,
	)

	return nil
}

func (p *Programs) Exchange(values []byte) error {
	var a, b int
	_, err := fmt.Fscanf(bytes.NewBuffer(values), "%d/%d", &a, &b)
	if err != nil {
		return err
	}

	(*p)[a], (*p)[b] = (*p)[b], (*p)[a]

	return nil
}

func (p *Programs) Partner(values []byte) error {
	if len(values) != 3 || values[1] != '/' {
		return fmt.Errorf("invalid partner value: %s", values)
	}
	a := bytes.IndexByte(*p, values[0])
	b := bytes.IndexByte(*p, values[2])
	if a < 0 || b < 0 {
		return fmt.Errorf("invalid partner value: %s", values)
	}

	(*p)[a], (*p)[b] = (*p)[b], (*p)[a]

	return nil
}
