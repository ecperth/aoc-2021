package day12

import (
	"fmt"
	"github.com/ecperth/aoc-2021/util"
	"strconv"
	"strings"
	"unicode"
)

type cave struct {
	name     string
	children []*cave
}

func isCaveInSlice(s []*cave, x *cave) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

func doesSliceHaveADupe(s []*cave) bool {

	for i1, v1 := range s {
		for i2, v2 := range s {
			if i1 != i2 && v1 == v2 {
				return true
			}
		}
	}
	return false
}

var nameToCaveMap = make(map[string]*cave)

func Solve() {
	// https://adventofcode.com/2021/day/12

	inputStr, err := util.ScanLines("input/day12.txt")
	if err != nil {
		fmt.Println(err)
	}

	//Build the cave graph

	nameToCaveMap["start"] = &cave{"start", []*cave{}}
	nameToCaveMap["end"] = &cave{"end", []*cave{}}

	for _, line := range inputStr {
		AtoB := strings.Split(line,"-")

		_, bFound := nameToCaveMap[AtoB[1]]
		if !bFound{
			nameToCaveMap[AtoB[1]] = &cave{AtoB[1],[]*cave{}}
		}

		_, AFound := nameToCaveMap[AtoB[0]]
		if !AFound{
			nameToCaveMap[AtoB[0]] = &cave{AtoB[0],[]*cave{}}
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

	for _, concreteCave := range nameToCaveMap["start"].children{
		totalPaths += findPathsFromCavePart1(concreteCave, []*cave{})
	}

	return totalPaths
}

func part2() int {

	totalPaths := 0

	for _, concreteCave := range nameToCaveMap["start"].children{
		totalPaths += findPathsFromCavePart2(concreteCave, []*cave{})
	}

	return totalPaths

}

func findPathsFromCavePart1(cave *cave, visitedSmallCaves []*cave) int{

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

func findPathsFromCavePart2(cave *cave, visitedSmallCaves []*cave) int{

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