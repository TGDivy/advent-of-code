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

func Main(raw_input string) {
	dir_map := parseInput(raw_input)
	fmt.Println(dir_map)
	fmt.Println("Solution:")
	fmt.Println("S1: ")
	fmt.Println("S2: ")
}
