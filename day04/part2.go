package day04

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type card struct {
	id           int
	luckyNumbers []int
	cardNumbers  []int
}

func Part2() {
	fmt.Println("Start Day4 Part2")
	input, err := os.ReadFile("day04/part1.input")
	check(err)

	string_input := string(input)

	lines := strings.Split(string_input, "\n")
	cards := getCards(lines)

	result := doTheStuff(cards)

	fmt.Printf("Day3 Part2: %v\n", result)
}

func doTheStuff(cards []card) int {
	amounts := make([]int, len(cards))
	for i, c := range cards {
		amounts[i]++
		hit := 1
		for _, num := range c.cardNumbers {
			if slices.Contains(c.luckyNumbers, num) {
				amounts[i+hit] += amounts[i]
				hit++
			}
		}

	}
	return sum(amounts)
}
func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func getCards(lines []string) []card {
	cards := make([]card, 0)
	for i, line := range lines {
		var c card
		if line == "" {
			continue
		}
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
		var cardNumbers []int
		for _, val := range cardNumbersStrings {
			if val == "" {
				continue
			}
			v, err := strconv.Atoi(strings.TrimSpace(val))
			check(err)
			cardNumbers = append(cardNumbers, v)
		}
		c.id = i
		c.luckyNumbers = luckyNumbers
		c.cardNumbers = cardNumbers
		cards = append(cards, c)
	}
	return cards
}
