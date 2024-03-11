package day01

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Part1() {
	fmt.Println("Start day1 part1")
	input, err := os.ReadFile("day01/input.part1")
	string_input := string(input)
	check(err)
	lines := strings.Split(string_input, "\n")
	var allInt []int = make([]int, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		var ints []int = make([]int, 0)
		for idx := range line {
			n, err := strconv.Atoi(string(line[idx]))
			if err != nil {
				continue
			}
			if len(ints) == 2 {
				ints[1] = n
				continue
			}
			ints = append(ints, n)
			ints = append(ints, n)
		}
		// fmt.Println(ints)
		res, err := strconv.Atoi(fmt.Sprint(ints[0]) + "" + fmt.Sprint(ints[1]))
		check(err)
		allInt = append(allInt, res)
	}
	// fmt.Println(allInt)
	var result int
	for _, n := range allInt {
		result += n
	}
	fmt.Printf("Part 1: %v\n", result)
}
