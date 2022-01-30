package day7

import (
	"fmt"
	"github.com/ecperth/aoc-2021/util"
	"math"
	"strconv"
	"strings"
)

func Solve() {
	// https://adventofcode.com/2021/day/7

	inputStr, err := util.ScanLines("input/day7.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1: " + strconv.Itoa(part1(inputStr)))
	fmt.Println("Part 2: " + strconv.Itoa(part2(inputStr)))
}

func part1(input []string) int {

	crabPositions := util.StringsToInts(strings.Split(input[0], ","))
	maxPos, minPos := util.GetMaxAndMinIntFromSlice(crabPositions)
	var bestFuelUsed float64

	for pos := minPos; pos <= maxPos; pos++{
		fuelUsed := float64(0)
		for _, crab := range crabPositions{
			fuelUsed += math.Abs(float64(crab - pos))
		}
		if pos == minPos || fuelUsed < bestFuelUsed {
			bestFuelUsed = fuelUsed
		}
	}

	return int(bestFuelUsed)
}

func part2(input []string) int {

	crabPositions := util.StringsToInts(strings.Split(input[0], ","))
	maxPos, minPos := util.GetMaxAndMinIntFromSlice(crabPositions)
	var bestFuelUsed int

	//This block is me sucking at math / being too lazy
	//to work out a formula to get fuel used from steps.
	//Instead just initialise a map to use later :)
	var fuelForSteps = make(map[int]int)
	runningFuelCount := 0
	for i := 1; i <= maxPos; i++ {
		runningFuelCount += i
		fuelForSteps[i] = runningFuelCount
	}

	for pos := minPos; pos <= maxPos; pos++{
		fuelUsed := 0

		for _, crab := range crabPositions{
			steps := int(math.Abs(float64(crab - pos)))

			value, _ := fuelForSteps[steps]
			fuelUsed += value
		}
		if pos == minPos || fuelUsed < bestFuelUsed {
			bestFuelUsed = fuelUsed
		}
	}

	return bestFuelUsed

}