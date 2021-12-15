package main

import (
	"AdventOfCode2021/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Consider only horizonal lines (x1 = x2 or y1 = y2)
// At how many points do at least 2 lines overlap
func partOne() {
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	var (
		coordCounts = map[string]int{}
	)

	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(line, " -> ")
		coords1 := strings.Split(split[0], ",")
		coords2 := strings.Split(split[1], ",")
		if coords1[0] != coords2[0] && coords1[1] != coords2[1] {
			continue // ignore diagonal lines for part 1
		}

		coord1X, _ := strconv.Atoi(coords1[0])
		coord1Y, _ := strconv.Atoi(coords1[1])
		coord2X, _ := strconv.Atoi(coords2[0])
		coord2Y, _ := strconv.Atoi(coords2[1])

		if coord1X == coord2X {
			min := coord1Y
			max := coord2Y
			if coord1Y > coord2Y {
				min = coord2Y
				max = coord1Y
			}

			for i := min; i <= max; i++ {
				indexStr := fmt.Sprintf("%d,%d", coord1X, i)
				if _, ok := coordCounts[indexStr]; ok {
					coordCounts[indexStr] = coordCounts[indexStr] + 1
				} else {
					coordCounts[indexStr] = 1
				}
			}
		} else { // Y is equal
			min := coord1X
			max := coord2X
			if coord1X > coord2X {
				min = coord2X
				max = coord1X
			}

			for i := min; i <= max; i++ {
				indexStr := fmt.Sprintf("%d,%d", i, coord1Y)
				if _, ok := coordCounts[indexStr]; ok {
					coordCounts[indexStr] = coordCounts[indexStr] + 1
				} else {
					coordCounts[indexStr] = 1
				}
			}
		}
	}

	dangerousCount := 0
	for _, val := range coordCounts {
		if val >= 2 {
			dangerousCount++
		}
	}

	log.Println("Number of dangerous spots in part 1:", dangerousCount)
}

// Now allow diagonal lines (they will be 45 degrees)
func partTwo() {
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	var (
		coordCounts = map[string]int{}
	)

	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(line, " -> ")
		coords1 := strings.Split(split[0], ",")
		coords2 := strings.Split(split[1], ",")

		coord1X, _ := strconv.Atoi(coords1[0])
		coord1Y, _ := strconv.Atoi(coords1[1])
		coord2X, _ := strconv.Atoi(coords2[0])
		coord2Y, _ := strconv.Atoi(coords2[1])

		// Assume we start at coord1
		xIncrement := 1
		if coord2X < coord1X {
			xIncrement = -1
		} else if coord1X == coord2X {
			xIncrement = 0
		}
		yIncrement := 1
		if coord2Y < coord1Y {
			yIncrement = -1
		} else if coord1Y == coord2Y {
			yIncrement = 0
		}

		// We're going to update coord1 until it's equal to coord2
		for {
			indexStr := fmt.Sprintf("%d,%d", coord1X, coord1Y)
			if _, ok := coordCounts[indexStr]; ok {
				coordCounts[indexStr] = coordCounts[indexStr] + 1
			} else {
				coordCounts[indexStr] = 1
			}

			if coord1X == coord2X && coord1Y == coord2Y {
				break
			}

			coord1X += xIncrement
			coord1Y += yIncrement
		}
	}

	dangerousCount := 0
	for _, val := range coordCounts {
		if val >= 2 {
			dangerousCount++
		}
	}

	log.Println("Number of dangerous spots in part 2:", dangerousCount)
}

func main() {
	partOne()
	partTwo()
}
