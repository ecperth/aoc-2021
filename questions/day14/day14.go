package day14

import (
	"fmt"
	"github.com/ecperth/aoc-2021/util"
	"strconv"
	"strings"
)

func Solve() {
	// https://adventofcode.com/2021/day/14

	inputStr, err := util.ScanLines("input/day14.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1: " + strconv.Itoa(part1(inputStr)))
	fmt.Println("Part 2: " + strconv.Itoa(part2(inputStr)))
}

type insertionRule struct{
	pos 	int
	char 	string
}

func part1(input []string) int {

	polymerTemplate := input[0]
	var insertionRules = make(map[string]string)
	for line := 2; line < len(input); line++{
		buffer := strings.Split(input[line], " -> ")
		insertionRules[buffer[0]] = buffer[1]
	}

	for step := 1; step <= 10; step++{
		pairPointer := 0
		var insertionBuffer []insertionRule
		for {
			elementPair := polymerTemplate[pairPointer:pairPointer + 2]

			insertionChar, exists := insertionRules[elementPair]
			if exists{
				insertionBuffer = append(insertionBuffer, insertionRule{pairPointer + 1, insertionChar})
			}

			pairPointer++
			if pairPointer == len(polymerTemplate) - 1{
				break
			}
		}

		for i, insertion := range insertionBuffer{
			polymerTemplate = polymerTemplate[:insertion.pos + i] + insertion.char + polymerTemplate[insertion.pos + i:]
		}
	}

	var charCounts = make(map[int32]int)
	for _, char := range polymerTemplate{
		currentVal, exists := charCounts[char]
		if exists {
			charCounts[char] = currentVal + 1
		} else{
			charCounts[char] = 1
		}
	}

	minCount, maxCount := len(polymerTemplate), 0
	for _, count := range charCounts{
		if count < minCount{
			minCount = count
		}
		if count > maxCount{
			maxCount = count
		}
	}

	return maxCount - minCount
}

func part2(input []string) int {

    return 0

}