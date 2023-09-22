package y22

import "advent-of-code/utils"

func Day02Part1() {
	scanner := utils.OpenDataFile("y22/d02.data")

	totalPoints := 0

	/*
		A, X = Rock 	= 1
		B, Y = Paper	= 2
		C, Z = Scissors	= 3

		Loss = 0
		Draw = 3
		Win  = 6

		A, X = 1 + 3
		A, Y = 2 + 6
		A, Z = 3 + 0

		B, X = 1 + 0
		B, Y = 2 + 3
		B, Z = 3 + 6

		C, X = 1 + 6
		C, Y = 2 + 0
		C, Z = 3 + 3
	*/

	for scanner.Scan() {
		current_line := scanner.Text()
		if current_line == "A X" {
			totalPoints += 4
		}
		if current_line == "A Y" {
			totalPoints += 8
		}
		if current_line == "A Z" {
			totalPoints += 3
		}
		if current_line == "B X" {
			totalPoints += 1
		}
		if current_line == "B Y" {
			totalPoints += 5
		}
		if current_line == "B Z" {
			totalPoints += 9
		}
		if current_line == "C X" {
			totalPoints += 7
		}
		if current_line == "C Y" {
			totalPoints += 2
		}
		if current_line == "C Z" {
			totalPoints += 6
		}
	}

	println(totalPoints)
}

func Day02Part2() {
	scanner := utils.OpenDataFile("y22/d02.data")

	totalPoints := 0

	/*
		A, X = Lose + Scissors	= 0 + 3 = 3
		A, Y = Draw + Rock	= 3 + 1 = 4
		A, Z = Win  + Paper	= 6 + 2 = 8

		B, X = Lose + Rock	= 0 + 1 = 1
		B, Y = Draw + Paper	= 3 + 2 = 5
		B, Z = Win  + Scissors	= 6 + 3 = 9

		C, X = Lose + Paper	= 0 + 2 = 2
		C, Y = Draw + Scissors	= 3 + 3 = 6
		C, Z = Win  + Rock	= 6 + 1 = 7

	*/

	for scanner.Scan() {
		current_line := scanner.Text()
		if current_line == "A X" {
			totalPoints += 3
		}
		if current_line == "A Y" {
			totalPoints += 4
		}
		if current_line == "A Z" {
			totalPoints += 8
		}
		if current_line == "B X" {
			totalPoints += 1
		}
		if current_line == "B Y" {
			totalPoints += 5
		}
		if current_line == "B Z" {
			totalPoints += 9
		}
		if current_line == "C X" {
			totalPoints += 2
		}
		if current_line == "C Y" {
			totalPoints += 6
		}
		if current_line == "C Z" {
			totalPoints += 7
		}
	}

	println(totalPoints)
}
