package main

import (
	"AdventOfCode2021/common"
	"fmt"
	"log"
	"sort"
	"strings"
)

// Insert the vals between the pairs
// After 10 steps, what is quantity of most common element minus quantity of least common element?
func partOne() {
	var (
		sequenceStr  string
		chunkMapping = map[string]string{}
	)

	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	for i, line := range lines {
		if line == "" {
			continue
		}

		if i == 0 {
			sequenceStr = line
			continue
		}

		split := strings.Split(line, " -> ")
		chunkMapping[split[0]] = split[1]
	}

	for i := 0; i < 10; i++ {
		var (
			tempStr = ""
		)

		for j := 0; j < len(sequenceStr)-1; j++ {
			lookupKey := sequenceStr[j : j+2]
			lookupVal := chunkMapping[lookupKey]
			// log.Println("Looking at lookup key:", lookupKey, "has val:", lookupVal)
			tempStr = fmt.Sprintf("%s%s%s", tempStr, string(lookupKey[0]), lookupVal)
			// log.Println("tempString:", tempStr)
			// Make sure to add the last char
			if j == len(sequenceStr)-2 {
				tempStr = fmt.Sprintf("%s%s", tempStr, string(lookupKey[1]))
			}
		}
		sequenceStr = tempStr
	}

	// Get character counts
	var charCounts = map[rune]int{}
	for _, char := range sequenceStr {
		if _, ok := charCounts[char]; ok {
			charCounts[char]++
		} else {
			charCounts[char] = 1
		}
	}
	charVals := []int{}
	for _, val := range charCounts {
		charVals = append(charVals, val)
	}
	sort.Ints(charVals)

	log.Println("Part 1 - Most common char count:", charVals[len(charVals)-1], "least common char count:", charVals[0], "difference:", charVals[len(charVals)-1]-charVals[0])
}

// Need to do it 40 times now. Previous solution wont work. The string exponentially grows too much
// Original str: ONHOOSCKBSVHBNKFKSBK
// What if we have a map of all two char possibilities to count
// first str -> ON -> 1, NH -> 1
// then every iteration, we make a new map of all possibilities
// loop over old possibilities -> for each 2 chars, we add 1 to 2char[0]newChar and newChar2char[1]
// at the end, loop over map, for each pair of 2 chars, just

// 2 maps, each iteration, we have map of all pairs AND map of char count
func partTwo() {
	var (
		sequenceStr  string
		chunkMapping = map[string]string{}
	)

	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	for i, line := range lines {
		if line == "" {
			continue
		}

		if i == 0 {
			sequenceStr = line
			continue
		}

		split := strings.Split(line, " -> ")
		chunkMapping[split[0]] = split[1]
	}
	log.Println(sequenceStr)

	var (
		charCounts  = map[string]int{}
		charPairMap = map[string]int{}
	)
	// Set the pair and char map vals up for the beginning
	for i := 0; i < len(sequenceStr)-1; i++ {
		lookupKey := sequenceStr[i : i+2]

		if _, ok := charPairMap[lookupKey]; ok {
			charPairMap[lookupKey]++
		} else {
			charPairMap[lookupKey] = 1
		}

		// add the first char of the pair to the char map
		if _, char1Ok := charCounts[string(lookupKey[0])]; char1Ok {
			charCounts[string(lookupKey[0])]++
		} else {
			charCounts[string(lookupKey[0])] = 1
		}

		// add the second pair char to the char map if we're on the last iteration
		if i == len(sequenceStr)-2 {
			if _, char1Ok := charCounts[string(lookupKey[1])]; char1Ok {
				charCounts[string(lookupKey[1])]++
			} else {
				charCounts[string(lookupKey[1])] = 1
			}
		}
	}

	for i := 0; i < 40; i++ {
		// Start with an EMPTY pair map and reuse the charCount pair map since we don't delete chars
		// For each pair key/val in the original pair map, add `val` at TWO new keys in this new pairMap
		// (each of the two new pairs it will create)
		// Also add `val` to the char map at the new character added in the middle of the new pair we're making
		var (
			newPairMap = map[string]int{}
		)

		for oldPair, oldVal := range charPairMap {
			insertChar := chunkMapping[oldPair] // the new char
			newPair1 := fmt.Sprintf("%s%s", string(oldPair[0]), insertChar)
			newPair2 := fmt.Sprintf("%s%s", insertChar, string(oldPair[1]))

			// Add first new pair to pair map
			if _, p1OK := newPairMap[newPair1]; p1OK {
				newPairMap[newPair1] += oldVal
			} else {
				newPairMap[newPair1] = oldVal
			}

			// Add second new pair to pair map
			if _, p2OK := newPairMap[newPair2]; p2OK {
				newPairMap[newPair2] += oldVal
			} else {
				newPairMap[newPair2] = oldVal
			}

			// Add the new insert char cound to the char map
			if _, charSeen := charCounts[insertChar]; charSeen {
				charCounts[insertChar] += oldVal
			} else {
				charCounts[insertChar] = oldVal
			}

			// Need to copy newPairMap over to charPairMap
			charPairMap = newPairMap
		}
	}

	charCountSlice := []int{}
	for _, val := range charCounts {
		charCountSlice = append(charCountSlice, val)
	}
	sort.Ints(charCountSlice)
	log.Println("Part 2 - Most common char count:", charCountSlice[len(charCountSlice)-1], "least common char count:", charCountSlice[0], "difference:", charCountSlice[len(charCountSlice)-1]-charCountSlice[0])

}

func main() {
	partOne()
	partTwo()
}
