package dayXXX

import (
	util "adventOfCode/2021/util"
	"fmt"
	"strconv"
)

func Solve() {
	// https://adventofcode.com/2021/day/XXX

	inputStr, err := util.ScanLines("input/dayXXX.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1: " + strconv.Itoa(part1(inputStr)))
	fmt.Println("Part 2: " + strconv.Itoa(part2(inputStr)))
}

func part1(input []string) int {

	return 0
}

func part2(input []string) int {

    return 0

}