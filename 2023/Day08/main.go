package main

import (
	"strings"
	"sync"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Node struct {
	left  string
	right string
}

func buildNetwork(input string) map[string]Node {
	nodeDefinitions := strings.Split(input, "\n")[2:]
	nodes := make(map[string]Node, len(nodeDefinitions))
	for _, nodeDefinition := range nodeDefinitions {
		split := strings.Split(nodeDefinition, " = ")
		nodes[split[0]] = Node{
			left:  split[1][1:4],
			right: split[1][6:9],
		}
	}
	return nodes
}

func traverseGraph(graph map[string]Node, instructions string, start string, goal string, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	steps := 0
	nodeName := start
	for !strings.HasSuffix(nodeName, goal) {
		if nodeName == graph[nodeName].left && nodeName == graph[nodeName].right {
			c <- -1
			return
		}
		if instructions[steps%len(instructions)] == 'L' {
			nodeName = graph[nodeName].left
		} else {
			nodeName = graph[nodeName].right
		}
		steps++
	}
	c <- steps
}

func run(input string) (any, any) {
	instructions := strings.Split(input, "\n")[0]
	network := buildNetwork(input)
	var wg sync.WaitGroup

	chan1 := make(chan int, 1)
	chan2 := make(chan int)
	wg.Add(1)
	go traverseGraph(network, instructions, "AAA", "ZZZ", chan1, &wg)
	for key := range network {
		if key[2] == 'A' {
			wg.Add(1)
			go traverseGraph(network, instructions, key, "Z", chan2, &wg)
		}
	}
	go func() {
		wg.Wait()
		close(chan1)
		close(chan2)
	}()
	part2 := <-chan2
	for i := range chan2 {
		part2 = pkg.LCM(part2, i)
	}
	return <-chan1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
