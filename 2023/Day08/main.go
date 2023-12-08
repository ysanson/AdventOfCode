package main

import "strings"

type Node struct {
	name  string
	left  string
	right string
}

func getInstructions(input string) string {
	return strings.Split(input, "\n")[0]
}

func buildNode(nodeDefinition string) Node {
	split := strings.Split(nodeDefinition, " = ")

	return Node{
		name:  split[0],
		right: split[1][1:4],
		left:  split[1][6:9],
	}
}

func buildNetwork(input string) map[string]Node {
	nodeDefinitions := strings.Split(input, "\n")[:2]
	var node Node
	nodes := make(map[string]Node, len(nodeDefinitions))
	for _, nodeDefinition := range nodeDefinitions {
		node = buildNode(nodeDefinition)
		nodes[node.name] = node
	}
	return nodes
}

func traverseGraph(graph map[string]Node, instructions string, current Node, goal string, steps int) int {
	if current.name == goal {
		return steps
	}

	return traverseGraph(graph, instructions, nil, goal, steps+1)
}
