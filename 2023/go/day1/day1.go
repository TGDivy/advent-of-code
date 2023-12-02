package day1

import (
	"fmt"
	"strings"
)

// challenge description https://adventofcode.com/2023/day/1
func Day1(raw_input string) {
	// Raw Input for this challenge is a strings separated by new lines.

	strings := strings.Split(raw_input, "\n")
	fmt.Printf("Day 1 Challenge! %q", strings)
}
