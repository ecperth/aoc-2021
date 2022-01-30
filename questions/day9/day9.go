package day9

import (
	"fmt"
	"github.com/ecperth/aoc-2021/util"
	"sort"
	"strconv"
)

func Solve() {
	// https://adventofcode.com/2021/day/9

	inputStr, err := util.ScanLines("input/day9.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1: " + strconv.Itoa(part1(inputStr)))
	fmt.Println("Part 2: " + strconv.Itoa(part2(inputStr)))
}

func part1(input []string) int {

	risk := 0

	for lineIndex, line := range input {
		for col := 0; col < len(line); col++ {
			intValue := util.Uint8ToInt(line[col])
			if col != 0 && intValue >= util.Uint8ToInt(line[col - 1]){
				continue
			}
			if lineIndex != 0 && intValue >= util.Uint8ToInt(input[lineIndex - 1][col]){
				continue
			}
			if col != len(line) - 1 && intValue >= util.Uint8ToInt(line[col + 1]){
				continue
			}
			if lineIndex != len(input) - 1 && intValue >= util.Uint8ToInt(input[lineIndex + 1][col]){
				continue
			}
			risk += 1 + intValue
		}
	}

	return risk
}

func part2(input []string) int {

	//1. I am in a basin if I am next to another basin and I am not a 9
	//2. All low points (from part 1) are in basins and all basins contain low points

	// I am going to try do this by starting from the lowpoint and recursively
	// checking surrounding points to build but the basin.

	var lowPoints [][2]int

	for lineIndex, line := range input {
		for col := 0; col < len(line); col++ {
			intValue := util.Uint8ToInt(line[col])
			if col != 0 && intValue >= util.Uint8ToInt(line[col - 1]){
				continue
			}
			if lineIndex != 0 && intValue >= util.Uint8ToInt(input[lineIndex - 1][col]){
				continue
			}
			if col != len(line) - 1 && intValue >= util.Uint8ToInt(line[col + 1]){
				continue
			}
			if lineIndex != len(input) - 1 && intValue >= util.Uint8ToInt(input[lineIndex + 1][col]){
				continue
			}
			lowPoints = append(lowPoints, [2]int{lineIndex, col})
		}
	}


	var basinSizes []int
	for _, point := range lowPoints {
		basin := [][2]int{point}
		checkedPoints := [][2]int{point}
		basin,_ = buildDaBasin(point, input, checkedPoints, basin)
		basinSizes = append(basinSizes, len(basin))
	}

	sort.Ints(basinSizes)
	basinCount := len(basinSizes)

    return basinSizes[basinCount - 1] * basinSizes[basinCount - 2] * basinSizes[basinCount - 3]
}

func buildDaBasin(point [2]int, cave []string, checkedPoints, basin [][2]int) ([][2]int, [][2]int){
	checkedPoints = append(checkedPoints, point)

	leftPoint := [2]int{point[0], point[1] - 1}
	if point[1] != 0 && !util.IsPointInSlice(checkedPoints, leftPoint) && util.Uint8ToInt(cave[point[0]][point[1] - 1]) != 9{
		basin = append(basin, leftPoint)
		basin, checkedPoints = buildDaBasin(leftPoint, cave, checkedPoints, basin)
	}

	upPoint := [2]int{point[0] - 1, point[1]}
	if point[0] != 0 && !util.IsPointInSlice(checkedPoints, upPoint) && util.Uint8ToInt(cave[point[0] - 1][point[1]]) != 9{
		basin = append(basin, upPoint)
		basin, checkedPoints = buildDaBasin(upPoint, cave, checkedPoints, basin)
	}

	rightPoint := [2]int{point[0], point[1] + 1}
	if point[1] != len(cave[0]) - 1 && !util.IsPointInSlice(checkedPoints, rightPoint) && util.Uint8ToInt(cave[point[0]][point[1] + 1]) != 9{
		basin = append(basin, rightPoint)
		basin, checkedPoints = buildDaBasin(rightPoint, cave, checkedPoints, basin)
	}

	downPoint := [2]int{point[0] + 1, point[1]}
	if point[0] != len(cave) - 1 && !util.IsPointInSlice(checkedPoints, downPoint) && util.Uint8ToInt(cave[point[0] + 1][point[1]]) != 9{
		basin = append(basin, downPoint)
		basin, checkedPoints = buildDaBasin(downPoint, cave, checkedPoints, basin)
	}

	return basin, checkedPoints
}
