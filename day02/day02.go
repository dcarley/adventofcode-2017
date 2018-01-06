package day02

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

func Part1(input io.Reader) (int, error) {
	checksumFunc := func(nums []int) int {
		var low, high *int

		for i, num := range nums {
			if low == nil || num < *low {
				low = &nums[i]
			}
			if high == nil || num > *high {
				high = &nums[i]
			}
		}

		return *high - *low
	}

	return bothParts(input, checksumFunc)
}

func Part2(input io.Reader) (int, error) {
	checksumFunc := func(nums []int) int {
		for x, _ := range nums {
			for y, _ := range nums {
				if x == y {
					continue
				}

				if nums[x]%nums[y] == 0 {
					return nums[x] / nums[y]
				}
			}
		}

		return 0
	}

	return bothParts(input, checksumFunc)
}

func bothParts(input io.Reader, checksumFunc func([]int) int) (int, error) {
	var checksum int

	tableScanner := bufio.NewScanner(input)
	for tableScanner.Scan() {
		lineScanner := bufio.NewScanner(bytes.NewBuffer(tableScanner.Bytes()))
		lineScanner.Split(bufio.ScanWords)

		var nums []int
		for lineScanner.Scan() {
			num, err := strconv.Atoi(lineScanner.Text())
			if err != nil {
				return 0, err
			}

			nums = append(nums, num)
		}

		if err := lineScanner.Err(); err != nil {
			return 0, err
		}

		checksum += checksumFunc(nums)
	}

	if err := tableScanner.Err(); err != nil {
		return 0, err
	}

	return checksum, nil
}
