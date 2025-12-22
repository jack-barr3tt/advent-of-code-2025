package main

import (
	"os"
	"regexp"
	"strings"

	"github.com/jack-barr3tt/gostuff/graphs"
	lp "github.com/jack-barr3tt/gostuff/linear_programming"
	"github.com/jack-barr3tt/gostuff/parsing"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
)

var lightingRegex = regexp.MustCompile(`\[[.#]+\]`)
var buttonRegex = regexp.MustCompile(`\((\d+,)*\d+\)`)
var joltageRegex = regexp.MustCompile(`\{(\d+,)*\d+\}`)
var tokens = []regexp.Regexp{*lightingRegex, *buttonRegex, *joltageRegex}

func flip(l rune) rune {
	if l == '.' {
		return '#'
	} else {
		return '.'
	}
}

func computePart1(lighting string, buttons [][]int) int {
	init := string(slicestuff.Repeat('.', len(lighting)))

	graph := graphs.NewVirtualGraph(func(n *graphs.Node) []graphs.Edge {
		edges := []graphs.Edge{}

		for _, b := range buttons {
			newState := []rune(n.Name)
			for _, pos := range b {
				newState[pos] = flip(newState[pos])
			}
			edges = append(edges, graphs.Edge{
				Node: string(newState),
				Cost: 1,
			})
		}

		return edges
	}, init)

	_, len := graph.ShortestPath(init, lighting, func(n graphs.Node) int {
		return 1
	})
	return len
}

func computePart2(buttons [][]int, joltage string) int {
	pJoltage := stringstuff.GetNums(joltage)

	constraints := []lp.Constraint{}

	for i, j := range pJoltage {
		row := []float64{}
		for _, b := range buttons {
			if slicestuff.IndexOf(i, b) != -1 {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		constraint := lp.Constraint{
			Value:        float64(j),
			Coefficients: row,
			Type:         lp.EQ,
		}

		constraints = append(constraints, constraint)
	}

	problem := lp.Problem{
		Objective:   slicestuff.Repeat(1.0, len(buttons)),
		Constraints: constraints,
	}

	solution := problem.Solve(true, true)

	return int(solution.Value)
}

func main() {
	data, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(data), "\n")

	part1 := 0
	part2 := 0

	for _, line := range lines {
		lighting := ""
		buttons := [][]int{}
		joltage := ""

		for rest, token := parsing.NextToken(line, tokens); token != ""; rest, token = parsing.NextToken(rest, tokens) {
			if lightingRegex.MatchString(token) {
				lighting = token[1 : len(token)-1]
			} else if buttonRegex.MatchString(token) {
				buttons = append(buttons, stringstuff.GetNums(token))
			} else if joltageRegex.MatchString(token) {
				joltage = token[1 : len(token)-1]
			}
		}

		part1 += computePart1(lighting, buttons)
		part2 += computePart2(buttons, joltage)
	}

	println(part1)
	println(part2)
}
