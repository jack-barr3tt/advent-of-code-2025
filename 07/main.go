package main

import (
	"os"

	mapstuff "github.com/jack-barr3tt/gostuff/maps"
	mazestuff "github.com/jack-barr3tt/gostuff/maze"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	"github.com/jack-barr3tt/gostuff/types"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	maze := mazestuff.NewMaze(string(data))
	start := maze.LocateAll('S')[0]

	beams := map[types.Point]int{start: 1}
	completeBeams := map[types.Point]int{}
	splitPoints := map[types.Point]bool{}

	for len(mapstuff.Keys(beams)) > 0 {
		for beam := range beams {
			newPos, ok := maze.Move(beam, types.South)

			if !ok {
				completeBeams[beam] += beams[beam]
				delete(beams, beam)
				continue
			}

			if maze.At(newPos) == '^' {
				l, _ := maze.Move(newPos, types.West)
				r, _ := maze.Move(newPos, types.East)

				splitPoints[newPos] = true
				beams[l] += beams[beam]
				beams[r] += beams[beam]
			} else {
				beams[newPos] += beams[beam]
			}

			delete(beams, beam)
		}
	}

	part1 := len(splitPoints)
	part2 := slicestuff.Sum(mapstuff.Values(completeBeams))

	println(part1)
	println(part2)
}
