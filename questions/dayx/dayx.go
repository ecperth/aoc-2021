package day6

import (
	util "adventOfCode/2021/util"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	// https://adventofcode.com/2021/day/6

	inputStr, err := util.ScanLines("input/day6.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1: " + strconv.Itoa(part1(inputStr)))
	fmt.Println("Part 2: " + strconv.Itoa(part2(inputStr)))
}

func part1(input []string) int {

	lanternFish := util.StringsToInts(strings.Split(input[0], ","))

	for i := 1; i <= 80; i++{
		newFish := 0
		for f, fish := range lanternFish{
			if fish == 0 {
				lanternFish[f] = 6
				newFish++
			} else {
				lanternFish[f] = fish - 1
			}

		}
		lanternFish = append(lanternFish, util.NewInts(newFish, 8)...)

	}

	return len(lanternFish)
}

func part2(input []string) int {

    return 0

}