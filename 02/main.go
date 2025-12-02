package main

import (
	"os"
	"strconv"
	"strings"

	stringstuff "github.com/jack-barr3tt/gostuff/strings"
)

func repeatOnce(num int) bool {
	strNum := strconv.Itoa(num)

	if len(strNum)%2 != 0 {
		return false
	}

	mid := len(strNum) / 2
	return strNum[:mid] == strNum[mid:]
}

func repeatAny(num int) bool {
	strNum := strconv.Itoa(num)

	for width := 1; width <= len(strNum)/2; width++ {
		if len(strNum)%width != 0 {
			continue
		}

		same := true
		for start := 0; start <= len(strNum)-2*width; start += width {
			if strNum[start:start+width] != strNum[start+width:start+2*width] {
				same = false
				break
			}
		}
		if same {
			return true
		}
	}

	return false
}

func main() {
	data, _ := os.ReadFile("input.txt")

	ranges := strings.Split(string(data), ",")

	part1 := 0
	part2 := 0

	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		start := stringstuff.GetNum(bounds[0])
		end := stringstuff.GetNum(bounds[1])

		for i := start; i <= end; i++ {
			if repeatOnce(i) {
				part1 += i
			}
			if repeatAny(i) {
				part2 += i
			}
		}
	}

	println(part1)
	println(part2)
}
