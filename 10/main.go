package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/jack-barr3tt/gostuff/graphs"
	"github.com/jack-barr3tt/gostuff/nums"
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
	init := strings.Join(slicestuff.Map(func(_ int) string { return "0" }, pJoltage), ",")

	graph := graphs.NewVirtualGraph(func(n *graphs.Node) []graphs.Edge {
		edges := []graphs.Edge{}

		state := stringstuff.GetNums(n.Name)

		for _, b := range buttons {
			newState := make([]int, len(state))
			copy(newState, state)

			over := false

			for _, pos := range b {
				newState[pos] = (newState[pos] + 1)
				if newState[pos] > pJoltage[pos] {
					over = true
					break
				}
			}

			if over {
				continue
			}

			edges = append(edges, graphs.Edge{
				Node: strings.Join(slicestuff.Map(func(i int) string {
					return strconv.Itoa(i)
				}, newState), ","),
				Cost: 1,
			})
		}

		return edges
	}, init)

	_, len := graph.ShortestPath(init, joltage, func(n graphs.Node) int {
		curr := stringstuff.GetNums(n.Name)

		diff := 0
		for i := range curr {
			diff += pJoltage[i] - curr[i]
		}

		return diff
	})

	return len
}

func main() {
	data, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(data), "\n")

	part1 := 0
	part2 := 0

	slicestuff.ParallelMap(func(line string) bool {
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

		return true
	}, lines, nums.Min(10, len(lines)))

	println(part1)
	println(part2)
}
