package day05

import (
	"fmt"
	"os"
)

func Part2() {
	fmt.Println("Start Day5 Part2")
	input, err := os.ReadFile("day05/part1.input")
	check(err)

	string_input := string(input)

	var currentAlmanac alamanac
	currentAlmanac.createFromInput(string_input)
	// fmt.Println(currentAlmanac)

	result := -1
	for i := range currentAlmanac.seeds {
		if i%2 != 0 {
			continue
		}
		for j := currentAlmanac.seeds[i]; j <= currentAlmanac.seeds[i]+currentAlmanac.seeds[i+1]; j++ {
			soil := checkMap(currentAlmanac.seedToSoil, j)
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
	}
	fmt.Printf("Day5 Part2: %v\n", result)
}
