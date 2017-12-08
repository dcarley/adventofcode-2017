package day06

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Part1(bank []int) int {
	var (
		rounds    int
		duplicate bool
	)

	prevBanks := map[string]struct{}{}
	prevBanks[fmt.Sprintf("%v", bank)] = struct{}{}

	for !duplicate {
		rounds++
		bank = Rebalance(bank)

		key := fmt.Sprintf("%v", bank)
		if _, exists := prevBanks[key]; exists {
			duplicate = true
			break
		}
		prevBanks[key] = struct{}{}
	}

	return rounds
}

func Parse(input io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	var bank []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return []int{}, err
		}

		bank = append(bank, val)
	}

	return bank, scanner.Err()
}

type Highest struct {
	Index, Blocks int
}

func Rebalance(bank []int) []int {
	var highest Highest
	for i, blocks := range bank {
		if blocks > highest.Blocks || blocks == highest.Blocks && i < highest.Index {
			highest = Highest{
				Index:  i,
				Blocks: blocks,
			}
		}
	}

	blocks := bank[highest.Index]
	bank[highest.Index] = 0
	index := highest.Index

	for i := 0; i < blocks; i++ {
		index++
		if index >= len(bank) {
			index = 0
		}

		bank[index]++
	}

	return bank
}
