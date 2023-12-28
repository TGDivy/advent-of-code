package day8

import (
	"fmt"
	"strings"

	"github.com/schollz/progressbar/v3"
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

func hasEnded(nodes []string) bool {
	for _, node := range nodes {
		lastChar := []rune(node)[2]
		if lastChar != rune('Z') {
			return false
		}
	}
	return true
}

func FindEndForGhosts(dir_map Map) int {
	steps := 0

	current_nodes := []string{}
	for node := range dir_map.Network {
		lastChar := []rune(node)[2]
		if lastChar == rune('A') {
			current_nodes = append(current_nodes, node)
		}
	}

	bar := progressbar.Default(-1)
	directionRunes := []rune(dir_map.Directions)
	for !hasEnded(current_nodes) {
		direction := directionRunes[steps%len(directionRunes)]
		for i, name := range current_nodes {
			current_node := dir_map.Network[name]
			if direction == rune('R') {
				current_node = dir_map.Network[current_node.Right]
			} else {
				current_node = dir_map.Network[current_node.Left]
			}
			current_nodes[i] = current_node.Value
		}
		steps += 1
		bar.Add(1)
	}

	return steps
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmArray(arr []int) int {
	ans := arr[0]
	for i := 1; i < len(arr); i++ {
		ans = lcm(ans, arr[i])
	}
	return ans
}

func FindEndForGhosts2(dir_map Map, start_node string) int {
	steps := 0

	current_node := dir_map.Network[start_node]
	directionRunes := []rune(dir_map.Directions)
	for !hasEnded([]string{current_node.Value}) {
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
	// s2 := FindEndForGhosts(dir_map)

	s2steps := []int{}
	for node := range dir_map.Network {
		lastChar := []rune(node)[2]
		if lastChar == rune('A') {
			s2steps = append(s2steps, FindEndForGhosts2(dir_map, node))
		}
	}

	s2 := lcmArray(s2steps)

	fmt.Println("Solution:")
	fmt.Println("S1: ", s1)
	fmt.Println("S2: ", s2)
}
