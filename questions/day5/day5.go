package day5

import (
"github.com/ecperth/aoc-2021/util"
"fmt"
"math"
"strconv"
"strings"
)

func Solve() {
	// https://adventofcode.com/2021/day/5

	inputStr, err := util.ScanLines("input/day5.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1: " + strconv.Itoa(part1(inputStr)))
	fmt.Println("Part 2: " + strconv.Itoa(part2(inputStr)))
}

func part1(input []string) int {

	var lines = make(map[[2]int]int)

	// This sucks and is hacky / non generic. Handling each line type differently
	for _, row := range input {
		startPos, endPos := parseStartAndEndCoordinate(row)
		//Ignore non horizontal or vertical
		if startPos[0] != endPos[0] && startPos[1] != endPos[1] {
			continue
		}

		if startPos[0] != endPos[0] {
			//horizontal
			y := startPos[1]
			xMin, xMax := getMinMax(startPos[0], endPos[0])
			for x := xMin; x <= xMax; x++ {
				pos := [2]int{x, y}
				currentValue, exists := lines[pos]
				if exists {
					lines[pos] = currentValue + 1
				} else {
					lines[pos] = 1
				}
			}
		} else {
			//vertical
			x := startPos[0]
			yMin, yMax := getMinMax(startPos[1], endPos[1])
			for y := yMin; y <= yMax; y++ {
				pos := [2]int{x, y}
				currentValue, exists := lines[pos]
				if exists {
					lines[pos] = currentValue + 1
				} else {
					lines[pos] = 1
				}
			}
		}
	}

	resultCounter := 0
	for _, value := range lines {
		if value >= 2 {
			resultCounter++
		}
	}

	return resultCounter

}

func part2(input []string) int {

	var lines = make(map[[2]int]int)

	for _, row := range input {
		startPos, endPos := parseStartAndEndCoordinate(row)

		dirVector := getDirVector(startPos, endPos)
		length := getLength(startPos, endPos)

		for i := 0; i <= length; i++ {

			x := startPos[0] + dirVector[0] * i
			y := startPos[1] + dirVector[1] * i

			if x < 0 || y < 0{
				fmt.Println("huh")
			}

			pos := [2]int{x, y}
			currentValue, exists := lines[pos]
			if exists {
				lines[pos] = currentValue + 1
			} else {
				lines[pos] = 1
			}
		}
	}

	resultCounter := 0
	for _, value := range lines {
		if value >= 2 {
			resultCounter++
		}
	}

	return resultCounter
}

func parseStartAndEndCoordinate(input string) ([]int, []int) {

	coordinates := strings.Split(input, " -> ")
	startPos := util.StringsToInts(strings.Split(coordinates[0], ","))
	endPos := util.StringsToInts(strings.Split(coordinates[1], ","))

	return startPos, endPos
}

func getMinMax(val1, val2 int) (int, int) {

	if val1 < val2 {
		return val1, val2
	}
	return val2, val1
}

func getDirVector(point1, point2 []int) [2]int{

	return [2]int{compareInt(point1[0], point2[0]), compareInt(point1[1], point2[1])}
}

func getLength(point1 []int, point2 []int) int{

	return int(math.Max(
		math.Abs(float64(point1[0] - point2[0])), math.Abs(float64(point1[1] - point2[1]))))
}

func compareInt(i1, i2 int) int{

	if i1 > i2{
		return -1
	} else if i1 < i2{
		return 1
	}
	return 0
}
