package day03

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	fmt.Println("Start Day3 Part2")
	input, err := os.ReadFile("day03/part1.input")
	string_input := string(input)
	check(err)
	// fmt.Println(string_input)
	result := doTheMagic2(string_input)
	fmt.Printf("Day3 Part1: %v\n", result)
}

func doTheMagic2(input string) int {
	lines := strings.Split(input, "\n")
	enginePartNumbers := make([]partNumber, 0)
	theEngine := make([][]string, (len(lines) - 1))
	for y, line := range lines {
		if line == "" {
			continue
		}
		theEngine[y] = make([]string, len(line))
		isPartNumber := false
		var part partNumber
		for x, rune_value := range line {
			value := string(rune_value)
			theEngine[y][x] = value
			_, err := strconv.Atoi(value)
			if err != nil {
				if isPartNumber == false {
					continue
				}
				part.newPoint(newPoint(x-1, y), false)
				enginePartNumbers = append(enginePartNumbers, part)
				//fmt.Printf("Added Part: %v\n", part)
				part.clear()
				isPartNumber = false
				continue
			}
			if isPartNumber == false {
				part.newPoint(newPoint(x, y), true)
			}
			part.digits = append(part.digits, value)
			if x == len(line)-1 {
				part.newPoint(newPoint(x-1, y), false)
				enginePartNumbers = append(enginePartNumbers, part)
				//fmt.Printf("Added Part: %v\n", part)
				part.clear()
			}
			isPartNumber = true
		}
	}
	result := 0
	adStars := make(map[point]int)
	for _, part := range enginePartNumbers {
		legit, star := checkLegit2(part.start, part.end, theEngine)
		if legit == true {
			part.calcValue()
			val, starExists := adStars[star]
			if starExists == false {
				adStars[star] = part.value
				continue
			}
			result += part.value * val
			// fmt.Printf("Adding: %v\n", part)
		}
	}
	return result
}

func checkLegit2(start point, end point, theEngine [][]string) (bool, point) {
	for x := start.x - 1; x <= end.x+1; x++ {
		if x < 0 || x >= len(theEngine[0]) {
			continue
		}
		for y := start.y - 1; y <= start.y+1; y++ {
			if y < 0 || y >= len(theEngine) {
				continue
			}
			if checkForSymbol2(theEngine[y][x]) == true {
				return true, newPoint(x, y)
			}
		}
	}
	return false, newPoint(-1, -1)
}

func checkForSymbol2(s string) bool {
	if s != "*" {
		return false
	}
	return true
}
