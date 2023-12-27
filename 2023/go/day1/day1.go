package day1

import (
	"fmt"
	"math"
	"strconv"
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

func getNumberWithWords(input string) int {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var arr [9]int

	for i, word := range words {
		indWord := strings.Index(input, word)
		if indWord == -1 {
			indWord = math.MaxInt64
		}
		indDigit := strings.Index(input, strconv.Itoa(i+1))
		if indDigit == -1 {
			indDigit = math.MaxInt64
		}

		arr[i] = indWord
		if indDigit < indWord {
			arr[i] = indDigit
		}
	}

	mini := 0
	minv := math.MaxInt
	for i, v := range arr {
		if v < minv {
			mini = i + 1
			minv = v
		}
	}

	var arr2 [9]int

	for i, word := range words {
		indWord := strings.LastIndex(input, word)
		indDigit := strings.LastIndex(input, strconv.Itoa(i+1))
		arr2[i] = int(math.Max(float64(indWord), float64(indDigit)))
	}

	maxi := 0
	maxv := 0
	for i, v := range arr2 {
		if v > maxv {
			maxi = i + 1
			maxv = v
		}
	}
	num := mini*10 + maxi
	fmt.Println(num)
	return num
}

// challenge description https://adventofcode.com/2023/day/1
func Day1(raw_input string) {
	// Raw Input for this challenge is a strings separated by new lines.

	strs := strings.Split(raw_input, "\n")
	sum1, sum2 := 0, 0

	for _, v := range strs {
		num := getNumber(v)
		sum1 += num
		num2 := getNumberWithWords(v)
		sum2 += num2

		// if i == 10 {
		// 	break
		// }
	}

	fmt.Println("Solution 1:")
	fmt.Println(sum1)
	fmt.Println(sum2)
}
