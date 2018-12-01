package day06

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Part1(bank []int) int {
	return bothParts(bank, false)
}

func Part2(bank []int) int {
	return bothParts(bank, true)
}

func bothParts(bank []int, returnDiff bool) int {
	var rounds, prevRound int

	prevBanks := map[string]int{
		fmt.Sprintf("%v", bank): 1,
	}

	for prevRound == 0 {
		rounds++
		bank = Rebalance(bank)

		key := fmt.Sprintf("%v", bank)
		if round, exists := prevBanks[key]; exists {
			prevRound = round
			break
		}
		prevBanks[key] = rounds
	}

	if returnDiff {
		return rounds - prevRound
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
