package y22

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
)

func DayOne() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path + "/y22/day_one.data")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var elf = 0
	var elves = []int{}

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
