package main

import (
	"AdventOfCode2021/common"
	"log"
	"math"
	"sort"
)

func partOne() {
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	var (
		onesCounts = make([]int, len(lines[0]), len(lines[0])) // number of ones seen in each digit space
		cutoff     = len(lines) / 2
		gamma      = 0
		epsilon    = 0
	)

	for _, line := range lines {
		for i, char := range line {
			if char == '1' {
				onesCounts[i] = onesCounts[i] + 1
			}
		}
	}

	exponent := 0
	for i := len(onesCounts) - 1; i >= 0; i-- {
		if onesCounts[i] > cutoff {
			gamma += int(math.Pow(2, float64(exponent)))
		} else {
			epsilon += int(math.Pow(2, float64(exponent)))
		}
		exponent++
	}

	log.Printf("Gamma: %d Epsilon: %d Product: %d", gamma, epsilon, gamma*epsilon)
}

type Match struct {
	BinStr string
	Num    int
}

func partTwo() {
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	// sort in ascending order first and then
	// look for the best match with LEAST common bits
	sort.Strings(lines)

	leastCommonBitMatch := getCommonMatch(false, lines, 0)
	log.Println(leastCommonBitMatch)

	mostCommonBitMatch := getCommonMatch(true, lines, 0)
	log.Println(mostCommonBitMatch)

	exponent := 0
	leastCommonMatchNum := 0
	for i := len(leastCommonBitMatch) - 1; i >= 0; i-- {
		if leastCommonBitMatch[i] == '1' {
			leastCommonMatchNum += int(math.Pow(2, float64(exponent)))
		}
		exponent++
	}

	exponent = 0
	mostCommonMatchNum := 0
	for i := len(mostCommonBitMatch) - 1; i >= 0; i-- {
		if mostCommonBitMatch[i] == '1' {
			mostCommonMatchNum += int(math.Pow(2, float64(exponent)))
		}
		exponent++
	}

	log.Printf("Least common bit match: %d Most common bit match: %d Product: %d", leastCommonMatchNum, mostCommonMatchNum, leastCommonMatchNum*mostCommonMatchNum)
}

func getCommonMatch(mostCommon bool, nums []string, currPos int) string {
	var (
		onesCount float32
		totalNums = len(nums)
		cutoff    = float32(totalNums) / 2
	)
	// first get the counts of the 1s in each position
	for _, num := range nums {
		if num[currPos] == '1' {
			onesCount++
		}
	}

	// Grab everything that has '1' in the currPos position if:
	// - mostCommon is true AND onesCount >= cutoff
	// - mostCommon is false AND onesCount < cutoff
	// Grab everything that has '0' in the currPos position if:
	// - mostCommon is true AND onesCount < cutoff
	// - mostCOmmon is false AND onesCount >= cutoff
	newNums := []string{}
	if (mostCommon && onesCount >= cutoff) || (!mostCommon && onesCount < cutoff) {
		// get everything with '1' in currPos
		for _, num := range nums {
			if num[currPos] == '1' {
				newNums = append(newNums, num)
			}
		}
	} else if (mostCommon && onesCount < cutoff) || (!mostCommon && onesCount >= cutoff) {
		// get everything with '0' in currPos
		for _, num := range nums {
			if num[currPos] == '0' {
				newNums = append(newNums, num)
			}
		}
	}

	if len(newNums) == 1 {
		return newNums[0]
	}

	return getCommonMatch(mostCommon, newNums, currPos+1)
}

func main() {
	// partOne()
	partTwo()
}
