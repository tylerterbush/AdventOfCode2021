package main

import (
	"AdventOfCode2021/common"
	"log"
	"strconv"
	"strings"
)

// Fish makes a new one once every 7 days
// Not synchronized. One may have 2 days left before It makes a new fish and one may have 5
// Can model a fish as an int that shows how many days it has left
// NEW lanternfish need 9 total days before they can make their first spawn, but then it's 7 after
// when a timer is 0, next day make it 6 and make a spawn with a counter of 8
// How many fish would there be after 80 days?
func getNumberOfFish(days int) {
	var (
		fishMap = map[int]int{
			-1: 0, // placeholder for fish that make new fish
			0:  0,
			1:  0,
			2:  0,
			3:  0,
			4:  0,
			5:  0,
			6:  0,
			7:  0,
			8:  0,
		}
	)

	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	split := strings.Split(lines[0], ",")
	for _, str := range split {
		num, _ := strconv.Atoi(str)
		fishMap[num]++
	}

	// Loop over the map 80 times
	// 0 -> 8 set x-1 as curr val
	// then check -1 val, add that number to 8 and 7
	for day := 0; day < days; day++ {
		for i := 0; i <= 8; i++ {
			fishMap[i-1] = fishMap[i]
		}
		fishMap[8] = fishMap[-1]
		fishMap[6] += fishMap[-1]
	}

	numFish := 0
	for i := 0; i <= 8; i++ {
		numFish += fishMap[i]
	}

	log.Printf("Total number of fish after % days: %d", days, numFish)
}

func main() {
	getNumberOfFish(80)
	getNumberOfFish(256)
}
