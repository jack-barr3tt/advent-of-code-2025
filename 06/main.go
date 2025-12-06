package main

import (
	"os"
	"regexp"
	"strings"

	"github.com/jack-barr3tt/gostuff/chars"
	"github.com/jack-barr3tt/gostuff/parsing"
	stringstuff "github.com/jack-barr3tt/gostuff/strings"
)

var numsRegex = regexp.MustCompile(`((\d+)\s+)+`)
var opRegex = regexp.MustCompile(`[\+\*]`)
var toks = []regexp.Regexp{*opRegex}

func performOperation(a, b int, op rune) int {
	switch op {
	case '+':
		return a + b
	case '*':
		return a * b
	default:
		panic("unknown operation")
	}
}

func main() {
	data, _ := os.ReadFile("input.txt")

	nums := [][]int{}
	ops := []rune{}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if numsRegex.MatchString(line) {
			nums = append(nums, stringstuff.GetNums(line))
		} else {
			for line, op := parsing.NextToken(line, toks); op != ""; line, op = parsing.NextToken(line, toks) {
				ops = append(ops, []rune(op)[0])
			}
		}
	}

	part1 := 0

	for c := 0; c < len(nums[0]); c++ {
		res := nums[0][c]
		for r := 1; r < len(nums); r++ {
			res = performOperation(res, nums[r][c], ops[c])
		}
		part1 += res
	}

	println(part1)

	part2 := 0

	opNo := len(ops) - 1
	tempNums := []int{}
	for c := len(lines[0]) - 1; c >= 0; c-- {
		num := 0
		for r := 0; r < len(lines); r++ {
			if chars.CharIsDigit(rune(lines[r][c])) {
				num = num*10 + int(lines[r][c]-'0')
			}
		}

		if num != 0 || c == 0 {
			tempNums = append(tempNums, num)
		}

		if num == 0 || c == 0 {
			res := tempNums[0]
			for i := 1; i < len(tempNums); i++ {
				res = performOperation(res, tempNums[i], ops[opNo])
			}
			part2 += res

			opNo--
			tempNums = []int{}
		}
	}

	println(part2)
}
