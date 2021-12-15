package main

import (
	"AdventOfCode2021/common"
	"fmt"
	"log"
	"sort"
	"strconv"
)

var (
	seen = map[string]bool{}
)

// A low point is a point where its surrounding neighbors (4)
// are all higher than it
// That point's risk level is its height+1
// What is the sum of the risk levels of all low points?
func partOne() {
	var (
		grid               = [100][100]int{}
		row                = 0
		lowPointsTotalRisk = 0
	)
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)
	for _, line := range lines {
		if line == "" {
			continue
		}

		col := 0
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			grid[row][col] = num
			col++
		}
		row++
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			var (
				upIsGreater    = false
				downIsGreater  = false
				leftIsGreater  = false
				rightIsGreater = false
			)

			// up
			if i == 0 || grid[i-1][j] > grid[i][j] {
				upIsGreater = true
			}

			// down
			if i == 99 || grid[i+1][j] > grid[i][j] {
				downIsGreater = true
			}

			// left
			if j == 0 || grid[i][j-1] > grid[i][j] {
				leftIsGreater = true
			}

			// right
			if j == 99 || grid[i][j+1] > grid[i][j] {
				rightIsGreater = true
			}

			if upIsGreater && downIsGreater && leftIsGreater && rightIsGreater {
				lowPointsTotalRisk += grid[i][j] + 1
			}
		}
	}

	log.Println("Part One - Total Risk Level:", lowPointsTotalRisk)
}

type Point struct {
	Row int
	Col int
}

func partTwo() {
	var (
		grid       = [100][100]int{}
		row        = 0
		lowPoints  = []Point{}
		basinSizes = []int{}
	)
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)
	for _, line := range lines {
		if line == "" {
			continue
		}

		col := 0
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			grid[row][col] = num
			col++
		}
		row++
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			var (
				upIsGreater    = false
				downIsGreater  = false
				leftIsGreater  = false
				rightIsGreater = false
			)

			// up
			if i == 0 || grid[i-1][j] > grid[i][j] {
				upIsGreater = true
			}

			// down
			if i == 99 || grid[i+1][j] > grid[i][j] {
				downIsGreater = true
			}

			// left
			if j == 0 || grid[i][j-1] > grid[i][j] {
				leftIsGreater = true
			}

			// right
			if j == 99 || grid[i][j+1] > grid[i][j] {
				rightIsGreater = true
			}

			if upIsGreater && downIsGreater && leftIsGreater && rightIsGreater {
				lowPoints = append(lowPoints, Point{Row: i, Col: j})
			}
		}
	}

	log.Println("got low points:", lowPoints)

	// Do BFS from each low point and only expand if neighboring cell is greater than
	// current val. Make sure to not hit same cell twice. Store sizes in a list,
	// then return the top 3 biggest ones multiplied together
	for _, lowPoint := range lowPoints {
		basinSizes = append(basinSizes, recursiveBFSHelper(lowPoint, grid, -1))
	}

	sort.Ints(basinSizes)
	log.Println("Part 2 - 3 largest basin sizes multiplied together:", basinSizes[len(basinSizes)-1]*basinSizes[len(basinSizes)-2]*basinSizes[len(basinSizes)-3])
}

func recursiveBFSHelper(point Point, grid [100][100]int, lastValue int) int {
	// End if you hit a 9, or if we've seen this point already, or if the point is
	// out of bounds
	if point.Row < 0 || point.Row >= 100 || point.Col < 0 || point.Col >= 100 {
		return 0
	}

	curVal := grid[point.Row][point.Col]
	if curVal <= lastValue || curVal == 9 {
		return 0
	}

	pointKey := fmt.Sprintf("%d %d", point.Row, point.Col)
	if _, ok := seen[pointKey]; ok {
		return 0
	}

	seen[pointKey] = true
	return 1 +
		recursiveBFSHelper(Point{Row: point.Row + 1, Col: point.Col}, grid, curVal) +
		recursiveBFSHelper(Point{Row: point.Row - 1, Col: point.Col}, grid, curVal) +
		recursiveBFSHelper(Point{Row: point.Row, Col: point.Col + 1}, grid, curVal) +
		recursiveBFSHelper(Point{Row: point.Row, Col: point.Col - 1}, grid, curVal)
}

func main() {
	// partOne()
	partTwo()
}
