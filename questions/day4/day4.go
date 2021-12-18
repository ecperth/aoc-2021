package day4

import (
	util "adventOfCode/2021/util"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	// https://adventofcode.com/2021/day/4

	inputStr, err := util.ScanLines("input/day4.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1: " + strconv.Itoa(part1(inputStr)))
	fmt.Println("Part 2: " + strconv.Itoa(part2(inputStr)))
}

func part1(input []string) int {

	numbers := strings.Split(input[0], ",")
	boards := initialiseBoards(input)

	for _, number := range numbers {

		for boardIndex := 0; boardIndex < len(boards); boardIndex++ {
			searchBoardForNumber(boards[boardIndex], number)
		}

		//CHECK FOR WINNER
		for boardIndex := 0; boardIndex < len(boards); boardIndex++ {

			if checkCols(boards[boardIndex]) {
				numberInt, _ := strconv.Atoi(number)
				return calculateScore(boards[boardIndex], numberInt)
			}

			if checkRows(boards[boardIndex]) {
				numberInt, _ := strconv.Atoi(number)
				return calculateScore(boards[boardIndex], numberInt)
			}

		}
	}

	return 0

}

func part2(input []string) int {

	numbers := strings.Split(input[0], ",")
	boards := initialiseBoards(input)
	var winningBoards []int

	for _, number := range numbers {
		for boardIndex := 0; boardIndex < len(boards); boardIndex++ {

			if util.IsIntInSlice(winningBoards, boardIndex) {
				continue
			}

			searchBoardForNumber(boards[boardIndex], number)
		}

		//CHECK FOR WINNER
		for boardIndex := 0; boardIndex < len(boards); boardIndex++ {

			if util.IsIntInSlice(winningBoards, boardIndex) {
				continue
			}

			if checkCols(boards[boardIndex]) {
				winningBoards = append(winningBoards, boardIndex)
				if len(winningBoards) == len(boards) {
					numberInt, _ := strconv.Atoi(number)
					return calculateScore(boards[boardIndex], numberInt)
				}
				continue
			}

			if checkRows(boards[boardIndex]) {
				winningBoards = append(winningBoards, boardIndex)
				if len(winningBoards) == len(boards) {
					numberInt, _ := strconv.Atoi(number)
					return calculateScore(boards[boardIndex], numberInt)
				}
			}
		}
	}

	return 0
}

func initialiseBoards(input []string) [][][]string {

	var boards [][][]string
	var boardBuffer [][]string

	for i := 2; i < len(input); i++ {

		if i%6 == 1 {
			continue
		}

		row := strings.Fields(input[i])
		boardBuffer = append(boardBuffer, row)

		if i%6 == 0 {
			boards = append(boards, boardBuffer)
			boardBuffer = nil
		}
	}

	return boards
}

func searchBoardForNumber(board [][]string, number string) {
	for rowIndex := 0; rowIndex < len(board); rowIndex++ {
		for valueIndex, value := range board[rowIndex] {
			if value == number {
				board[rowIndex][valueIndex] = "X"
			}
		}
	}
}

func checkCols(board [][]string) bool {
	for colIndex := 0; colIndex < len(board[0]); colIndex++ {
		winningBoard := true
		for rowIndex := 0; rowIndex < len(board[colIndex]); rowIndex++ {
			if board[rowIndex][colIndex] != "X" {
				winningBoard = false
				break
			}
		}
		if winningBoard {
			return true
		}
	}
	return false
}

func checkRows(board [][]string) bool {
	for rowIndex := 0; rowIndex < len(board); rowIndex++ {
		winningBoard := true
		for colIndex := 0; colIndex < len(board[rowIndex]); colIndex++ {
			if board[rowIndex][colIndex] != "X" {
				winningBoard = false
				break
			}
		}
		if winningBoard {
			return true
		}
	}
	return false
}

func calculateScore(board [][]string, number int) int {

	score := 0

	for rowIndex := 0; rowIndex < len(board); rowIndex++ {
		for _, value := range board[rowIndex] {
			if value != "X" {
				numberValue, _ := strconv.Atoi(value)
				score += numberValue
			}
		}
	}

	return score * number
}
