package main

import (
	"os"
	"strings"

	rangestuff "github.com/jack-barr3tt/gostuff/ranges"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
	"github.com/jack-barr3tt/gostuff/types"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	parts := strings.Split(string(data), "\n\n")

	ranges := []types.Range{}

	for _, line := range strings.Split(parts[0], "\n") {
		ranges = append(ranges, rangestuff.ParseRange(line))
	}

	ingredients := stringstuff.GetNums(parts[1])

	part1 := 0

	for _, i := range ingredients {
		for _, r := range ranges {
			if r.Contains(i) {
				part1++
				break
			}
		}
	}

	println(part1)

	part2 := 0

	for _, r := range rangestuff.CombineRanges(ranges) {
		part2 += r.Width()
	}

	println(part2)
}
