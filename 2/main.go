package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
var part int

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		count := 0
		lines := strings.Split(input, "\n")
		for _, line := range lines {
			list := strings.Fields(line)
			nums := []int{}

			for _, item := range list {
				num, _ := strconv.Atoi(item)
				nums = append(nums, num)
			}

			if isSafe(nums) {
				count++
			}
		}

		fmt.Println("Output:", count)
	} else {
		count := 0
		lines := strings.Split(input, "\n")
		for _, line := range lines {
			list := strings.Fields(line)
			nums := []int{}

			for _, item := range list {
				num, _ := strconv.Atoi(item)
				nums = append(nums, num)
			}

			if isSafeWithTolerance(nums) {
				count++
			}
		}

		fmt.Println("Output:", count)
	}
}

func isSafe(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	dir := ""
	if nums[0] < nums[1] {
		dir = "INC"
	} else {
		dir = "DEC"
	}

	for i := 0; i < len(nums)-1; i++ {
		curr := nums[i]
		next := nums[i+1]
		diff := abs(curr - next)
		if diff > 3 || diff < 1 {
			return false
		}
		if dir == "INC" && curr > next {
			return false
		}
		if dir == "DEC" && curr < next {
			return false
		}
	}

	return true
}

func isSafeWithTolerance(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	if isSafe(nums) {
		return true
	}

	for i := 0; i < len(nums); i++ {
		cand := []int{}
		cand = append(cand, nums[:i]...)
		cand = append(cand, nums[i+1:]...)
		if isSafe(cand) {
			return true
		}
	}

	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
