package util

import "strconv"

//TODO replace some of these methods with generics
//https://tip.golang.org/doc/tutorial/generics

func IsIntInSlice(s []int, x int) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

func IsStringInSlice(s []string, x string) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

func IsPointInSlice(s [][2]int, cord[2]int) bool {
	for _, v := range s {
		if v == cord {
			return true
		}
	}
	return false
}

func GetMaxAndMinIntFromSlice(s []int) (int, int) {
	maxValue := 0
	minValue := -1

	for _, element := range s {
		if element > maxValue {
			maxValue = element
		} else if minValue == -1 || element < minValue {
			minValue = element
		}
	}

	return maxValue, minValue
}

func NewUniformIntSlice(len, value int) []int {
	s := make([]int, len)
	for i := range s {
		s[i] = value
	}
	return s
}

func Uint8ToInt(i uint8) int {
	intValue, _ := strconv.Atoi(string(i))
	return intValue
}