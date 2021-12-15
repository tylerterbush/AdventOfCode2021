package main

import (
	"AdventOfCode2021/common"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

// seven segment display

//  aaaa
// b    c
// b    c
//  dddd
// e    f
// e    f
//  gggg
//
var (
	NumSegments = map[int][]rune{
		0: []rune{'a', 'b', 'c', 'e', 'f', 'g'},
		1: []rune{'c', 'f'},
		2: []rune{'a', 'c', 'd', 'e', 'g'},
		3: []rune{'a', 'c', 'd', 'f', 'g'},
		4: []rune{'b', 'c', 'd', 'f'},
		5: []rune{'a', 'b', 'd', 'f', 'g'},
		6: []rune{'a', 'b', 'd', 'e', 'f', 'g'},
		7: []rune{'a', 'c', 'f'},
		8: []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'},
		9: []rune{'a', 'b', 'c', 'd', 'f', 'g'},
	}

	NormalMapping = map[string]string{
		"abcefg":  "0",
		"cf":      "1",
		"acdeg":   "2",
		"acdfg":   "3",
		"bcdf":    "4",
		"abdfg":   "5",
		"abdefg":  "6",
		"acf":     "7",
		"abcdefg": "8",
		"abcdfg":  "9",
	}
)

type Test struct {
	DigitWiring []string
	DisplayNum  []string
}

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	var r ByRune = StringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}

// Wires have been connected randomly
// We have a four-digit display
// They have been connected randomly for each of the 4 seven-segment numbers
// Unique number digits are (1, 4, 7, 8)
// Part one: just get the number of times 1,4,7, or 8 appear in output
func partOne() {
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	tests := []Test{}
	for _, line := range lines {
		split := strings.Split(line, " | ")
		digitWiring := strings.Split(split[0], " ")
		displayNum := strings.Split(split[1], " ")
		tests = append(tests, Test{DigitWiring: digitWiring, DisplayNum: displayNum})
	}

	count1478 := 0

	for _, test := range tests {
		for _, outputDigit := range test.DisplayNum {
			switch len(outputDigit) {
			case len(NumSegments[1]), len(NumSegments[4]), len(NumSegments[7]), len(NumSegments[8]):
				count1478++
			}
		}
	}

	log.Println("Part One - number of times 1, 4, 7 or 8 appears in output:", count1478)
}

