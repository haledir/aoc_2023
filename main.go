package main

import (
	"fmt"
	_ "github.com/haledir/aoc_2023/day01"
	_ "github.com/haledir/aoc_2023/day02"
	_ "github.com/haledir/aoc_2023/day03"
	_ "github.com/haledir/aoc_2023/day04"
	"github.com/haledir/aoc_2023/day05"
	_ "github.com/haledir/aoc_2023/day06"
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
	// day03.Part1()
	// day03.Part2()
	// day04.Part1()
	// day04.Part2()
	// day05.Part1()
	day05.Part2()

	// day06.Part1()
	// day06.Part2()
}
