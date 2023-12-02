package day1

import (
	"fmt"
	"strings"
)

func isDigit(c rune) bool {
	return (c >= '0' && c <= '9')
}

func getNumber(str string) int {
	num := 0

	index := strings.LastIndexFunc(str, isDigit)
	if index == -1 {
		return 0
	}
	num += int(str[index]) - '0'
	for _, c := range str {
		if c >= '0' && c <= '9' {
			num += 10 * (int(c) - '0')
			break
		}
	}
	return num % 100
}

// challenge description https://adventofcode.com/2023/day/1
func Day1(raw_input string) {
	// Raw Input for this challenge is a strings separated by new lines.

	strs := strings.Split(raw_input, "\n")
	sum := 0

	for _, v := range strs {
		num := getNumber(v)
		sum += num
	}

	fmt.Println("Solution:")
	fmt.Println(sum)
}
