package main

import (
	"os"
	"slices"
	"strings"

	mazestuff "github.com/jack-barr3tt/gostuff/maze"
	"github.com/jack-barr3tt/gostuff/nums"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
	"github.com/jack-barr3tt/gostuff/types"
)

func inside(maze mazestuff.Maze[rune], p1, p2 types.Point) bool {
	found := false
	for x := nums.Min(p1[0], p2[0]); x <= nums.Max(p1[0], p2[0]); x++ {
		for y := nums.Min(p1[1], p2[1]); y <= nums.Max(p1[1], p2[1]); y++ {
			if maze.At(types.Point{x, y}) == '*' {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	return !found
}

func area(p1, p2 types.Point) int {
	dir := p1.DirectionTo(p2)
	x := nums.Abs(dir[0]) + 1
	y := nums.Abs(dir[1]) + 1

	return x * y
}

func main() {
	data, _ := os.ReadFile("test_input.txt")

	lines := strings.Split(string(data), "\n")

	points := []types.Point{}

	for _, line := range lines {
		points = append(points, types.PointFromSlice(stringstuff.GetNums(line)))
	}

	l, r := types.Point{0, 0}, types.Point{0, 0}

	for _, p1 := range points {
		for _, p2 := range points {
			if area(p1, p2) > area(l, r) {
				l = p1
				r = p2
			}
		}
	}
	part1 := area(l, r)

	println(part1)

	xs := slicestuff.Unique(slicestuff.Map(func(p types.Point) int { return p[0] }, points))
	ys := slicestuff.Unique(slicestuff.Map(func(p types.Point) int { return p[1] }, points))

	slices.Sort(xs)
	slices.Sort(ys)

	maze := mazestuff.NewBlankMaze(len(xs)*2+1, len(ys)*2+1, '.')

	cpoints := slicestuff.Map(func(p types.Point) types.Point {
		return types.Point{
			slicestuff.IndexOf(p[0], xs)*2 + 1,
			slicestuff.IndexOf(p[1], ys)*2 + 1,
		}
	}, points)

	for i, p := range cpoints {
		maze.Set(p, '#')

		ni := (i + 1) % len(cpoints)
		dir := p.DirectionTo(cpoints[ni]).Unit()

		for pos := p.UnsafeMove(dir); pos != cpoints[ni]; pos = pos.UnsafeMove(dir) {
			maze.Set(pos, 'X')
		}
	}

	maze.FloodFill(types.Point{0, 0}, '.', '*')

	l, r = types.Point{0, 0}, types.Point{0, 0}

	for i1, p1 := range points {
		cp1 := cpoints[i1]
		for i2, p2 := range points {
			cp2 := cpoints[i2]
			if area(p1, p2) > area(l, r) && inside(maze, cp1, cp2) {
				l = p1
				r = p2
			}
		}
	}
	part2 := area(l, r)

	println(part2)
}
