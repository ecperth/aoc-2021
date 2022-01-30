package day12

import (
	"fmt"
	"github.com/ecperth/aoc-2021/util"
	"strconv"
	"strings"
	"unicode"
)

type Cave struct {
	name     string
	children []*Cave
}

func isCaveInSlice(s []*Cave, x *Cave) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

func doesSliceHaveADupe(s []*Cave) bool {

	for i1, v1 := range s {
		for i2, v2 := range s {
			if i1 != i2 && v1 == v2 {
				return true
			}
		}
	}
	return false
}

var nameToCaveMap = make(map[string]*Cave)

func Solve() {
	// https://adventofcode.com/2021/day/12

	inputStr, err := util.ScanLines("input/day12.txt")
	if err != nil {
		fmt.Println(err)
	}

	//Build the cave graph

	nameToCaveMap["start"] = &Cave{"start", []*Cave{}}
	nameToCaveMap["end"] = &Cave{"end", []*Cave{}}

	for _, line := range inputStr {
		AtoB := strings.Split(line,"-")

		_, bFound := nameToCaveMap[AtoB[1]]
		if !bFound{
			nameToCaveMap[AtoB[1]] = &Cave{AtoB[1],[]*Cave{}}
		}

		_, AFound := nameToCaveMap[AtoB[0]]
		if !AFound{
			nameToCaveMap[AtoB[0]] = &Cave{AtoB[0],[]*Cave{}}
		}

		nameToCaveMap[AtoB[0]].children = append(nameToCaveMap[AtoB[0]].children, nameToCaveMap[AtoB[1]])
		nameToCaveMap[AtoB[1]].children = append(nameToCaveMap[AtoB[1]].children, nameToCaveMap[AtoB[0]])
	}

	//---------------------------------

	fmt.Println("Part 1: " + strconv.Itoa(part1()))
	fmt.Println("Part 2: " + strconv.Itoa(part2()))
}

func part1() int {

	totalPaths := 0

	for _, cave := range nameToCaveMap["start"].children{
		totalPaths += findPathsFromCavePart1(cave, []*Cave{})
	}

	return totalPaths
}

func part2() int {

	totalPaths := 0

	for _, cave := range nameToCaveMap["start"].children{
		totalPaths += findPathsFromCavePart2(cave, []*Cave{})
	}

	return totalPaths

}

func findPathsFromCavePart1(cave *Cave, visitedSmallCaves []*Cave) int{

	if cave.name == "end"{
		return 1
	}
	if cave.name == "start" {
		return 0
	}

	pathsFromHere := 0

	if unicode.IsLower(rune(cave.name[0])){
		visitedSmallCaves = append(visitedSmallCaves, cave)
	}

	for _, nextCave := range cave.children{

		if !isCaveInSlice(visitedSmallCaves, nextCave){
			pathsFromHere += findPathsFromCavePart1(nextCave, visitedSmallCaves)
		}
	}

	return pathsFromHere
}

func findPathsFromCavePart2(cave *Cave, visitedSmallCaves []*Cave) int{

	if cave.name == "end"{
		return 1
	}
	if cave.name == "start" {
		return 0
	}

	pathsFromHere := 0

	if unicode.IsLower(rune(cave.name[0])){
		visitedSmallCaves = append(visitedSmallCaves, cave)
	}

	for _, nextCave := range cave.children{

		if !isCaveInSlice(visitedSmallCaves, nextCave) || !doesSliceHaveADupe(visitedSmallCaves){
			pathsFromHere += findPathsFromCavePart2(nextCave, visitedSmallCaves)
		}
	}

	return pathsFromHere
}