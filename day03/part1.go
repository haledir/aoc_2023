package day03

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func newPoint(p_x, p_y int) point {
	return point{
		x: p_x,
		y: p_y,
	}
}

type partNumber struct {
	start  point
	end    point
	digits []string
	value  int
}

func (pN *partNumber) newPoint(p_point point, start bool) {
	if start == true {
		pN.start = p_point
	} else {
		pN.end = p_point
	}
}

func (pN *partNumber) clear() {
	pN.digits = nil
	pN.value = 0
	pN.start = newPoint(-1, -1)
	pN.end = newPoint(-1, -1)
}

func (pN *partNumber) calcValue() {
	value := ""
	for _, v := range pN.digits {

		value = value + "" + fmt.Sprintf(v)
	}
	v, err := strconv.Atoi(value)
	check(err)
	pN.value = v
}

func Part1() {
	fmt.Println("Start Day3 Part1")
	input, err := os.ReadFile("day03/part1.input")
	string_input := string(input)
	check(err)
	// fmt.Println(string_input)
	result := doTheMagic(string_input)
	fmt.Printf("Day3 Part1: %v\n", result)
}

func doTheMagic(input string) int {
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
				fmt.Printf("Added Part: %v\n", part)
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
				fmt.Printf("Added Part: %v\n", part)
				part.clear()
			}
			isPartNumber = true
		}
	}
	result := 0
	for _, part := range enginePartNumbers {
		legit := checkLegit(part.start, part.end, theEngine)
		if legit == true {
			part.calcValue()
			result += part.value
			fmt.Printf("Adding: %v\n", part)
		}
	}
	return result
}

func checkLegit(start point, end point, theEngine [][]string) bool {
	for x := start.x - 1; x <= end.x+1; x++ {
		if x < 0 || x >= len(theEngine[0]) {
			continue
		}
		for y := start.y - 1; y <= start.y+1; y++ {
			if y < 0 || y >= len(theEngine) {
				continue
			}
			if checkForSymbol(theEngine[y][x]) == true {
				return true
			}
		}
	}
	return false
}

func checkForSymbol(s string) bool {
	_, err := strconv.Atoi(s)
	if err == nil || s == "." {
		return false
	}
	return true
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
