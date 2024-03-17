package main

import (
	"fmt"
	// "github.com/haledir/aoc_2023/day01"
	// "github.com/haledir/aoc_2023/day02"
	"github.com/haledir/aoc_2023/day03"
)

type point struct {
	x int
	y int
}

type partNumber struct {
	start point
	end   point
}

func main() {
	fmt.Println("Start AOC")
	// day01.Part1()
	// day01.Part2()
	// day02.Part1()
	// day02.Part2()
	day03.Part1()
	day03.Part2()
}
