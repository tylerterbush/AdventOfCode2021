package main

import (
	"AdventOfCode2021/common"
	"log"
	"strconv"
	"strings"
)

func partOne() {
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	horizontalPos := 0
	depth := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		num, err := strconv.Atoi(split[1])
		common.FatalIf(err)
		switch split[0] {
		case "forward":
			horizontalPos += num
		case "down":
			depth += num
		case "up":
			depth -= num
		}
	}

	log.Printf("Horizontal: %d Depth: %d Product: %d", horizontalPos, depth, horizontalPos*depth)
}

func partTwo() {
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	horizontalPos := 0
	depth := 0
	aim := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		num, err := strconv.Atoi(split[1])
		common.FatalIf(err)
		switch split[0] {
		case "forward":
			horizontalPos += num
			depth += (aim * num)
		case "down":
			aim += num
		case "up":
			aim -= num
		}
	}

	log.Printf("Horizontal: %d Depth: %d Product: %d", horizontalPos, depth, horizontalPos*depth)
}

func main() {
	partOne()
	partTwo()
}
