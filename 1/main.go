package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	lines := strings.Split(input, "\n")
	list1, list2 := []int{}, []int{}
	for _, line := range lines {
		numbers := strings.Fields(line)
		number1, _ := strconv.Atoi(numbers[0])
		number2, _ := strconv.Atoi(numbers[1])

		list1 = append(list1, number1)
		list2 = append(list2, number2)
	}

	if part == 1 {
		fmt.Println("Output:", part1(list1, list2))
	} else {
		fmt.Println("Output:", part2(list1, list2))
	}
}

func part1(list1 []int, list2 []int) int {
	sort.Ints(list1)
	sort.Ints(list2)

	result := 0
	for i := 0; i < len(list1) && i < len(list2); i++ {
		result += abs(list2[i] - list1[i])
	}

	return result
}

func part2(list1 []int, list2 []int) int {
	counts := map[int]int{}
	for _, number := range list2 {
		counts[number]++
	}

	result := 0
	for _, number := range list1 {
		result += number * counts[number]
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
