package day11

import (
	"fmt"
	"github.com/ecperth/aoc-2021/util"
	"strconv"
)

var input [][]int

func Solve() {
	// https://adventofcode.com/2021/day/11

	inputStr, err := util.ScanLines("input/day11.txt")
	if err != nil {
		fmt.Println(err)
	}

	input = util.StringsToIntMatrix(inputStr)
	fmt.Println("Part 1: " + strconv.Itoa(part1()))

	input = util.StringsToIntMatrix(inputStr)
	fmt.Println("Part 2: " + strconv.Itoa(part2()))
}

func part1() int {

	flashCount := 0

	for step := 1; step <= 100; step++{

		var flashedThisTurn [][2]int

		for rowIndex, row := range input{
			for colIndex := range row {
				flashedThisTurn = incrementPussy(flashedThisTurn, colIndex, rowIndex)
			}
		}

		flashCount += len(flashedThisTurn)
	}

	return flashCount
}

func part2() int {

	stepCount := 0

	for {

		stepCount++

		var flashedThisTurn [][2]int

		for rowIndex, row := range input{
			for colIndex := range row {
				flashedThisTurn = incrementPussy(flashedThisTurn, colIndex, rowIndex)
			}
		}

		if len(flashedThisTurn) == len(input) * len(input[0]){
			return stepCount
		}

	}

}

func incrementPussy(flashedThisTurn [][2]int, colIndex, rowIndex int) [][2]int{

	if util.IsPointInSlice(flashedThisTurn, [2]int{colIndex, rowIndex}){
		return flashedThisTurn
	}

	if input[rowIndex][colIndex] == 9{

		input[rowIndex][colIndex] = 0
		flashedThisTurn = append(flashedThisTurn, [2]int{colIndex, rowIndex})

		for j := -1; j <= 1; j++{
			if rowIndex + j < 0 || rowIndex + j >= len(input){
				continue
			}
			for i := -1; i <= 1; i++{
				if colIndex + i < 0 || colIndex + i >= len(input[rowIndex]){
					continue
				}

				flashedThisTurn = incrementPussy(flashedThisTurn, colIndex + i, rowIndex + j)
			}
		}
	} else{
		input[rowIndex][colIndex] = input[rowIndex][colIndex] + 1
	}

	return flashedThisTurn
}