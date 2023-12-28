package day5

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type RangeMapping struct {
	DestinationStart int
	SourceStart      int
	RangeLen         int
}

type CategoryMapping struct {
	from   string
	to     string
	ranges []RangeMapping
}

type Almanac struct {
	categoryMaps map[string]CategoryMapping
	seeds        []int
}

func parse(raw_input string) Almanac {
	var almanac Almanac
	parts := strings.Split(raw_input, "\n\n")

	seedsStr := strings.Split(strings.TrimSpace(strings.Split(parts[0], ":")[1]), " ")
	seeds := make([]int, len(seedsStr))
	for i, str := range seedsStr {
		seeds[i], _ = strconv.Atoi(str)
	}
	almanac.seeds = seeds

	maps := make(map[string]CategoryMapping)

	for i := 1; i < len(parts); i += 1 {
		var categoryMap CategoryMapping
		lines := strings.Split(parts[i], "\n")
		name := strings.Split(lines[0], " ")[0]
		nameParts := strings.Split(name, "-")
		categoryMap.to, categoryMap.from = nameParts[2], nameParts[0]

		var mappings []RangeMapping
		for j := 1; j < len(lines); j += 1 {
			var m RangeMapping
			values := strings.Split(strings.TrimSpace(lines[j]), " ")
			m.DestinationStart, _ = strconv.Atoi(values[0])
			m.SourceStart, _ = strconv.Atoi(values[1])
			m.RangeLen, _ = strconv.Atoi(values[2])
			mappings = append(mappings, m)
		}
		categoryMap.ranges = mappings

		maps[categoryMap.from] = categoryMap
	}

	almanac.categoryMaps = maps
	return almanac
}

func mapper(categoryMap CategoryMapping, seed int) int {
	for _, rangeMap := range categoryMap.ranges {
		if seed >= rangeMap.SourceStart && seed < (rangeMap.SourceStart+rangeMap.RangeLen) {
			return rangeMap.DestinationStart + seed - rangeMap.SourceStart
		}
	}
	return seed
}

func rec(alamanc Almanac, categoryMap CategoryMapping, end string, seed int) int {
	newValue := mapper(categoryMap, seed)
	if categoryMap.to == end {
		return newValue
	}
	fmt.Println(categoryMap.to, categoryMap.from)
	categoryMapNew, ok := alamanc.categoryMaps[categoryMap.to]
	if !ok {
		return 0
	}
	return rec(alamanc, categoryMapNew, end, newValue)
}

func minCategory(alamanc Almanac) int {
	start, end := "seed", "location"

	mappedSeeds := make([]int, len(alamanc.seeds))

	minV := math.MaxInt
	for i, seed := range alamanc.seeds {
		mappedSeeds[i] = rec(alamanc, alamanc.categoryMaps[start], end, seed)
		if mappedSeeds[i] < minV {
			minV = mappedSeeds[i]
		}
	}

	return minV
}

func Main(raw_input string) {
	almanac := parse(raw_input)
	s1 := minCategory(almanac)

	fmt.Println("Solution: ")
	fmt.Println(s1)
}
