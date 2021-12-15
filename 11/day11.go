package main

import (
	"AdventOfCode2021/common"
	"log"
	"strconv"
)

// 100 octopuses in a 10x10 grid
// octopus can be seen as an int which is their energy level 0-9
// Step
// - each octopus increases by 1
// - anyone greater than 9 flashes, this increases energy level of neighbors by 1
//   An octopus can only flash ONCE per step
// - Any octopus that flashed this round goes back to 0
// How many flashes are there after 100 steps
type Octopus struct {
	Energy   int
	DidFlash bool
}

func partOne() {
	var (
		octopi       = [10][10]*Octopus{}
		totalFlashes = 0
	)
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	row := 0
	for _, line := range lines {
		col := 0
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			octopi[row][col] = &Octopus{Energy: num, DidFlash: false}
			col++
		}
		row++
		col = 0
	}

	for day := 0; day < 100; day++ {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				octopi[i][j].Energy++
			}
		}

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				bfsHelper(i, j, octopi, true)
			}
		}

		// reset values
		thisFlashCount := 0
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if octopi[i][j].DidFlash {
					thisFlashCount++
					octopi[i][j].DidFlash = false
					octopi[i][j].Energy = 0
				}
			}
		}

		totalFlashes += thisFlashCount
	}

	log.Println("Part 1 - total flashes after 100 days:", totalFlashes)
}

func bfsHelper(row int, col int, octopi [10][10]*Octopus, firstRun bool) {
	if row < 0 || row >= 10 || col < 0 || col >= 10 {
		return
	}

	if octopi[row][col].DidFlash {
		return
	}

	// This thing did not flash yet and it's a valid coordinate
	if !(firstRun) {
		octopi[row][col].Energy++
	}

	if octopi[row][col].Energy > 9 {
		octopi[row][col].DidFlash = true
		bfsHelper(row-1, col, octopi, false)
		bfsHelper(row-1, col+1, octopi, false)
		bfsHelper(row-1, col-1, octopi, false)
		bfsHelper(row, col+1, octopi, false)
		bfsHelper(row, col-1, octopi, false)
		bfsHelper(row+1, col+1, octopi, false)
		bfsHelper(row+1, col, octopi, false)
		bfsHelper(row+1, col-1, octopi, false)
	}
}

func partTwo() {
	var (
		octopi = [10][10]*Octopus{}
	)
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	row := 0
	for _, line := range lines {
		col := 0
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			octopi[row][col] = &Octopus{Energy: num, DidFlash: false}
			col++
		}
		row++
		col = 0
	}

	day := 1
	for {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				octopi[i][j].Energy++
			}
		}

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				bfsHelper(i, j, octopi, true)
			}
		}

		// reset values
		thisFlashCount := 0
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if octopi[i][j].DidFlash {
					thisFlashCount++
					octopi[i][j].DidFlash = false
					octopi[i][j].Energy = 0
				}
			}
		}

		if thisFlashCount == 100 {
			break
		}

		day++
	}

	log.Println("Part 2 - they all flash simultaneously on day:", day)
}

func main() {
	partOne()
	partTwo()
}
