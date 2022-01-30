package main

import (
	"github.com/ecperth/aoc-2021/questions"
	"fmt"
	"time"
)

func main() {
	var day int

	fmt.Println("What day?")
	fmt.Scanln(&day)
	t := time.Now().UnixNano()
	questions.Solve(day)
	fmt.Println("compute time: ", (time.Now().UnixNano()-t)/1000, "ms")

}
