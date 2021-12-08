package questions

import (
	"adventOfCode/2021/questions/day1"
	"adventOfCode/2021/questions/day2"
	"adventOfCode/2021/questions/day3"
)

func Solve(day int) {

	switch day {
	case 1:
		day1.Solve()
	case 2:
		day2.Solve()
	case 3:
		day3.Solve()
	}
}
