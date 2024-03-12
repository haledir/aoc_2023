package day01

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	fmt.Println("Start day1 part2")
	input, err := os.ReadFile("day01/part1.input")
	string_input := string(input)
	check(err)
	lines := strings.Split(string_input, "\n")
	var digits []string = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var allInt []int = make([]int, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		var ints []int = make([]int, 0)
		for idx := range line {
			n, err := strconv.Atoi(string(line[idx]))
			if err != nil {
				for idy, v := range digits {
					if (idx + len(v)) > len(line) {
						continue
					}

					if line[idx:(idx+len(v))] == v {
						// fmt.Printf("idx is %v, ", idx)
						// fmt.Printf("line[idx:idx+len(v)] is %s | ", line[idx:(idx+len(v))])
						// fmt.Printf("v is %v\n", v)
						ints = append(ints, idy)
					}

					continue
				}
				continue
			}
			ints = append(ints, n)
		}
		// fmt.Println(ints)
		res, err := strconv.Atoi(fmt.Sprint(ints[0]) + "" + fmt.Sprint(ints[len(ints)-1]))
		check(err)
		allInt = append(allInt, res)
	}
	// fmt.Println(allInt)
	var result int
	for _, n := range allInt {
		result += n
	}
	fmt.Printf("Part 2: %v\n", result)
}
