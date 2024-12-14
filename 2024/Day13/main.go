package main

import (
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Machine struct {
	aX     int
	aY     int
	bX     int
	bY     int
	prizeX int
	prizeY int
}

func parseMachine(input string) Machine {
	lines := strings.Split(input, "\n")
	machine := Machine{}
	btnA := strings.Split(lines[0], "+")
	machine.aX = pkg.MustAtoi(btnA[1][:strings.Index(btnA[1], ",")])
	machine.aY = pkg.MustAtoi(btnA[2])
	btnB := strings.Split(lines[1], "+")
	machine.bX = pkg.MustAtoi(btnB[1][:strings.Index(btnB[1], ",")])
	machine.bY = pkg.MustAtoi(btnB[2])
	prize := strings.Split(lines[2], "=")
	machine.prizeX = pkg.MustAtoi(prize[1][:strings.Index(prize[1], ",")])
	machine.prizeY = pkg.MustAtoi(prize[2])
	return machine
}

func solveDoubleEquations(machine Machine) int {
	a1, b1, c1 := machine.aX, machine.bX, machine.prizeX
	a2, b2, c2 := machine.aY, machine.bY, machine.prizeY

	det := a1*b2 - a2*b1
	if det == 0 {
		return 0
	}
	detX := c1*b2 - c2*b1
	detY := a1*c2 - a2*c1

	// Calcul des solutions
	x := float64(detX) / float64(det)
	y := float64(detY) / float64(det)

	if pkg.IsInteger(x) && pkg.IsInteger(y) && x <= 100 && y <= 100 {
		return int(x*3 + y)
	}
	return 0
}

func solveCorrectedEquations(machine Machine) int {
	correction := 10000000000000
	a1, b1, c1 := machine.aX, machine.bX, machine.prizeX+correction
	a2, b2, c2 := machine.aY, machine.bY, machine.prizeY+correction

	det := a1*b2 - a2*b1
	if det == 0 {
		return 0
	}
	detX := c1*b2 - c2*b1
	detY := a1*c2 - a2*c1

	// Calcul des solutions
	x := float64(detX) / float64(det)
	y := float64(detY) / float64(det)

	if pkg.IsInteger(x) && pkg.IsInteger(y) {
		return int(x*3 + y)
	}
	return 0

}

func run(input string) (interface{}, interface{}) {
	inputs := strings.Split(input, "\n\n")
	part1, part2 := 0, 0
	for _, in := range inputs {
		machine := parseMachine(in)
		part1 += solveDoubleEquations(machine)
		part2 += solveCorrectedEquations(machine)
	}
	return part1, part2
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
