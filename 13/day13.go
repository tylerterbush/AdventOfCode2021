package main

import (
	"AdventOfCode2021/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*

  x 0 1 2 3 4 5 6
Y 0  0 0 0 0 0 0 0
  1
  2

*/
// 0,0 is top left
// first value x increases to the right
// second value y increases downward
// For part 1 only focus on the first fold

// If we fold along X:
//  -  If the fold is AFTER the midpoint
//     - take anything to the right, get its X offset from the X fold.
//     - set its X position to the midpoint minus the offset
//  - If the fold is BEFORE the midpoint
type Point struct {
	X int
	Y int
}
type Fold struct {
	Axis string
	Num  int
}

func partOne() {
	var (
		points       = map[string]Point{}
		foldCommands = []Fold{}
		largestX     = 0
	)
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "fold along") {
			trim := strings.TrimPrefix(line, "fold along ")
			split := strings.Split(trim, "=")
			num, _ := strconv.Atoi(split[1])
			foldCommands = append(foldCommands, Fold{
				Axis: split[0],
				Num:  num,
			})
		} else {
			split := strings.Split(line, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			if x > largestX {
				largestX = x
			}
			points[fmt.Sprintf("%d %d", x, y)] = Point{X: x, Y: y}
		}
	}

	log.Println("Initial number of points:", len(points))

	firstFold := foldCommands[0]

	if firstFold.Axis == "x" {
		if firstFold.Num > largestX/2 {
			for key, val := range points {
				if val.X > firstFold.Num {
					newX := firstFold.Num - (val.X - firstFold.Num)
					newKey := fmt.Sprintf("%d %d", newX, val.Y)
					if _, ok := points[newKey]; ok {
						// new point already exists, remove old one from the map
						delete(points, key)
					} else {
						points[newKey] = Point{X: newX, Y: val.Y}
						delete(points, key) // still delete old key
					}
				}
			}
		} else {
			// do it the other way
		}
	}

	log.Println("New number of points:", len(points))
}

func partTwo() {
	var (
		points       = map[string]Point{}
		foldCommands = []Fold{}
		largestX     = 0
		largestY     = 0
	)
	lines, err := common.GetLinesFromFile("input.txt")
	common.FatalIf(err)

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "fold along") {
			trim := strings.TrimPrefix(line, "fold along ")
			split := strings.Split(trim, "=")
			num, _ := strconv.Atoi(split[1])
			foldCommands = append(foldCommands, Fold{
				Axis: split[0],
				Num:  num,
			})
		} else {
			split := strings.Split(line, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			if x > largestX {
				largestX = x
			}
			if y > largestY {
				largestY = y
			}
			points[fmt.Sprintf("%d %d", x, y)] = Point{X: x, Y: y}
		}
	}

	log.Println("Initial number of points:", len(points))

	for _, foldCommand := range foldCommands {
		if foldCommand.Axis == "x" {
			for key, val := range points {
				if val.X > foldCommand.Num {
					newX := foldCommand.Num - (val.X - foldCommand.Num)
					newKey := fmt.Sprintf("%d %d", newX, val.Y)
					if _, ok := points[newKey]; ok {
						delete(points, key)
					} else {
						points[newKey] = Point{X: newX, Y: val.Y}
						delete(points, key) // still delete old key
					}
				}
			}
			largestX = foldCommand.Num
		} else if foldCommand.Axis == "y" {
			for key, val := range points {
				if val.Y > foldCommand.Num {
					newY := foldCommand.Num - (val.Y - foldCommand.Num)
					newKey := fmt.Sprintf("%d %d", val.X, newY)
					if _, ok := points[newKey]; ok {
						// new point already exists, remove old one from the map
						delete(points, key)
					} else {
						points[newKey] = Point{X: val.X, Y: newY}
						delete(points, key) // still delete old key
					}
				}
			}
			largestY = foldCommand.Num
		}

		for i := 0; i < largestX; i++ {
			for j := 0; j < largestY; j++ {
				key := fmt.Sprintf("%d %d", i, j)
				if _, ok := points[key]; ok {
					fmt.Print("0")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
	}

	log.Println("new number of points:", len(points))
}

func main() {
	// partOne()
	partTwo()
}
