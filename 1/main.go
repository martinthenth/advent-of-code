package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	utilities "github.com/martinthenth/advent-of-code"
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

	if part == 1 {
		answer := part1(input)
		utilities.CopyToClipboard(fmt.Sprintf("%v", answer))
		fmt.Println("Output:", answer)
	} else {
		answer := part2(input)
		utilities.CopyToClipboard(fmt.Sprintf("%v", answer))
		fmt.Println("Output:", answer)
	}
}

func part1(input string) int {
	return 0
}

func part2(input string) int {
	return 0
}
