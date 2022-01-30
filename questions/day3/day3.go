package day3

import (
	"fmt"
	"github.com/ecperth/aoc-2021/util"
	"math"
	"strconv"
)

func Solve() {
	// https://adventofcode.com/2021/day/3

	inputStr, err := util.ScanLines("input/day3.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1: " + strconv.Itoa(part1(inputStr)))
	fmt.Println("Part 2: " + strconv.Itoa(part2(inputStr)))
}

func part1(input []string) int {

	rowCount := len(input)
	bitLength := len(input[0])

	bitCounts := make([]int, bitLength)
	for i := range bitCounts {
		bitCounts[i] = 0
	}

	for i := 0; i < rowCount; i++ {
		for j, bit := range input[i] {
			if byte(bit) == '1' {
				bitCounts[j]++
			}
		}
	}

	gammaRate := 0
	epsilonRate := 0
	for i, bit := range bitCounts {
		if float64(bit)/float64(rowCount) > 0.5 {
			gammaRate += int(math.Pow(float64(2), float64(bitLength-(i+1))))
		} else {
			epsilonRate += int(math.Pow(float64(2), float64(bitLength-(i+1))))
		}
	}

	return gammaRate * epsilonRate
}

func part2(input []string) int {

	bitLength := len(input[0])

	generatorRatings := getEntriesWithMostOrLeastCommonBitAtIndex(input, 0, true)
	scrubberRatings := getEntriesWithMostOrLeastCommonBitAtIndex(input, 0, false)
	index := 1

	for len(generatorRatings) > 1 {
		generatorRatings = getEntriesWithMostOrLeastCommonBitAtIndex(generatorRatings, index, true)
		index = index + 1
	}

	index = 1
	for len(scrubberRatings) > 1 {
		scrubberRatings = getEntriesWithMostOrLeastCommonBitAtIndex(scrubberRatings, index, false)
		index = index + 1
	}

	generatorRating := 0
	scrubberRating := 0
	for i, bit := range generatorRatings[0] {
		if byte(bit) == '1' {
			generatorRating += int(math.Pow(float64(2), float64(bitLength-(i+1))))
		}
	}

	for i, bit := range scrubberRatings[0] {
		if byte(bit) == '1' {
			scrubberRating += int(math.Pow(float64(2), float64(bitLength-(i+1))))
		}
	}

	return generatorRating * scrubberRating
}

func getEntriesWithMostOrLeastCommonBitAtIndex(input []string, index int, mostCommon bool) []string {

	rowCount := len(input)
	bitCount := 0
	desiredBit := '1'
	var returnArray []string

	for i := 0; i < rowCount; i++ {
		if input[i][index] == '1' {
			bitCount++
		}
	}

	if mostCommon {
		if float64(bitCount)/float64(rowCount) < 0.5 {
			desiredBit = '0'
		}
	} else {
		if float64(bitCount)/float64(rowCount) >= 0.5 {
			desiredBit = '0'
		}
	}

	for i := 0; i < rowCount; i++ {
		if input[i][index] == byte(desiredBit) {
			returnArray = append(returnArray, input[i])
		}
	}

	return returnArray
}
