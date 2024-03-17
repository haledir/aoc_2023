package day04

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	fmt.Println("Start Day4 Part1")
	input, err := os.ReadFile("day04/part1.input")
	check(err)

	string_input := string(input)

	lines := strings.Split(string_input, "\n")

	result := 0
	// fmt.Println(string_input)
	for _, line := range lines {
		if line == "" {
			continue
		}
		result += getCardInfo(line)
	}

	fmt.Printf("Day3 Part1: %v\n", result)
}

func getCardInfo(line string) int {
	input := strings.Split(line, ":")[1]
	parts := strings.Split(input, "|")
	luckyNumbersStrings := strings.Split(parts[0], " ")
	var luckyNumbers []int
	for _, val := range luckyNumbersStrings {
		if val == "" {
			continue
		}
		v, err := strconv.Atoi(strings.TrimSpace(val))
		check(err)
		luckyNumbers = append(luckyNumbers, v)
	}
	cardNumbersStrings := strings.Split(parts[1], " ")
	var amount float64 = 0
	for _, val := range cardNumbersStrings {
		if val == "" {
			continue
		}
		v, err := strconv.Atoi(strings.TrimSpace(val))
		check(err)
		if slices.Contains(luckyNumbers, v) {
			// fmt.Printf("%v is in %v\n", v, luckyNumbers)
			amount++
		}
	}
	if amount == 0 {
		return int(amount)
	}
	amount--
	res := int(math.Pow(2, amount))

	// fmt.Printf("Line: %v - has amount: %v with res: %v\n", input, amount, res)
	return res
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
