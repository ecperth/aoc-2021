package day1

import (
	util "adventOfCode/util"
	"fmt"
	"strconv"
)

func Solve() {
	inputStr, err := util.ScanLines("input/day1.txt")
	if err != nil {
		fmt.Println(err)
	}

	inputInt := util.StringsToInts(inputStr)

	fmt.Println("Part 1: " + strconv.Itoa(part1(inputInt)))
	fmt.Println("Part 2: " + strconv.Itoa(part2(inputInt)))
}

func part1(input []int) int {

	increaseCounter := 0

	for i := 1; i < len(input); i++ {
		if input[i-1] < input[i] {
			increaseCounter++
		}
	}

	return increaseCounter
}

func part2(input []int) int {

	increaseCounter := 0

	for i := 3; i < len(input); i++ {
		if input[i-3] < input[i] {
			increaseCounter++
		}
	}

	return increaseCounter
}
