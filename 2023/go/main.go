package main

import (
	"log"

	"github.com/TGDivy/advent-of-code/2023/go/day1"
)

type (
	ChallengeFunc func()
	ChallengesMap map[int]ChallengeFunc
)

func buildChallenges() ChallengesMap {
	cmap := make(ChallengesMap)
	cmap[1] = day1.Day1

	return cmap
}

func main() {
	// log.SetPrefix("greetings: ")
	log.SetFlags(0)

	challengesMap := buildChallenges()
	day := 1
	function, ok := challengesMap[day]

	if !ok {
		log.Fatal("challenge not initialized for: Day-", day)
	}

	function()
}
