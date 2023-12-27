package day2

import (
	"fmt"
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

func Main(raw_input string) {
	strs := strings.Split(raw_input, "\n")
	sum1, sum2 := 0, 0

	for _, v := range strs {
		game := parseGameStr(v)
		fmt.Println(game)
		break
	}
	fmt.Println("Solution:")
	fmt.Println(sum1)
	fmt.Println(sum2)
}
