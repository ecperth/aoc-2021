package util

func IsIntInSlice(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func NewUniformIntSlice(len, value int) []int {
	s := make([]int, len)
	for i := range s {
		s[i] = value
	}
	return s
}
