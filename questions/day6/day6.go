package day6

import (
	"fmt"
	"github.com/ecperth/aoc-2021/util"
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
		lanternFish = append(lanternFish, util.NewUniformIntSlice(newFish, 8)...)

	}

	return len(lanternFish)
}

func part2(input []string) int {
	//Note there must be a way of calculating for each starting point
	//how many fish will result from that fish. Just need to figure something like
	//that out and then sum.

	//Nope couldn't come up with any arithmetic way of solving. Turns out i actually
	//should be keeping track of number of fish on each day instead of the array of fish.

	lanternFish := util.StringsToInts(strings.Split(input[0], ","))

	var dayToFishCountMap = make(map[int]int)
	for _, fish := range lanternFish{
		dayToFishCountMap[fish] = dayToFishCountMap[fish] + 1
	}

	for i := 1; i <= 256; i++ {
		var nextDayToFishCountMap = make(map[int]int)
		for day, count := range dayToFishCountMap{
			if day == 0{
				nextDayToFishCountMap[6] += count
				nextDayToFishCountMap[8] += count
			} else {
				nextDayToFishCountMap[day - 1] += count
			}
		}
		dayToFishCountMap = nextDayToFishCountMap
	}

	totalFish := 0
	for _, count := range dayToFishCountMap{
		totalFish += count
	}

	return totalFish
}