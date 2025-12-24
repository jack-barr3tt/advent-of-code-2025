package main

import (
	"os"
	"strings"

	graphstuff "github.com/jack-barr3tt/gostuff/graphs"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(data), "\n")

	g := graphstuff.NewEmptyGraph()

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		node := parts[0]

		g.AddNode(node, slicestuff.Map(func(n string) graphstuff.Edge {
			return graphstuff.Edge{
				Node: n,
				Cost: 1,
			}
		}, strings.Split(parts[1], " ")))
	}

	g.AddNode("out", []graphstuff.Edge{})

	part1 := len(g.AllPaths("you", "out"))

	println(part1)

	pathsSvrFft := g.CountPaths("svr", "fft")
	pathsFftDac := g.CountPaths("fft", "dac")
	pathsDacOut := g.CountPaths("dac", "out")

	pathsSvrDac := g.CountPaths("svr", "dac")
	pathsDacFft := g.CountPaths("dac", "fft")
	pathsFftOut := g.CountPaths("fft", "out")

	part2 := pathsSvrFft*pathsFftDac*pathsDacOut + pathsSvrDac*pathsDacFft*pathsFftOut

	println(part2)
}
