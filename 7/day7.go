package main

import (
	"AdventOfCode2021/common"
	"log"
	"math"
	"strconv"
	"strings"
)

// can only move horizontally
// make all of their horizonal positions match while using as little fuel as possible
// costs 1 fuel to move 1 space
// Idea: bring the stragglers in. Furthest one out comes in
// so sort the list first that way we can see who's out the furthest
type SubmarinePosition struct {
	Position int
	Count    int
}

func partOne() {
	var (
		lowestFuel = 0
		nums       = []int{}
	)

	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	split := strings.Split(lines[0], ",")
	for _, str := range split {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}

	for i := 0; i < len(nums); i++ {
		basePos := nums[i]
		fuel := 0
		for _, f := range nums {
			fuel += int(math.Abs(float64(basePos - f)))
		}

		if lowestFuel == 0 || fuel < lowestFuel {
			lowestFuel = fuel
		}
	}

	log.Println("Lowest amount of fuel needed to get all subs in same spot:", lowestFuel)
}

func partTwo() {
	var (
		lowestFuel = 0
		nums       = []int{}
		maxNum     = 0
	)

	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	split := strings.Split(lines[0], ",")
	for _, str := range split {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
		if num > maxNum {
			maxNum = num
		}
	}

	for i := 0; i < maxNum; i++ {
		basePos := i
		fuel := 0
		for _, f := range nums {
			diff := int(math.Abs(float64(basePos - f)))
			for q := 1; q <= diff; q++ {
				fuel += q
			}
		}

		log.Println("Fuel:", fuel)
		if lowestFuel == 0 || fuel < lowestFuel {
			log.Println("found a lowest fuel")
			lowestFuel = fuel
		}
	}

	log.Println("Lowest amount of fuel needed to get all subs in same spot Part 2:", lowestFuel)
}

// DIDN'T WORK :(
// func partOne() {
// 	var submarinePositionsMap = map[int]*SubmarinePosition{}
// 	var totalFuelSpent = 0
// 	// var submarinePositions = []SubmarinePosition{}
//
// 	lines, err := common.GetLinesFromFile("input.txt")
// 	common.FatalIf(err)
//
// 	split := strings.Split(lines[0], ",")
// 	for _, str := range split {
// 		num, _ := strconv.Atoi(str)
// 		if _, ok := submarinePositionsMap[num]; ok {
// 			submarinePositionsMap[num].Count = submarinePositionsMap[num].Count + 1
// 		} else {
// 			submarinePositionsMap[num] = &SubmarinePosition{
// 				Count:    1,
// 				Position: num,
// 			}
// 		}
// 	}
//
// 	submarinePositions := make([]*SubmarinePosition, 0, len(submarinePositionsMap))
// 	for _, value := range submarinePositionsMap {
// 		submarinePositions = append(submarinePositions, value)
// 	}
// 	sort.Slice(submarinePositions, func(i, j int) bool {
// 		return submarinePositions[i].Position < submarinePositions[j].Position
// 	})
//
// 	for {
// 		var (
// 			totalPos = len(submarinePositions)
// 		)
// 		if totalPos == 1 { // exit the loop if everyone is in the same spot now
// 			break
// 		}
//
// 		costToMoveFromLeft := submarinePositions[0].Count * (submarinePositions[1].Position - submarinePositions[0].Position)
// 		costToMoveFromRight := submarinePositions[totalPos-1].Count * (submarinePositions[totalPos-1].Position - submarinePositions[totalPos-2].Position)
//
// 		if costToMoveFromLeft < costToMoveFromRight {
// 			submarinePositions[1].Count += submarinePositions[0].Count
// 			totalFuelSpent += costToMoveFromLeft
// 			submarinePositions = submarinePositions[1:]
// 		} else { // cost to move from right is greater or they're equal
// 			submarinePositions[totalPos-2].Count += submarinePositions[totalPos-1].Count
// 			totalFuelSpent += costToMoveFromRight
// 			submarinePositions = submarinePositions[:totalPos-1]
// 		}
// 	}
//
// 	log.Println("Total fuel spent to get all submarines in the same position:", totalFuelSpent)
// }

//func partTwo() {
//	lines, err := common.GetLinesFromFile("input.txt")
//	common.FatalIf(err)
//}

func main() {
	partOne()
	partTwo()
}
