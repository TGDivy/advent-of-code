package day8

import (
	"fmt"
	"strings"
)

type NetworkNode struct {
	Value string
	Left  string
	Right string
}

type Map struct {
	Network    map[string]NetworkNode
	Directions string
}

func parseInput(raw_input string) Map {
	var dir_map Map

	parts := strings.Split(raw_input, "\n\n")
	dir_map.Directions = strings.TrimSpace(parts[0])

	network := make(map[string]NetworkNode)
	for _, line := range strings.Split(parts[1], "\n") {
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		lineParts := strings.Split(line, "=")
		head := strings.TrimSpace(lineParts[0])
		nodes := strings.Split(lineParts[1], ",")
		var networkNode NetworkNode
		networkNode.Value = head
		networkNode.Left = nodes[0]
		networkNode.Right = nodes[1]

		network[head] = networkNode
	}
	dir_map.Network = network
	return dir_map
}

func FindEnd(dir_map Map) int {
	steps := 0

	current_node, end_node := dir_map.Network["AAA"], "ZZZ"
	directionRunes := []rune(dir_map.Directions)
	for current_node.Value != end_node {
		direction := directionRunes[steps%len(directionRunes)]

		if direction == rune('R') {
			current_node = dir_map.Network[current_node.Right]
		} else {
			current_node = dir_map.Network[current_node.Left]
		}

		steps += 1
	}

	return steps
}

func Main(raw_input string) {
	dir_map := parseInput(raw_input)

	s1 := FindEnd(dir_map)

	fmt.Println("Solution:")
	fmt.Println("S1: ", s1)
	fmt.Println("S2: ")
}
