package day10

import (
	"fmt"
	"github.com/ecperth/aoc-2021/util"
	"sort"
	"strconv"
)

func Solve() {
	// https://adventofcode.com/2021/day/10

	inputStr, err := util.ScanLines("input/day10.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1: " + strconv.Itoa(part1(inputStr)))
	fmt.Println("Part 2: " + strconv.Itoa(part2(inputStr)))
}




func part1(input []string) int {

	openBrackets := []string{"[", "(", "{", "<"}
	var bracketMap = make(map[string]string)
	bracketMap["]"] = "["
	bracketMap[")"] = "("
	bracketMap["}"] = "{"
	bracketMap[">"] = "<"

	var bracketPointsMap = make(map[string]int)
	bracketPointsMap["]"] = 57
	bracketPointsMap[")"] = 3
	bracketPointsMap["}"] = 1197
	bracketPointsMap[">"] = 25137

	var charStack, errorChars []string

	for _, line := range input {
		for _, char := range line {
			bracket := string(char)
			if util.IsStringInSlice(openBrackets, bracket){
				charStack = append(charStack, bracket)
			} else {
				if bracketMap[bracket] != charStack[len(charStack) - 1]{
					errorChars = append(errorChars, bracket)
					break
				} else {
					charStack = charStack[:len(charStack)-1]
				}
			}
		}
	}

	syntaxErrorScore := 0
	for _, errorChar := range errorChars {
		syntaxErrorScore += bracketPointsMap[errorChar]
	}

	return syntaxErrorScore
}

func part2(input []string) int {

	openBrackets := []string{"[", "(", "{", "<"}
	var bracketMap = make(map[string]string)
	bracketMap["]"] = "["
	bracketMap[")"] = "("
	bracketMap["}"] = "{"
	bracketMap[">"] = "<"

	var reverseBracketMap = make(map[string]string)
	for key := range bracketMap {
		reverseBracketMap[bracketMap[key]] = key
	}

	var bracketPointsMap = make(map[string]int)
	bracketPointsMap["("] = 1
	bracketPointsMap["["] = 2
	bracketPointsMap["{"] = 3
	bracketPointsMap["<"] = 4

	var autoCompleteScores []int

	for _, line := range input {
		var charStack []string
		lineCorrupt := false
		for _, char := range line {
			bracket := string(char)
			if util.IsStringInSlice(openBrackets, bracket){
				charStack = append(charStack, bracket)
			} else {
				if bracketMap[bracket] != charStack[len(charStack) - 1]{
					lineCorrupt = true
					break
				}
				charStack = charStack[:len(charStack)-1]
			}
		}

		if lineCorrupt{
			continue
		}

		autoCompleteScore := 0
		for index := range charStack {
			autoCompleteScore = autoCompleteScore * 5 + bracketPointsMap[charStack[len(charStack) - 1 - index]]
		}
		autoCompleteScores = append(autoCompleteScores, autoCompleteScore)
	}

	sort.Ints(autoCompleteScores)

	return autoCompleteScores[(len(autoCompleteScores) - 1) / 2]

}