package main

import (
	"os"
	"strconv"
	"strings"

	numstuff "github.com/jack-barr3tt/gostuff/nums"
	slicestuff "github.com/jack-barr3tt/gostuff/slices"
)

func maxValueWithKDigits(nums []int, k int) int {
	n := len(nums)

	selected := []int{}
	start := 0

	for i := 0; i < k; i++ {
		end := n - (k - i) + 1

		maxIdx := start
		for j := start + 1; j < end; j++ {
			if nums[j] > nums[maxIdx] {
				maxIdx = j
			}
		}

		selected = append(selected, nums[maxIdx])
		start = maxIdx + 1
	}

	result := 0
	for i, digit := range selected {
		result += digit * numstuff.Pow(10, k-i-1)
	}
	return result
}

func main() {
	data, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(data), "\n")

	part1 := 0
	part2 := 0

	for _, line := range lines {
		nums := slicestuff.Map(func(el string) int {
			v, _ := strconv.Atoi(el)
			return v
		}, strings.Split(line, ""))

		maxValue1 := maxValueWithKDigits(nums, 2)
		part1 += maxValue1

		maxValue2 := maxValueWithKDigits(nums, 12)
		part2 += maxValue2
	}

	println(part1)
	println(part2)
}