func partTwo() {
	var (
		total = 0
	)

	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	tests := []Test{}
	for _, line := range lines {
		split := strings.Split(line, " | ")
		digitWiring := strings.Split(split[0], " ")
		displayNum := strings.Split(split[1], " ")
		tests = append(tests, Test{DigitWiring: digitWiring, DisplayNum: displayNum})
	}

	// Loop over input wirings
	// get segments used for unique numbers (1,4,7,8)
	for _, test := range tests {
		// if i != 100 {
		// 	continue
		// }
		var (
			wiringByLengths            = map[int][]string{}
			randomSegmentToRealSegment = map[string]string{}
			sixWiring                  string
			zeroWiring                 string
		)
		// Get the wiring maps for this test
		for _, wiring := range test.DigitWiring {
			if _, ok := wiringByLengths[len(wiring)]; ok {
				wiringByLengths[len(wiring)] = append(wiringByLengths[len(wiring)], wiring)
			} else {
				wiringByLengths[len(wiring)] = []string{wiring}
			}
		}
		log.Println(wiringByLengths)

		// FIGURE OUT WHAT THE DIGITS ARE
		// can first first what maps to segment 'A' by comparing 1 and 7 to find the
		// one character that isn't in 1
		oneWiring := wiringByLengths[2][0]
		sevenWiring := wiringByLengths[3][0]
		for _, char := range sevenWiring {
			if !strings.Contains(oneWiring, string(char)) {
				randomSegmentToRealSegment[string(char)] = "a"
			}
		}

		// Can now figure out 'C' and 'F' by taking into account the two characters
		// that make up the number 1. Loop over digits that have length 6 to find the
		// one that DOES NOT have the two chars that make up 1. This is digit number
		// 6. The one that it DOESN'T have is 'C' and the one it does have is 'F'
		for _, w := range wiringByLengths[6] {
			if !(strings.Contains(w, string(oneWiring[0])) && strings.Contains(w, string(oneWiring[1]))) {
				sixWiring = w
				if strings.Contains(w, string(oneWiring[0])) {
					randomSegmentToRealSegment[string(oneWiring[0])] = "f"
					randomSegmentToRealSegment[string(oneWiring[1])] = "c"
				} else {
					randomSegmentToRealSegment[string(oneWiring[0])] = "c"
					randomSegmentToRealSegment[string(oneWiring[1])] = "f"
				}
			}
		}

		// Can now figure out 'B' and 'D' by looking at 4, taking out the two characters
		// for 'c' and 'f' which we can get from the 1 random mapping
		// Then we can find 0 by looping over the mappings with length 6 to find the
		// only one that does not have both the "c" and "f" random chars
		fourWiring := wiringByLengths[4][0]
		fourWiringLeftover := strings.Replace(fourWiring, string(oneWiring[0]), "", -1)
		fourWiringLeftover = strings.Replace(fourWiringLeftover, string(oneWiring[1]), "", -1)
		log.Println("Four wiringLeftover:", fourWiringLeftover)
		for _, w2 := range wiringByLengths[6] {
			if !(strings.Contains(w2, string(fourWiringLeftover[0])) && strings.Contains(w2, string(fourWiringLeftover[1]))) {
				zeroWiring = w2
				if strings.Contains(w2, string(fourWiringLeftover[0])) {
					randomSegmentToRealSegment[string(fourWiringLeftover[0])] = "b"
					randomSegmentToRealSegment[string(fourWiringLeftover[1])] = "d"
				} else {
					randomSegmentToRealSegment[string(fourWiringLeftover[0])] = "d"
					randomSegmentToRealSegment[string(fourWiringLeftover[1])] = "b"
				}
			}
		}

		// Now get the last two segments 'g' and 'e' by grabbing 9 from the list of
		// wirings with length 6. This will be the last one in the list that we haven't used yet
		// log.Println("Zero Wiring:", zeroWiring, "Six Wiring:", sixWiring)
		for _, w := range wiringByLengths[6] {
			if w != zeroWiring && w != sixWiring {
				nineWiring := w
				// found the 9 wiring. Loop over all keys of our randomSegmentToRealSegment
				// delete this char from the string. The remaining char is 'g'
				for key, _ := range randomSegmentToRealSegment {
					nineWiring = strings.Replace(nineWiring, key, "", 1)
				}

				randomSegmentToRealSegment[nineWiring] = "g"
			}
		}

		// Get the 'e' mapping now
		eightWiring := wiringByLengths[7][0]
		for key, _ := range randomSegmentToRealSegment {
			eightWiring = strings.Replace(eightWiring, key, "", 1)
		}
		randomSegmentToRealSegment[eightWiring] = "e"

		log.Println("Final mapping for this test:", randomSegmentToRealSegment)

		// Now get the digit and add it to the total
		digits := ""
		for _, num := range test.DisplayNum {
			log.Println(num)
			actualString := ""
			for _, char := range num {
				actualString = fmt.Sprintf("%s%s", actualString, randomSegmentToRealSegment[string(char)])
			}

			// Now sort the string + see what number it goes to
			log.Println("actual String:", actualString)
			actualString = SortStringByCharacter(actualString)
			log.Println("actual string after sorting:", actualString)
			digits = fmt.Sprintf("%s%s", digits, NormalMapping[actualString])
		}

		number, _ := strconv.Atoi(digits)
		total += number
	}

	log.Println("Part 2 - sum of all numbers:", total)
}

func main() {
	partOne()
	partTwo()
}
