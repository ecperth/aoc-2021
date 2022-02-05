package questions

import (
	"github.com/ecperth/aoc-2021/questions/day1"
	"github.com/ecperth/aoc-2021/questions/day10"
	"github.com/ecperth/aoc-2021/questions/day11"
	"github.com/ecperth/aoc-2021/questions/day12"
	"github.com/ecperth/aoc-2021/questions/day13"
	"github.com/ecperth/aoc-2021/questions/day14"
	"github.com/ecperth/aoc-2021/questions/day2"
	"github.com/ecperth/aoc-2021/questions/day3"
	"github.com/ecperth/aoc-2021/questions/day4"
	"github.com/ecperth/aoc-2021/questions/day5"
	"github.com/ecperth/aoc-2021/questions/day6"
	"github.com/ecperth/aoc-2021/questions/day7"
	"github.com/ecperth/aoc-2021/questions/day8"
	"github.com/ecperth/aoc-2021/questions/day9"
)

func Solve(day int) {

	switch day {
	case 1:
		day1.Solve()
	case 2:
		day2.Solve()
	case 3:
		day3.Solve()
	case 4:
		day4.Solve()
	case 5:
		day5.Solve()
	case 6:
		day6.Solve()
	case 7:
		day7.Solve()
	case 8:
		day8.Solve()
	case 9:
		day9.Solve()
	case 10:
		day10.Solve()
	case 11:
		day11.Solve()
	case 12:
		day12.Solve()
	case 13:
		day13.Solve()
	case 14:
		day14.Solve()
	}
}
