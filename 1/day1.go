package main

import (
	"AdventOfCode2021/common"
	"log"
	"strconv"
)

func partOne() {
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	totalIncreases := 0

	for i := 1; i < len(lines); i++ {
		line1 := lines[i-1]
		line2 := lines[i]
		num1, err1 := strconv.Atoi(line1)
		common.FatalIf(err1)
		num2, err2 := strconv.Atoi(line2)
		common.FatalIf(err2)

		if num2 > num1 {
			totalIncreases++
		}
	}

	log.Printf("Total increases: %d", totalIncreases)
}

func partTwo(windowSize int) {
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	nums := []int{}
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		common.FatalIf(err)
		nums = append(nums, num)
	}

	// get starting total
	prevTotal := 0
	for i := 0; i < windowSize; i++ {
		prevTotal += nums[i]
	}

	log.Println("first total:", prevTotal)

	windowIncreases := 0

	for i := 1; i < len(nums)-windowSize+1; i++ {
		currTotal := 0
		for j := i; j < i+windowSize; j++ {
			currTotal += nums[j]
		}
		if currTotal > prevTotal {
			windowIncreases++
		}
		prevTotal = currTotal
	}

	log.Println("Total Window Increases:", windowIncreases)
}

func main() {
	partOne()
	partTwo(3)
}
