package day13

import (
	"fmt"
	"github.com/ecperth/aoc-2021/util"
	"strconv"
	"strings"
)

func Solve() {
	// https://adventofcode.com/2021/day/13

	inputStr, err := util.ScanLines("input/day13.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1: " + strconv.Itoa(part1(inputStr)))
	fmt.Println("Part 2:")
	part2(inputStr)
}

func part1(input []string) int {

	//If fold is on y axis. All the "points" above the fold will remain
	//All the points below become (Y axis of fold + distance from fold, x)

	//Same logic works for x axis folds and going right to left

	//Get the initial points buffered
	var pointsBuffer [][2]int
	lineIndex := 0
	for {
		if input[lineIndex] == ""{
			lineIndex++
			break
		}
		point := strings.Split(input[lineIndex], ",")
		x, _ := strconv.Atoi(point[0])
		y, _ := strconv.Atoi(point[1])
		addToBuffer(&pointsBuffer, [2]int{x, y})
		lineIndex++
	}


	instruction := strings.Split(strings.TrimLeft(input[lineIndex], "fold along "), "=")
	depth, _ := strconv.Atoi(instruction[1])

	var newPointsBuffer [][2]int
	if instruction[0] == "y"{
		for _, point := range pointsBuffer{
			if point[1] < depth {
				addToBuffer(&newPointsBuffer, point)
			} else if point[1] > depth {
				point = [2]int{point[0], depth - (point[1] - depth)}
				addToBuffer(&newPointsBuffer, point)
			}
		}
	} else if instruction[0] == "x"{
		for _, point := range pointsBuffer{
			if point[0] < depth {
				addToBuffer(&newPointsBuffer, point)
			} else if point[0] > depth {
				point = [2]int{depth - (point[0] - depth), point[1]}
				addToBuffer(&newPointsBuffer, point)
			}
		}
	}

	return len(newPointsBuffer)
}

func addToBuffer(buffer *[][2]int , newPoint [2]int){

	pointExists := isPointInBuffer(buffer,newPoint)

	if !pointExists{
		*buffer = append(*buffer, newPoint)
	}
}

func isPointInBuffer(buffer *[][2]int , newPoint [2]int) bool{
	for _, point := range *buffer {
		if point == newPoint{
			return true
		}
	}
	return false
}

func part2(input []string) {

	var pointsBuffer [][2]int
	lineIndex := 0
	for {
		if input[lineIndex] == ""{
			lineIndex++
			break
		}
		point := strings.Split(input[lineIndex], ",")
		x, _ := strconv.Atoi(point[0])
		y, _ := strconv.Atoi(point[1])
		addToBuffer(&pointsBuffer, [2]int{x, y})
		lineIndex++
	}

	for i := lineIndex; i < len(input); i++{
		instruction := strings.Split(strings.TrimLeft(input[i], "fold along "), "=")
		depth, _ := strconv.Atoi(instruction[1])

		var newPointsBuffer [][2]int
		if instruction[0] == "y"{
			for _, point := range pointsBuffer{
				if point[1] < depth {
					addToBuffer(&newPointsBuffer, point)
				} else if point[1] > depth {
					point = [2]int{point[0], depth - (point[1] - depth)}
					addToBuffer(&newPointsBuffer, point)
				}
			}
		} else if instruction[0] == "x" {
			for _, point := range pointsBuffer {
				if point[0] < depth {
					addToBuffer(&newPointsBuffer, point)
				} else if point[0] > depth {
					point = [2]int{depth - (point[0] - depth), point[1]}
					addToBuffer(&newPointsBuffer, point)
				}
			}
		}

		pointsBuffer = newPointsBuffer
	}

	biggestX, biggestY := 0, 0
	for _, point := range pointsBuffer{
		if point[0] > biggestX{
			biggestX = point[0]
		}
		if point[1] > biggestY{
			biggestY = point[1]
		}
	}

	//Lets cheat and assume we got a
	for row := 0; row <= biggestY; row++{
		for col := 0; col <= biggestX; col++{
			if isPointInBuffer(&pointsBuffer, [2]int{col, row}){
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}