package day3

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func getFullNumber(grid [][]rune, i int, j int) int {
	num := 0

	chr := grid[i][j]
	k := 0
	for unicode.IsDigit(chr) {
		num += int(math.Pow10(k)) * int(chr-'0')
		grid[i][j] = '.'
		k += 1
		if j-k < 0 {
			break
		}
		chr = grid[i][j-k]
	}
	return num
}

func Main(raw_input string) {
	raw_input = strings.TrimSpace(raw_input)
	grid := strings.Split(raw_input, "\n")
	var runeGrid [][]rune

	for _, str := range grid {
		runeGrid = append(runeGrid, []rune(str))
	}

	for i, str := range runeGrid {
		for j := range str {
			getFullNumber(runeGrid, i, j)
		}
	}
	fmt.Println(grid)
}
