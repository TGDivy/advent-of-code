package day6

import (
	"fmt"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func filter(ss []string, test func(string) bool) []string {
	var ret []string
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

func parseData(raw_input string) []Race {
	var races []Race

	lines := strings.Split(raw_input, "\n")

	times := strings.Split(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " ")
	distances := strings.Split(strings.TrimSpace(strings.Split(lines[1], ":")[1]), " ")

	times = filter(times, func(s string) bool { return s != "" })
	distances = filter(distances, func(s string) bool { return s != "" })

	for i := 0; i < len(times); i++ {
		var race Race
		race.Time, _ = strconv.Atoi(times[i])
		race.Distance, _ = strconv.Atoi(distances[i])
		races = append(races, race)
	}

	return races
}

func parseDateFixedKerning(raw_input string) Race {
	var race Race
	lines := strings.Split(raw_input, "\n")

	times := strings.Split(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " ")
	distances := strings.Split(strings.TrimSpace(strings.Split(lines[1], ":")[1]), " ")

	times = filter(times, func(s string) bool { return s != "" })
	distances = filter(distances, func(s string) bool { return s != "" })

	time := strings.Join(times, "")
	distance := strings.Join(distances, "")

	race.Time, _ = strconv.Atoi(time)
	race.Distance, _ = strconv.Atoi(distance)

	return race
}

func possibleWaysToWin(race Race) int {
	count := 0

	for i := 1; i < race.Time; i++ {
		dist := (race.Time - i) * i
		if dist > race.Distance {
			count += 1
		}
	}
	return count
}

func Main(raw_input string) {
	races := parseData(raw_input)
	m1 := 1
	for _, race := range races {
		m1 *= possibleWaysToWin(race)
	}
	fmt.Println("Solution: ")
	fmt.Println("S1: ", m1)
	correctRace := parseDateFixedKerning(raw_input)
	s2 := possibleWaysToWin(correctRace)
	fmt.Println("S2: ", s2)
}
