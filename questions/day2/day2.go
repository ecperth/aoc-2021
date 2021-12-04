package day2

import (
	util "adventOfCode/util"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	inputStr, err := util.ScanLines("input/day2.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1: " + strconv.Itoa(part1(inputStr)))
	fmt.Println("Part 2: " + strconv.Itoa(part2(inputStr)))
}

func part1(input []string) int {

	depth := 0
	horizontal := 0

	for i := 0; i < len(input); i++ {
		command := strings.Split(input[i], " ")

		switch command[0] {
		case "forward":
			diff, _ := strconv.Atoi(command[1])
			horizontal += diff
		case "down":
			diff, _ := strconv.Atoi(command[1])
			depth += diff
		case "up":
			diff, _ := strconv.Atoi(command[1])
			depth -= diff
		}
	}

	return depth * horizontal
}

func part2(input []string) int {

	depthAim := 0

	depth := 0
	horizontal := 0

	for i := 0; i < len(input); i++ {
		command := strings.Split(input[i], " ")

		switch command[0] {
		case "forward":
			diff, _ := strconv.Atoi(command[1])
			horizontal += diff
			depth += depthAim * diff
		case "down":
			diff, _ := strconv.Atoi(command[1])
			depthAim += diff
		case "up":
			diff, _ := strconv.Atoi(command[1])
			depthAim -= diff
		}
	}

	return depth * horizontal
}
