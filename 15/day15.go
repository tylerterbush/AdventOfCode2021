package main

import (
	"AdventOfCode2021/common"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// Start at top left, want to get to bottom right.
// Can't move diagonally
// Number at each position is risk level
// Add up the risk of each space that you enter
// Find the path with lowest total risk

// Djikstra's - start at start node with distance 0
// Initially assign all unvisited nodes a distance of infinity
// For the current node, consider all of its unvisited neighbors and calculate their tentative distances through the current node
// When we are done considering all of the unvisited neighbors of the current node, mark the current node as visited and remove it from the unvisited set
type Point struct {
	Row      int
	Col      int
	Distance int
}

func partOne() {
	var (
		seenVertices = map[string]struct{}{
			"0 0": struct{}{},
		} // start seen list as empty except for the first index
		grid      = [100][100]float64{}
		distances = [100][100]float64{}
	)

	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	row := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		for col, char := range line {
			num, _ := strconv.Atoi(string(char))
			numFloat := float64(num)
			grid[row][col] = numFloat
			distances[row][col] = math.Inf(1)
		}
		row++
	}

	// dijkstraHelper(grid, distances, seenVertices, 0, 0, 0)

	// Try doing dijstra non-recursively
	var (
		curRow      = 0
		curCol      = 0
		curDistance = float64(0)
	)
	for {
		// Base case, if THIS is the end node, return
		if curRow == 99 && curCol == 99 {
			break
		}

		// Add this to seen
		seenVertices[fmt.Sprintf("%d %d", curRow, curCol)] = struct{}{}

		// Get valid neighbors
		neighbors := []Point{}
		// Up
		if _, ok := seenVertices[fmt.Sprintf("%d %d", curRow-1, curCol)]; !ok && curRow > 0 {
			neighbors = append(neighbors, Point{Row: curRow - 1, Col: curCol})
		}
		// Down
		if _, ok := seenVertices[fmt.Sprintf("%d %d", curRow+1, curCol)]; !ok && curRow < 99 {
			neighbors = append(neighbors, Point{Row: curRow + 1, Col: curCol})
		}
		// Left
		if _, ok := seenVertices[fmt.Sprintf("%d %d", curRow, curCol-1)]; !ok && curCol > 0 {
			neighbors = append(neighbors, Point{Row: curRow, Col: curCol - 1})
		}
		// Right
		if _, ok := seenVertices[fmt.Sprintf("%d %d", curRow, curCol+1)]; !ok && curCol < 99 {
			neighbors = append(neighbors, Point{Row: curRow, Col: curCol + 1})
		}

		// Set the distances of any valid neighbors if they're smaller than current value
		for _, neighbor := range neighbors {
			prevDistance := distances[neighbor.Row][neighbor.Col]
			testDistance := curDistance + grid[neighbor.Row][neighbor.Col]
			if testDistance < prevDistance {
				distances[neighbor.Row][neighbor.Col] = testDistance
			}
		}

		// Now find the Node with the smallest distance value and call Dijkstra's from that node
		var (
			nextRow int
			nextCol int
		)
		smallestDist := float64(-1)
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				if _, ok := seenVertices[fmt.Sprintf("%d %d", i, j)]; !ok {
					if smallestDist == float64(-1) || distances[i][j] < smallestDist {
						smallestDist = distances[i][j]
						nextRow = i
						nextCol = j
					}
				}
			}
		}
		curRow = nextRow
		curCol = nextCol
		curDistance = smallestDist
	}

	log.Println("Part 1 - Distance to last point:", distances[99][99])
}

