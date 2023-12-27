package day2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Set struct {
	Blue  int
	Red   int
	Green int
}

type Game struct {
	Sets []Set
	ID   int
}

var MAX_SET = Set{
	Green: 13,
	Red:   12,
	Blue:  14,
}

func parseGameStr(raw_string string) Game {
	parts := strings.Split(raw_string, ":")

	gameID, _ := strconv.Atoi(strings.Trim(parts[0], "Game "))

	var sets []Set

	for _, setStr := range strings.Split(parts[1], ";") {
		var set Set
		for _, color := range strings.Split(setStr, ",") {
			colorParts := strings.Split(strings.TrimSpace(color), " ")
			count, _ := strconv.Atoi(colorParts[0])

			switch colorParts[1] {
			case "blue":
				set.Blue = count
			case "green":
				set.Green = count
			case "red":
				set.Red = count
			}
		}
		sets = append(sets, set)
	}

	game := Game{
		ID:   gameID,
		Sets: sets,
	}
	return game
}

func validSet(set Set) bool {
	if set.Blue > MAX_SET.Blue || set.Red > MAX_SET.Red || set.Green > MAX_SET.Green {
		return false
	}
	return true
}

func GameValidMax(game Game) int {
	for _, set := range game.Sets {
		if !validSet(set) {
			return 0
		}
	}
	return game.ID
}

func GameMinRequired(game Game) int {
	var minSet Set

	for _, set := range game.Sets {
		minSet.Green = int(math.Max(float64(minSet.Green), float64(set.Green)))
		minSet.Blue = int(math.Max(float64(minSet.Blue), float64(set.Blue)))
		minSet.Red = int(math.Max(float64(minSet.Red), float64(set.Red)))
	}

	return minSet.Green * minSet.Blue * minSet.Red
}

func Main(raw_input string) {
	strs := strings.Split(strings.TrimSpace(raw_input), "\n")
	sum1, sum2 := 0, 0

	for _, v := range strs {
		game := parseGameStr(v)
		sum1 += GameValidMax(game)
		sum2 += GameMinRequired(game)
		fmt.Println(game)
	}
	fmt.Println("Solution:")
	fmt.Println(sum1)
	fmt.Println(sum2)
}
