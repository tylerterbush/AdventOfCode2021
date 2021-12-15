package main

import (
	"AdventOfCode2021/common"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jinzhu/copier"
)

// Find the winning board (only vertical or horizontal)
// Find the sum of all unmarked numbers on that board
// Multiply that by the number that was JUST called

type Board struct {
	Grid      [5][5]int
	ValidNums map[int]RowCol
	RowCount  [5]int
	ColCount  [5]int
	Done      bool
}

type RowCol struct {
	Row int
	Col int
}

func partOne() string {
	var (
		numsToCall = []int{}
		boards     = []Board{}
	)
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	currRow := 0
	board1 := Board{Grid: [5][5]int{}, ValidNums: map[int]RowCol{}, RowCount: [5]int{}, ColCount: [5]int{}}
	for i, line := range lines {
		// get numsToCall
		if i == 0 {
			numStrings := strings.Split(line, ",")
			for _, ns := range numStrings {
				num, _ := strconv.Atoi(ns)
				numsToCall = append(numsToCall, num)
			}
			continue
		} else if i == 1 {
			continue
		}

		if line == "" {
			newBoard := Board{}
			copier.Copy(&newBoard, &board1)
			boards = append(boards, newBoard)
			board1.Grid = [5][5]int{}
			board1.ValidNums = map[int]RowCol{}
			currRow = 0
			continue
		}
		numSplit := strings.Fields(line)
		for i, ns := range numSplit {
			num, _ := strconv.Atoi(ns)
			board1.Grid[currRow][i] = num
			board1.ValidNums[num] = RowCol{Row: currRow, Col: i}
		}
		currRow++
	}

	// log.Println(boards)

	for _, num := range numsToCall {
		for i := 0; i < len(boards); i++ {
			if rowCol, ok := boards[i].ValidNums[num]; ok {
				boards[i].Grid[rowCol.Row][rowCol.Col] = 0
				boards[i].RowCount[rowCol.Row]++
				boards[i].ColCount[rowCol.Col]++
				if boards[i].RowCount[rowCol.Row] == 5 || boards[i].ColCount[rowCol.Col] == 5 {
					totalUnfilled := 0
					for x := 0; x < 5; x++ {
						for y := 0; y < 5; y++ {
							totalUnfilled += boards[i].Grid[x][y]
						}
					}

					return fmt.Sprintf("Game over. Number called: %d Total of unfinished part of board: %d, Product: %d", num, totalUnfilled, num*totalUnfilled)
				}
			}
		}
	}
	return "nayfin"
}

func partTwo() string {
	var (
		numsToCall   = []int{}
		boards       = []Board{}
		returnString = "nayfin"
	)
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	currRow := 0
	board1 := Board{Grid: [5][5]int{}, ValidNums: map[int]RowCol{}, RowCount: [5]int{}, ColCount: [5]int{}, Done: false}
	for i, line := range lines {
		// get numsToCall
		if i == 0 {
			numStrings := strings.Split(line, ",")
			for _, ns := range numStrings {
				num, _ := strconv.Atoi(ns)
				numsToCall = append(numsToCall, num)
			}
			continue
		} else if i == 1 {
			continue
		}

		if line == "" {
			newBoard := Board{}
			copier.Copy(&newBoard, &board1)
			boards = append(boards, newBoard)
			board1.Grid = [5][5]int{}
			board1.ValidNums = map[int]RowCol{}
			currRow = 0
			continue
		}
		numSplit := strings.Fields(line)
		for i, ns := range numSplit {
			num, _ := strconv.Atoi(ns)
			board1.Grid[currRow][i] = num
			board1.ValidNums[num] = RowCol{Row: currRow, Col: i}
		}
		currRow++
	}

	for _, num := range numsToCall {
		for i := 0; i < len(boards); i++ {
			if boards[i].Done {
				continue
			}
			if rowCol, ok := boards[i].ValidNums[num]; ok {
				boards[i].Grid[rowCol.Row][rowCol.Col] = 0
				boards[i].RowCount[rowCol.Row]++
				boards[i].ColCount[rowCol.Col]++
				if boards[i].RowCount[rowCol.Row] == 5 || boards[i].ColCount[rowCol.Col] == 5 {
					totalUnfilled := 0
					for x := 0; x < 5; x++ {
						for y := 0; y < 5; y++ {
							totalUnfilled += boards[i].Grid[x][y]
						}
					}

					returnString = fmt.Sprintf("Game over. Number called: %d Total of unfinished part of board: %d, Product: %d", num, totalUnfilled, num*totalUnfilled)
					boards[i].Done = true
				}
			}
		}
	}
	return returnString
}

func main() {
	log.Println(partOne())
	log.Println(partTwo())
}