func partOneRedo() {
	var (
		seenVertices = map[string]struct{}{
			"0 0": struct{}{},
		} // start seen list as empty except for the first index
		grid = [100][100]float64{}
		// distances = [100][100]float64{}
		distancesSet = map[string]float64{}
	)

	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	row := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		for col, char := range line {
			num, _ := strconv.Atoi(string(char))
			numFloat := float64(num)
			grid[row][col] = numFloat
			// distances[row][col] = math.Inf(1)
		}
		row++
	}

	// dijkstraHelper(grid, distances, seenVertices, 0, 0, 0)

	// Try doing dijstra non-recursively
	var (
		curRow      = 0
		curCol      = 0
		curDistance = float64(0)
	)
	for {
		// Base case, if THIS is the end node, return
		if curRow == 99 && curCol == 99 {
			log.Println("Part 1 REDO - answer:", curDistance)
			break
		}

		// Add this to seen
		seenVertices[fmt.Sprintf("%d %d", curRow, curCol)] = struct{}{}

		// Get valid neighbors
		neighbors := []Point{}
		// Up
		if _, ok := seenVertices[fmt.Sprintf("%d %d", curRow-1, curCol)]; !ok && curRow > 0 {
			neighbors = append(neighbors, Point{Row: curRow - 1, Col: curCol})
		}
		// Down
		if _, ok := seenVertices[fmt.Sprintf("%d %d", curRow+1, curCol)]; !ok && curRow < 99 {
			neighbors = append(neighbors, Point{Row: curRow + 1, Col: curCol})
		}
		// Left
		if _, ok := seenVertices[fmt.Sprintf("%d %d", curRow, curCol-1)]; !ok && curCol > 0 {
			neighbors = append(neighbors, Point{Row: curRow, Col: curCol - 1})
		}
		// Right
		if _, ok := seenVertices[fmt.Sprintf("%d %d", curRow, curCol+1)]; !ok && curCol < 99 {
			neighbors = append(neighbors, Point{Row: curRow, Col: curCol + 1})
		}

		// Set the distances of any valid neighbors if they're smaller than current value
		for _, neighbor := range neighbors {
			// prevDistance := distances[neighbor.Row][neighbor.Col]
			// testDistance := curDistance + grid[neighbor.Row][neighbor.Col]
			// if testDistance < prevDistance {
			// 	distances[neighbor.Row][neighbor.Col] = testDistance
			// }
			testDistance := curDistance + grid[neighbor.Row][neighbor.Col]
			if prevDist, ok := distancesSet[fmt.Sprintf("%d %d", neighbor.Row, neighbor.Col)]; ok {
				if testDistance < prevDist {
					distancesSet[fmt.Sprintf("%d %d", neighbor.Row, neighbor.Col)] = testDistance
				}
			} else {
				distancesSet[fmt.Sprintf("%d %d", neighbor.Row, neighbor.Col)] = testDistance
			}
		}

		// Now find the Node with the smallest distance value and call Dijkstra's from that node
		var (
			nextRow int
			nextCol int
		)
		smallestDist := float64(-1)
		// for i := 0; i < 100; i++ {
		// 	for j := 0; j < 100; j++ {
		// 		if _, ok := seenVertices[fmt.Sprintf("%d %d", i, j)]; !ok {
		// 			if smallestDist == float64(-1) || distances[i][j] < smallestDist {
		// 				smallestDist = distances[i][j]
		// 				nextRow = i
		// 				nextCol = j
		// 			}
		// 		}
		// 	}
		// }
		for key, val := range distancesSet {
			keySplit := strings.Split(key, " ")
			rowTest, _ := strconv.Atoi(keySplit[0])
			colTest, _ := strconv.Atoi(keySplit[1])
			if _, ok := seenVertices[fmt.Sprintf("%d %d", rowTest, colTest)]; !ok {
				if smallestDist == float64(-1) || val < smallestDist {
					smallestDist = val
					nextRow = rowTest
					nextCol = colTest
				}
			}
		}

		curRow = nextRow
		curCol = nextCol
		curDistance = smallestDist
	}

	// log.Println("Part 1 Redo - Distance to last point:", distances[99][99])
}

// Assumptions:
// - Will only ever go right or down
// - Wont ever need to start at an infinite distance, only look at neighbors
//   where we've set the distance

