package day5

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type mapping struct {
	DestinationStart int
	SourceStart      int
	RangeLen         int
}

func parse(raw_input string) ([]int, map[string][]mapping) {
	parts := strings.Split(raw_input, "\n\n")

	seedsStr := strings.Split(strings.TrimSpace(strings.Split(parts[0], ":")[1]), " ")
	seeds := make([]int, len(seedsStr))
	for i, str := range seedsStr {
		seeds[i], _ = strconv.Atoi(str)
	}
	maps := make(map[string][]mapping)

	for i := 1; i < len(parts); i += 1 {
		lines := strings.Split(parts[i], "\n")
		name := strings.Split(lines[0], " ")[0]

		var mappings []mapping

		for j := 1; j < len(lines); j += 1 {
			var m mapping
			values := strings.Split(strings.TrimSpace(lines[j]), " ")

			m.DestinationStart, _ = strconv.Atoi(values[0])
			m.SourceStart, _ = strconv.Atoi(values[1])
			m.RangeLen, _ = strconv.Atoi(values[2])

			mappings = append(mappings, m)
		}

		maps[name] = mappings
	}

	return seeds, maps
}

func minCategory([]int, map[string][]mapping) int {
	minV := math.MaxInt

	return minV
}

func Main(raw_input string) {
	fmt.Println(raw_input)

	seeds, mappings := parse(raw_input)

	s1 := minCategory(seeds, mappings)

	fmt.Println("Solution: ")
	fmt.Println(s1)
}
