package main

import (
	"os"
	"strings"

	stringstuff "github.com/jack-barr3tt/gostuff/strings"
)

func getDirection(line string) int {
	if line[0] == 'L' {
		return -1
	}
	return 1
}

func main() {
	data, _ := os.ReadFile("input.txt")

	pos := 50
	part1 := 0
	part2 := 0

	for _, line := range strings.Split(string(data), "\n") {
		for i := 0; i < stringstuff.GetNum(line); i++ {
			pos += getDirection(line)
			if pos < 0 {
				pos = 99
			} else if pos > 99 {
				pos = 0
			}

			if pos == 0 {
				part2++
			}
		}
		if pos == 0 {
			part1++
		}
	}

	println(part1)
	println(part2)
}
