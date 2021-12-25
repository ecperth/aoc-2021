package util

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
