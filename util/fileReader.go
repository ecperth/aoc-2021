package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ScanLines(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func ScanWords(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil
}

func StringsToInts(input []string) []int {

	output := make([]int, len(input))

	for i, s := range input {
		output[i], _ = strconv.Atoi(s)
	}

	return output
}

func StringsToIntMatrix(input []string) [][]int {
	output := make([][]int, len(input))

	for i, row := range input{
		intRow := make([]int, len(row))

		chars := strings.Split(row,"")

		for j, char := range chars{
			intRow[j], _ = strconv.Atoi(char)
		}

		output[i] = intRow
	}
	return output
}
