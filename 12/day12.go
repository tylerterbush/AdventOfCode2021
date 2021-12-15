package main

import (
	"AdventOfCode2021/common"
	"log"
	"strings"
	"unicode"
)

// find the number of paths that start at 'start' and end at 'end'
// and don't visit small caves more than once
// there are Big caves and Small caves (capitalization of letter)
// Can visit big caves multiple times
type Cave struct {
	Name          string
	Small         bool
	AdjacentCaves []string
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func partOne() {
	var (
		// map of cave name -> cave struct (size, times visited, adjacent cave names)
		caves = map[string]*Cave{}
	)
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	for _, line := range lines {
		split := strings.Split(line, "-")
		if _, ok := caves[split[0]]; ok {
			caves[split[0]].AdjacentCaves = append(caves[split[0]].AdjacentCaves, split[1])
		} else {
			caves[split[0]] = &Cave{
				Name:          split[0],
				Small:         !IsUpper(split[0]),
				AdjacentCaves: []string{split[1]},
			}
		}

		if _, ok := caves[split[1]]; ok {
			caves[split[1]].AdjacentCaves = append(caves[split[1]].AdjacentCaves, split[0])
		} else {
			caves[split[1]] = &Cave{
				Name:          split[1],
				Small:         !IsUpper(split[1]),
				AdjacentCaves: []string{split[0]},
			}
		}
	}

	// start at start
	// loop over adjacent caves
	numPaths := recursiveHelper("start", caves, map[string]int{})
	log.Println("Part 1 - number of valid paths:", numPaths)
}

func recursiveHelper(caveName string, caves map[string]*Cave, visits map[string]int) int {
	cave := caves[caveName] // current cave
	if cave.Name == "end" {
		return 1
	}

	if _, ok := visits[caveName]; ok {
		visits[caveName] = visits[caveName] + 1
	} else {
		visits[caveName] = 1
	}
	pathCount := 0
	for _, cName := range cave.AdjacentCaves {
		adjCave := caves[cName]
		// see if we can go this way
		adjCaveVisits := 0
		if _, ok := visits[adjCave.Name]; ok {
			adjCaveVisits = visits[adjCave.Name]
		}
		if !adjCave.Small || (adjCave.Small && adjCaveVisits == 0) {
			// we can go this way
			newMap := map[string]int{}
			for k, v := range visits {
				newMap[k] = v
			}
			pathCount += recursiveHelper(adjCave.Name, caves, newMap)
		}
	}
	return pathCount
}

func recursiveHelper2(caveName string, caves map[string]*Cave, visits map[string]int, smallCaveVisitedTwice bool) int {
	cave := caves[caveName] // current cave
	if cave.Name == "end" {
		return 1
	}

	if _, ok := visits[caveName]; ok {
		visits[caveName] = visits[caveName] + 1
	} else {
		visits[caveName] = 1
	}
	pathCount := 0
	for _, cName := range cave.AdjacentCaves {
		adjCave := caves[cName]
		// see if we can go this way
		adjCaveVisits := 0
		if _, ok := visits[adjCave.Name]; ok {
			adjCaveVisits = visits[adjCave.Name]
		}
		if !adjCave.Small || (adjCave.Small && adjCaveVisits == 0) || (adjCave.Small && adjCaveVisits == 1 && !smallCaveVisitedTwice && adjCave.Name != "start") {
			// we can go this way
			newMap := map[string]int{}
			for k, v := range visits {
				newMap[k] = v
			}
			if adjCave.Small && adjCaveVisits == 1 && !smallCaveVisitedTwice {
				pathCount += recursiveHelper2(adjCave.Name, caves, newMap, true)
			} else {
				pathCount += recursiveHelper2(adjCave.Name, caves, newMap, smallCaveVisitedTwice)
			}
		}
	}
	return pathCount
}

func partTwo() {
	var (
		// map of cave name -> cave struct (size, times visited, adjacent cave names)
		caves = map[string]*Cave{}
	)
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	for _, line := range lines {
		split := strings.Split(line, "-")
		if _, ok := caves[split[0]]; ok {
			caves[split[0]].AdjacentCaves = append(caves[split[0]].AdjacentCaves, split[1])
		} else {
			caves[split[0]] = &Cave{
				Name:          split[0],
				Small:         !IsUpper(split[0]),
				AdjacentCaves: []string{split[1]},
			}
		}

		if _, ok := caves[split[1]]; ok {
			caves[split[1]].AdjacentCaves = append(caves[split[1]].AdjacentCaves, split[0])
		} else {
			caves[split[1]] = &Cave{
				Name:          split[1],
				Small:         !IsUpper(split[1]),
				AdjacentCaves: []string{split[0]},
			}
		}
	}

	// start at start
	// loop over adjacent caves
	numPaths := recursiveHelper2("start", caves, map[string]int{}, false)
	log.Println("Part 2 - number of valid paths now that small caves can be visited twice:", numPaths)
}

func main() {
	partOne()
	partTwo() // 25544 is too small
}
