package day8

import (
	"adventOfCode/2021/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solve() {
	// https://adventofcode.com/2021/day/8

	inputStr, err := util.ScanLines("input/day8.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1: " + strconv.Itoa(part1(inputStr)))
	fmt.Println("Part 2: " + strconv.Itoa(part2(inputStr)))
}

func part1(input []string) int {

	lookupLens := []int{2,4,3,7}

	counter := 0
	for _, line := range input {
		outputValues := strings.Split(strings.Split(line, " | ")[1], " ")
		for _, value := range outputValues {
			if util.IsIntInSlice(lookupLens, len(value)){
				counter++
			}
		}
	}

	return counter
}

func part2(input []string) int {

	//Model our displays with int -> char map
	//ie

	//dddd
	//e  a
	//e  a
	//ffff  => [0, f], [e, 1], [d, 2] ... etc clockwise ... [6, g]
	//g  b
	//g  b
	//cccc

	//For later deduction then we can say if there is an input of "ab" and
	//we know this represents a 1 then we can add [3, a], [4, b] to the map.

	//Now i need a way of mapping those map values back to an integer for later?

	//That is to say. I need to know when doing my final calculation that a
	//display with values in the 3 and 4 pos represents a 1. I guess best thing
	//to do is just define that up top...

	//Cant have varied lengths arrays or slices as keys... Just use string for now

	//2222
	//1  3
	//1  3
	//0000
	//6  4
	//6  4
	//5555

	//--------------------------------------------------------------

	//Ok so in the end i will follow this a planned order based on pretermined inference
	//steps instead of trying to generalise.

	//FOR EACH ROW:
	//1. Use set deduction logic to one by one map chars to positions in the display.
	//2. Use that mapping to determine which positions of the display are on for each output value
	//3. Use the displayToIntMap to map that to a specific int value.
	//4. Concat those ints to create the final output value for that row

	var displayToIntMap = make(map[string]string, 10)
	displayToIntMap["123456"] = "0"
	displayToIntMap["34"] = "1"
	displayToIntMap["02356"] = "2"
	displayToIntMap["02345"] = "3"
	displayToIntMap["0134"] = "4"
	displayToIntMap["01245"] = "5"
	displayToIntMap["012456"] = "6"
	displayToIntMap["234"] = "7"
	displayToIntMap["0123456"] = "8"
	displayToIntMap["012345"] = "9"

	answer := 0

	for _, line := range input {

		inputValues := strings.Split(strings.Split(line, " | ")[0], " ")
		outputValues := strings.Split(strings.Split(line, " | ")[1], " ")

		//1
		topChar := findTop(inputValues)
		topLeftChar, bottomLeftChar, middleChar := findTopLeftAndBottomLeftAndMiddle(inputValues)
		bottomChar := findBottomChar(inputValues, []string{topChar, topLeftChar, bottomLeftChar, middleChar})
		topRightChar, bottomRightChar := findTopRightAndBottomRight(inputValues, []string{topChar, topLeftChar, bottomLeftChar, middleChar, bottomChar}, middleChar)

		lineOutputString := ""

		for _, outputValue := range outputValues {
			var displayPosInts []int
			displayString := ""
			//2
			for _, outputChar := range outputValue {
				switch string(outputChar) {
				 case middleChar:
					displayPosInts = append(displayPosInts, 0)
					break
				case topLeftChar:
					displayPosInts = append(displayPosInts, 1)
					break
				case topChar:
					displayPosInts = append(displayPosInts, 2)
					break
				case topRightChar:
					displayPosInts = append(displayPosInts, 3)
					break
				case bottomRightChar:
					displayPosInts = append(displayPosInts, 4)
					break
				case bottomChar:
					displayPosInts = append(displayPosInts, 5)
					break
				case bottomLeftChar:
					displayPosInts = append(displayPosInts, 6)
					break
				}
			}
			sort.Ints(displayPosInts)
			for _, displayPos := range displayPosInts {
				displayString += strconv.Itoa(displayPos)
			}

			//3/4
			lineOutputString += displayToIntMap[displayString]
		}

		lineOutputInt, _ := strconv.Atoi(lineOutputString)
		answer += lineOutputInt

	}

    return answer
}

func findTop(inputs []string) string{

	var theOne, theSeven string

	for _, input := range inputs{
		if len(input) == 3{
			theSeven = input
		} else if len(input) == 2 {
			theOne = input
		}
	}

	for _, character := range theSeven{
		if !strings.Contains(theOne, string(character)){
			return string(character)
		}
	}

	return "donkey!"
}

func findTopLeftAndBottomLeftAndMiddle(inputs []string) (string, string, string){

	fiveLiners := make([]string, 3)
	fiveLinerCount := 0
	var theFour, theEight string

	var topLeftChar, bottomLeftChar, middleChar string

	for _, input := range inputs{
		if len(input) == 5{
			fiveLiners[fiveLinerCount] = input
			fiveLinerCount++
		} else if len(input) == 4{
			theFour = input
		} else if len(input) == 7{
			theEight = input
		}
	}

	for _, character := range theEight{
		count := 0
		for _, fiveLiner := range fiveLiners{
			if strings.Contains(fiveLiner, string(character)){
				count++
			}
		}
		if strings.Contains(theFour, string(character)){
			count++
		}
		switch count{
			case 1:
				bottomLeftChar = string(character)
				break
			case 2:
				topLeftChar = string(character)
				break
			case 4:
				middleChar = string(character)
				break
		}
	}

	return topLeftChar, bottomLeftChar, middleChar
}

func findBottomChar(inputs []string, takenChars []string) string{

	var theOne string
	for _, input := range inputs{
		if len(input) == 2{
			theOne = input
			break
		}
	}
	for _, character := range theOne {
		takenChars = append(takenChars, string(character))
	}

	for _, input := range inputs{
		if len(input) == 7{
			for _, char := range input {
				if !util.IsStringInSlice(takenChars, string(char)){
					return string(char)
				}
			}
		}
	}

	return "donkey!"
}

func findTopRightAndBottomRight(inputs, takenChars []string, middleChar string) (string, string){

	var topRightChar, bottomRightChar string

	var theOne string
	for _, input := range inputs{
		if len(input) == 2{
			theOne = input
			break
		}
	}

	sixLiners := make([]string, 3)
	sixLinerCount := 0

	for _, input := range inputs{
		if len(input) == 6{
			sixLiners[sixLinerCount] = input
			sixLinerCount++
		}
	}

	for _, char := range theOne{
		count := 0
		for _, sixLiner := range sixLiners{
			if strings.Contains(sixLiner, string(char)){
				count++
			}
		}
		if count == 2{
			topRightChar = string(char)
		} else{
			bottomRightChar = string(char)
		}
	}

	return topRightChar, bottomRightChar
}