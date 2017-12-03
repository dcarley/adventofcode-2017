package day02

import (
	"bufio"
	"bytes"
	"strconv"
)

func Part1(input []byte) (int, error) {
	var checksum int

	lines := bufio.NewScanner(bytes.NewBuffer(input))
	for lines.Scan() {
		sum, err := ChecksumLine(lines.Bytes())
		if err != nil {
			return 0, err
		}

		checksum += sum
	}

	if err := lines.Err(); err != nil {
		return 0, err
	}

	return checksum, nil
}

func ChecksumLine(line []byte) (int, error) {
	var low, high *int

	fields := bufio.NewScanner(bytes.NewBuffer(line))
	fields.Split(bufio.ScanWords)
	for fields.Scan() {
		num, err := strconv.Atoi(fields.Text())
		if err != nil {
			return 0, err
		}

		if low == nil || num < *low {
			low = &num
		}
		if high == nil || num > *high {
			high = &num
		}
	}

	if err := fields.Err(); err != nil {
		return 0, err
	}

	return *high - *low, nil
}
