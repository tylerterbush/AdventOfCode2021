package main

import (
	"AdventOfCode2021/common"
	"log"
	"sort"
)

var (
	characterScores = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	desiredOpening = map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
		">": "<",
	}
)

// Basically just a bracket opening/closing question
// Corrupted lines = a chunk closes with the wrong character
// Find the first illegal character on the line
//   ): 3 points.
//   ]: 57 points.
//   }: 1197 points.
//   >: 25137 points.
// What's the total error score of all the lines?
func partOne() {
	var (
		errorScore = 0
	)

	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	for _, line := range lines {
		var (
			lineSlice = []string{}
		)

		if line == "" {
			continue
		}

		done := false
		for _, char := range line {
			log.Println(lineSlice)
			switch string(char) {
			case ")", "}", "]", ">":
				recentChar := lineSlice[len(lineSlice)-1]
				if recentChar != desiredOpening[string(char)] {
					errorScore += characterScores[string(char)]
					done = true
				}
				lineSlice = lineSlice[0 : len(lineSlice)-1]
			default:
				lineSlice = append(lineSlice, string(char))
			}
			if done {
				break
			}
		}
	}

	log.Println("Part 1 - total error score:", errorScore)
}

// Now we're only concerned with the incomplete lines
// Start with a total score of 0. Then, for each character,
// multiply the total score by 5 and then increase the total
//  score by the point value given for the character in the following table:
/*
   ): 1 point.
   ]: 2 points.
   }: 3 points.
   >: 4 points.
*/
// Find the MIDDLE score for all of the incomplete lines
var (
	partTwoScoreMap = map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}
)

func partTwo() {
	var (
		incompleteScores = []int{}
	)

	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	for _, line := range lines {
		var (
			lineSlice = []string{}
		)

		errorFound := false
		for _, char := range line {
			switch string(char) {
			case ")", "}", "]", ">":
				recentChar := lineSlice[len(lineSlice)-1]
				if recentChar != desiredOpening[string(char)] {
					errorFound = true
				}
				lineSlice = lineSlice[0 : len(lineSlice)-1]
			default:
				lineSlice = append(lineSlice, string(char))
			}
			if errorFound {
				break
			}
		}

		if errorFound {
			continue
		}

		// This is an incomplete line
		thisScore := 0
		for i := len(lineSlice) - 1; i >= 0; i-- {
			thisScore = (thisScore * 5) + partTwoScoreMap[string(lineSlice[i])]
		}
		log.Println("this score:", thisScore)
		incompleteScores = append(incompleteScores, thisScore)
	}

	sort.Ints(incompleteScores)
	log.Println("Part 2 - the middle score for all of the incomplete lines:", incompleteScores[len(incompleteScores)/2])
}

func main() {
	// partOne()
	partTwo()
}
