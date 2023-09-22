package y22

import (
	"advent-of-code/utils"
	"log"
	"slices"
	"strconv"
)

func DayOne() {
	var elf = 0
	var elves = []int{}

	scanner := utils.OpenDataFile("y22/day_one.data")
	for scanner.Scan() {
		current_line := scanner.Text()
		if current_line == "" {
			elves = append(elves, elf)
			elf = 0
		} else {
			currentLineAsNum, err := strconv.Atoi(current_line)
			if err != nil {
				log.Fatal(err)
			}
			elf += currentLineAsNum
		}
	}

	slices.Sort(elves)
	for _, elf := range elves {
		println(elf)
	}
}
