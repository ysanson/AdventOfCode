package main

import (
	"strings"

	"github.com/ysanson/AdventOfCode/2023/pkg"
	"github.com/ysanson/AdventOfCode/2023/pkg/execute"
)

type Node struct {
	name  string
	left  string
	right string
}

func buildNode(nodeDefinition string) Node {
	split := strings.Split(nodeDefinition, " = ")
	return Node{
		name:  split[0],
		left:  split[1][1:4],
		right: split[1][6:9],
	}
}

func buildNetwork(input string) map[string]Node {
	nodeDefinitions := strings.Split(input, "\n")[2:]
	var node Node
	nodes := make(map[string]Node, len(nodeDefinitions))
	for _, nodeDefinition := range nodeDefinitions {
		node = buildNode(nodeDefinition)
		nodes[node.name] = node
	}
	return nodes
}

func traverseGraph(graph map[string]Node, instructions string, currentName string, goal string, steps int) int {
	if strings.HasSuffix(currentName, goal) {
		return steps
	}
	currentNode := graph[currentName]
	if currentNode.name == currentNode.right && currentNode.name == currentNode.left {
		return -1
	}
	currentInstruction := instructions[steps%len(instructions)]
	var nextNodeName string
	if currentInstruction == 'L' {
		nextNodeName = currentNode.left
	} else {
		nextNodeName = currentNode.right
	}

	return traverseGraph(graph, instructions, nextNodeName, goal, steps+1)
}

func run(input string) (any, any) {
	instructions := strings.Split(input, "\n")[0]
	network := buildNetwork(input)
	part1 := traverseGraph(network, instructions, "AAA", "ZZZ", 0)
	paths := make([]int, 0, 10)
	for _, node := range network {
		if node.name[2] == 'A' {
			length := traverseGraph(network, instructions, node.name, "Z", 0)
			paths = append(paths, length)
		}
	}
	part2 := pkg.LcmAll(paths[0], paths[1:]...)
	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
