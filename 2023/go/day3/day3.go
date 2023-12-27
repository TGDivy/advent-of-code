package day3

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func getFullNumber(grid [][]rune, i int, j int) int {
	num := 0

	if !unicode.IsDigit(grid[i][j]) {
		return 0
	}

	for j+1 < len(grid[i]) && unicode.IsDigit(grid[i][j+1]) {
		j += 1
	}

	chr := grid[i][j]
	k := 0
	for unicode.IsDigit(chr) {
		num += int(math.Pow10(k)) * int(chr-'0')
		grid[i][j-k] = '.'
		k += 1
		if j-k < 0 {
			break
		}
		chr = grid[i][j-k]
	}
	return num
}

type Direction struct {
	x int
	y int
}

var dirs = []Direction{
	{x: 0, y: 1},
	{x: 0, y: -1},

	{x: 1, y: 1},
	{x: 1, y: -1},
	{x: 1, y: 0},

	{x: -1, y: 1},
	{x: -1, y: -1},
	{x: -1, y: 0},
}

func Main(raw_input string) {
	raw_input = strings.TrimSpace(raw_input)
	grid := strings.Split(raw_input, "\n")
	// var runeGrid [][]rune
	//
	// for _, str := range grid {
	// 	runeGrid = append(runeGrid, []rune(str))
	// }
	//
	// sum1 := 0
	// for i, str := range runeGrid {
	// 	for j, chr := range str {
	// 		if !unicode.IsDigit(chr) && chr != '.' {
	// 			for _, dir := range dirs {
	// 				sum1 += getFullNumber(runeGrid, i+dir.x, j+dir.y)
	// 			}
	// 		}
	// 	}
	// }
	//
	// fmt.Println("Solution:")
	// fmt.Println("1: ", sum1)

	var runeGrid2 [][]rune

	for _, str := range grid {
		runeGrid2 = append(runeGrid2, []rune(str))
	}

	sum2 := 0
	for i, str := range runeGrid2 {
		for j, chr := range str {
			if chr == '*' {
				gearRatio := 1
				count := 0
				for _, dir := range dirs {
					num := getFullNumber(runeGrid2, i+dir.x, j+dir.y)
					if num > 0 {
						count += 1
						gearRatio *= num
					}
					if count > 2 {
						gearRatio = 0
						break
					}
				}
				if count == 2 {
					sum2 += gearRatio
				}
			}
		}
	}

	fmt.Println("2: ", sum2)
}
