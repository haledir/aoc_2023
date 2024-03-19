package day05

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type almanacMap struct {
	begin       int
	startNumber int
	numberRange int
}

func initAlmanacMap(input []string) almanacMap {
	begin, err := strconv.Atoi(strings.TrimSpace(string(input[0])))
	check(err)
	startNumber, err := strconv.Atoi(strings.TrimSpace(string(input[1])))
	check(err)
	numberRange, err := strconv.Atoi(strings.TrimSpace(string(input[2])))
	check(err)
	return almanacMap{
		begin:       begin,
		startNumber: startNumber,
		numberRange: numberRange,
	}
}

type alamanac struct {
	seeds        []int
	seedToSoil   []almanacMap
	soilToFert   []almanacMap
	fertToWater  []almanacMap
	waterToLight []almanacMap
	lightToTemp  []almanacMap
	tempToHum    []almanacMap
	humToLoc     []almanacMap
}

func (a *alamanac) createFromInput(input string) {
	parts := strings.Split(input, "\n\n")
	seeds := strings.Split(parts[0], ":")
	seeds = strings.Split(seeds[1], " ")
	for _, s := range seeds {
		if string(s) == "" {
			continue
		}
		i, err := strconv.Atoi(strings.TrimSpace(string(s)))
		check(err)
		a.seeds = append(a.seeds, i)
	}
	a.seedToSoil = generateAlmanac(parts[1])
	a.soilToFert = generateAlmanac(parts[2])
	a.fertToWater = generateAlmanac(parts[3])
	a.waterToLight = generateAlmanac(parts[4])
	a.lightToTemp = generateAlmanac(parts[5])
	a.tempToHum = generateAlmanac(parts[6])
	pSeven := parts[7]
	a.humToLoc = generateAlmanac(pSeven[:len(pSeven)-1])
}

func generateAlmanac(input string) []almanacMap {
	inputMap := strings.Split(input, "\n")
	a := make([]almanacMap, 0)
	for i := 1; i < len(inputMap); i++ {
		entries := strings.Split(inputMap[i], " ")
		if len(entries) > 0 {
			a = append(a, initAlmanacMap(entries))
		}
	}
	return a
}

func Part1() {
	fmt.Println("Start Day5 Part1")
	input, err := os.ReadFile("day05/part1.input")
	check(err)

	string_input := string(input)

	var currentAlmanac alamanac
	currentAlmanac.createFromInput(string_input)
	// fmt.Println(currentAlmanac)

	result := -1
	for _, seed := range currentAlmanac.seeds {
		soil := checkMap(currentAlmanac.seedToSoil, seed)
		fert := checkMap(currentAlmanac.soilToFert, soil)
		water := checkMap(currentAlmanac.fertToWater, fert)
		light := checkMap(currentAlmanac.waterToLight, water)
		temp := checkMap(currentAlmanac.lightToTemp, light)
		hum := checkMap(currentAlmanac.tempToHum, temp)
		loc := checkMap(currentAlmanac.humToLoc, hum)
		if result == -1 {
			result = loc
		}
		if loc < result {
			result = loc
		}
	}
	fmt.Printf("Day5 Part1: %v\n", result)
}

func checkMap(alMap []almanacMap, value int) int {
	for i := range alMap {
		if value > alMap[i].startNumber && value <= alMap[i].startNumber+alMap[i].numberRange {
			return value - alMap[i].startNumber + alMap[i].begin
		}
	}
	return value
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
