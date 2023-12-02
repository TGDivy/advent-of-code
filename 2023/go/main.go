package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/TGDivy/advent-of-code/2023/go/day1"
)

type (
	ChallengeFunc func(raw_input string)
	ChallengesMap map[int]ChallengeFunc
)

func buildChallenges() ChallengesMap {
	cmap := make(ChallengesMap)
	cmap[1] = day1.Day1

	return cmap
}

func getInput(day int, year int, inputs_path string, cookie string) string {
	log.Printf("Fetching for day %d, year %d", day, year)

	filename := filepath.Join(inputs_path, fmt.Sprintf("inputs/%d/day_%02d", year, day))

	fmt.Println(inputs_path, filename)
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	body := GetWithAOCCookie(url, cookie)
	WriteToFile(filename, body)

	return string(body)
}

func main() {
	// log.SetPrefix("greetings: ")
	log.SetFlags(0)

	challengesMap := buildChallenges()
	day := 1
	cookie := "123213"
	inputs_path_raw, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		panic(fmt.Sprintf("Unable to find the path to inputs folder %s", err))
	}
	inputs_path := strings.TrimSpace(string(inputs_path_raw))

	function, ok := challengesMap[day]

	if !ok {
		log.Fatal("challenge not initialized for: Day-", day)
	}

	raw_input := getInput(day, 2023, inputs_path, cookie)

	function(raw_input)
}
