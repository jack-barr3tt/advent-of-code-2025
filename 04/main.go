package main

import (
	"os"

	mazestuff "github.com/jack-barr3tt/gostuff/maze"
	"github.com/jack-barr3tt/gostuff/types"
)

var directions = []types.Direction{
	types.North,
	types.NorthEast,
	types.East,
	types.SouthEast,
	types.South,
	types.SouthWest,
	types.West,
	types.NorthWest,
}

func canAccess(maze mazestuff.Maze, roll types.Point) bool {
	count := 0
	for _, dir := range directions {
		if newPos, ok := maze.Move(roll, dir); ok && maze.At(newPos) == '@' {
			count++
		}
	}

	return count < 4
}

func main() {
	data, _ := os.ReadFile("input.txt")

	maze := mazestuff.NewMaze(string(data))

	rolls := maze.LocateAll('@')

	part1 := 0
	part2 := 0

	for _, roll := range rolls {
		count := 0
		for _, dir := range directions {
			if newPos, ok := maze.Move(roll, dir); ok && maze.At(newPos) == '@' {
				count++
			}
		}

		if canAccess(maze, roll) {
			part1++
		}
	}

	println(part1)

	for {
		found := false
		for _, roll := range rolls {
			if maze.At(roll) == '@' && canAccess(maze, roll) {
				maze.Set(roll, '.')
				found = true
				part2++
			}
		}
		if !found {
			break
		}
	}

	println(part2)
}
