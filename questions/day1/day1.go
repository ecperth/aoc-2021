package day1

import (
	util "adventOfCode/util"
	"fmt"
)

func Solve() {
	inputStr, err := util.ScanLines("input/day1.txt")
	if err != nil {
		fmt.Println(err)
	}

	inputInt := util.StringsToInts(inputStr)

	fmt.Println(part1(inputInt))
	fmt.Println(part2(inputInt))
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