func partTwo() {
	var (
		seenVertices = map[string]struct{}{
			"0 0": struct{}{},
		} // start seen list as empty except for the first index
		grid         = [500][500]float64{}
		distancesSet = map[string]float64{}
	)

	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	row := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		for col, char := range line {
			// Set the num in the normal grid
			num, _ := strconv.Atoi(string(char))
			numFloat := float64(num)
			grid[row][col] = numFloat

			// Set all the nums in the displaced grids
			for c := 0; c <= 4; c++ {
				for d := 0; d <= 4; d++ {
					if c == 0 && d == 0 {
						continue
					}
					newNum := numFloat
					for x := 0; x < c+d; x++ {
						newNum += 1
						if newNum == 10 {
							newNum = 1
						}
					}
					grid[row+(c*100)][col+(d*100)] = float64(newNum)
				}
			}
		}
		row++
	}

	var (
		curRow      = 0
		curCol      = 0
		curDistance = float64(0)
		numDone     = 0
	)
	for {
		// Base case, if THIS is the end node, return
		if curRow == 499 && curCol == 499 {
			log.Println("Part 2 - Distance to last point:", curDistance)
			break
		}

		// Add this to seen
		seenVertices[fmt.Sprintf("%d %d", curRow, curCol)] = struct{}{}
		delete(distancesSet, fmt.Sprintf("%d %d", curRow, curCol))

		numDone++
		if numDone%1000 == 0 {
			fmt.Printf("Finished %d nodes\n", numDone)
		}

		// Get valid neighbors
		neighbors := []Point{}
		// Up
		if _, ok := seenVertices[fmt.Sprintf("%d %d", curRow-1, curCol)]; !ok && curRow > 0 {
			neighbors = append(neighbors, Point{Row: curRow - 1, Col: curCol})
		}
		// Down
		if _, ok := seenVertices[fmt.Sprintf("%d %d", curRow+1, curCol)]; !ok && curRow < 499 {
			neighbors = append(neighbors, Point{Row: curRow + 1, Col: curCol})
		}
		// Left
		if _, ok := seenVertices[fmt.Sprintf("%d %d", curRow, curCol-1)]; !ok && curCol > 0 {
			neighbors = append(neighbors, Point{Row: curRow, Col: curCol - 1})
		}
		// Right
		if _, ok := seenVertices[fmt.Sprintf("%d %d", curRow, curCol+1)]; !ok && curCol < 499 {
			neighbors = append(neighbors, Point{Row: curRow, Col: curCol + 1})
		}

		// Set the distances of any valid neighbors if they're smaller than current value
		// If the prev distance doesn't exist at all, set it in the map
		for _, neighbor := range neighbors {
			testDistance := curDistance + grid[neighbor.Row][neighbor.Col]
			if prevDist, ok := distancesSet[fmt.Sprintf("%d %d", neighbor.Row, neighbor.Col)]; ok {
				if testDistance < prevDist {
					distancesSet[fmt.Sprintf("%d %d", neighbor.Row, neighbor.Col)] = testDistance
				}
			} else {
				distancesSet[fmt.Sprintf("%d %d", neighbor.Row, neighbor.Col)] = testDistance
			}
		}

		// Now find the Node with the smallest distance value and call Dijkstra's from that node
		var (
			nextRow int
			nextCol int
		)
		smallestDist := float64(-1)
		for key, val := range distancesSet {
			keySplit := strings.Split(key, " ")
			rowTest, _ := strconv.Atoi(keySplit[0])
			colTest, _ := strconv.Atoi(keySplit[1])
			if _, ok := seenVertices[fmt.Sprintf("%d %d", rowTest, colTest)]; !ok {
				if smallestDist == float64(-1) || val < smallestDist {
					smallestDist = val
					nextRow = rowTest
					nextCol = colTest
				}
			}
		}
		curRow = nextRow
		curCol = nextCol
		curDistance = smallestDist
	}
}

func main() {
	partOne()
	partOneRedo()
	partTwo()
}
