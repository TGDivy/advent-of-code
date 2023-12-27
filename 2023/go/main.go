package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/TGDivy/advent-of-code/2023/go/day1"
	"github.com/TGDivy/advent-of-code/2023/go/day2"
	"github.com/TGDivy/advent-of-code/2023/go/day3"
	"github.com/TGDivy/advent-of-code/2023/go/day4"
	"github.com/TGDivy/advent-of-code/2023/go/day5"
)

type (
	ChallengeFunc func(raw_input string)
	ChallengesMap map[int]ChallengeFunc
)

func buildChallenges() ChallengesMap {
	cmap := make(ChallengesMap)
	cmap[1] = day1.Day1
	cmap[2] = day2.Main
	cmap[3] = day3.Main
	cmap[4] = day4.Main
	cmap[5] = day5.Main
	return cmap
}

func getInput(day int, year int, inputs_path string, cookie string, force_fetch bool) string {
	log.Printf("Fetching for day %d, year %d", day, year)

	filename := filepath.Join(inputs_path, fmt.Sprintf("inputs/%d/day_%02d", year, day))

	if !force_fetch {
		data, err := os.ReadFile(filename)
		if err == nil {
			log.Printf("Input read from file: %s", filename)
			return string(data)
		}
	}

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
	year := flag.Int("year", 2023, "year of advent-of-code")
	day := flag.Int("day", 0, "day of advent-of-code")
	cookie := flag.String("cookie", "SET_COOKIE", "cookie for advent-of-code")
	force_fetch := flag.Bool("force_fetch", false, "force fetch from advent-of-code server")
	flag.Parse()

	inputs_path_raw, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		panic(fmt.Sprintf("Unable to find the path to inputs folder %s", err))
	}
	inputs_path := strings.TrimSpace(string(inputs_path_raw))

	function, ok := challengesMap[*day]

	if !ok {
		log.Fatal("challenge not initialized for: Day-", *day)
	}

	raw_input := getInput(*day, *year, inputs_path, *cookie, *force_fetch)
	raw_input = strings.TrimSpace(raw_input)
	function(raw_input)
}
