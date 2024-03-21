package day06

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
	fmt.Println("Start day6 part1")
	input, err := os.ReadFile("day06/part1.input")
	string_input := string(input)
	check(err)
	lines := strings.Split(string_input, "\n")
	ms := strings.Split(strings.TrimSpace(strings.Split(lines[0], ":")[1]), "   ")
	distances := strings.Split(strings.TrimSpace(strings.Split(lines[1], ":")[1]), "  ")
	result := 1

	for i, time := range ms {
		t, err := strconv.Atoi(strings.TrimSpace(time))
		check(err)
		dist, err := strconv.Atoi(strings.TrimSpace(distances[i]))
		check(err)

		z_res := 0
		for j := 0; j <= t; j++ {
			d := j * (t - j)
			if dist < d {
				z_res++
			}
		}
		result = result * z_res
	}
	fmt.Printf("Part 1: %v\n", result)
}
