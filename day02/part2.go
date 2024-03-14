package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	fmt.Println("Start Day2 Part2")
	input, err := os.ReadFile("day02/part1.input")
	string_input := string(input)
	check(err)
	result := parseInputPart2(string_input)
	fmt.Printf("Day2 Part3: %v\n", result)
}

func parseInputPart2(input string) int {
	lines := strings.Split(input, "\n")
	listOfGames := make([]game, 0)
	for i, line := range lines {
		if line == "" {
			continue
		}
		var gm game
		gm.Id = i + 1
		setsString := strings.Split(line, ":")[1]
		sets := strings.Split(setsString, ";")
		var gs gameSet
		gs.Init(0, 0, 0)
		for _, set := range sets {
			setParts := strings.Split(set, ",")
			for _, c := range setParts {
				var err error
				red := 0
				blue := 0
				green := 0
				temp := strings.Split(strings.TrimSpace(c), " ")
				switch temp[1] {
				case "red":
					red, err = strconv.Atoi(temp[0])
					check(err)
					if red > gs.Red {
						gs.Red = red
					}
				case "blue":
					blue, err = strconv.Atoi(temp[0])
					check(err)
					if blue > gs.Blue {
						gs.Blue = blue
					}
				case "green":
					green, err = strconv.Atoi(temp[0])
					check(err)
					if green > gs.Green {
						gs.Green = green
					}
				}
			}
			gm.MaxSet = gs
		}
		listOfGames = append(listOfGames, gm)
	}
	result := 0
	for _, g := range listOfGames {
		result += g.MaxSet.Red * g.MaxSet.Blue * g.MaxSet.Green
	}

	return result
}
