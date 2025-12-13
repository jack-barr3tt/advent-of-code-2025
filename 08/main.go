package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	graphstuff "github.com/jack-barr3tt/gostuff/graphs"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
	union_find "github.com/jack-barr3tt/gostuff/union_find"
)

type Point struct {
	X int
	Y int
	Z int
}

type PointPair struct {
	Point1   Point
	Point2   Point
	Distance float64
}

func distance(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	dz := p1.Z - p2.Z
	return float64(dx*dx + dy*dy + dz*dz)
}

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	points := []Point{}
	for _, line := range lines {
		nums := stringstuff.GetNums(line)
		points = append(points, Point{X: nums[0], Y: nums[1], Z: nums[2]})
	}

	g := graphstuff.NewEmptyGraph()
	for _, p := range points {
		node := fmt.Sprintf("%d,%d,%d", p.X, p.Y, p.Z)
		g.AddNode(node, []graphstuff.Edge{})
	}

	allPairs := []PointPair{}
	for i, p1 := range points {
		for j, p2 := range points {
			if i < j {
				allPairs = append(allPairs, PointPair{
					Point1:   p1,
					Point2:   p2,
					Distance: distance(p1, p2),
				})
			}
		}
	}

	sort.Slice(allPairs, func(i, j int) bool {
		return allPairs[i].Distance < allPairs[j].Distance
	})

	uf := union_find.New[string]()
	for _, p := range points {
		node := fmt.Sprintf("%d,%d,%d", p.X, p.Y, p.Z)
		uf.MakeSet(node)
	}

	n := 1000
	var part1, part2 int
	for i := 0; i < len(allPairs); i++ {
		pair := allPairs[i]
		node1 := fmt.Sprintf("%d,%d,%d", pair.Point1.X, pair.Point1.Y, pair.Point1.Z)
		node2 := fmt.Sprintf("%d,%d,%d", pair.Point2.X, pair.Point2.Y, pair.Point2.Z)

		if !uf.Connected(node1, node2) {
			subgraphs := g.Subgraphs()
			if len(subgraphs) == 2 {
				part2 = pair.Point1.X * pair.Point2.X
			}

			g.AddEdge(node1, node2, int(pair.Distance))
			g.AddEdge(node2, node1, int(pair.Distance))
			uf.Union(node1, node2)
		}

		if i == n-1 {
			subgraphSizes := slicestuff.Map(func(s string) int {
				return len(g.Connected(s))
			}, g.Subgraphs())

			sort.Slice(subgraphSizes, func(i, j int) bool {
				return subgraphSizes[i] > subgraphSizes[j]
			})

			if len(subgraphSizes) >= 3 {
				part1 = subgraphSizes[0] * subgraphSizes[1] * subgraphSizes[2]
			}
		}
	}

	println(part1)
	println(part2)
}
