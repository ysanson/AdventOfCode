package main

import (
	"slices"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Lens struct {
	box         int
	label       string
	focalLength int
}

const (
	ADD    int = 1
	REMOVE int = 2
)

type Boxes [][]Lens

func computeHash(str string) int {
	hash := 0
	for _, char := range str {
		hash += int(char)
		hash *= 17
		hash %= 256
	}
	return hash
}

func computeFocusingPower(boxes Boxes) int {
	res := 0

	for boxIdx, box := range boxes {
		for lensIdx, lens := range box {
			res += (boxIdx + 1) * (lensIdx + 1) * lens.focalLength
		}
	}

	return res
}

func extractLens(order string) (Lens, int) {
	actionIdx := strings.IndexAny(order, "-=")
	action := order[actionIdx]
	label := order[:actionIdx]
	if action == '-' {
		return Lens{
			box:         computeHash(label),
			label:       label,
			focalLength: 0,
		}, REMOVE
	} else {
		return Lens{
			box:         computeHash(label),
			label:       label,
			focalLength: pkg.MustAtoi(order[actionIdx+1:]),
		}, ADD
	}
}

func lensIndex(box []Lens, lens Lens) int {
	return slices.IndexFunc(box, func(elt Lens) bool {
		return elt.label == lens.label
	})
}

func addLens(boxes Boxes, lens Lens) Boxes {
	if idx := lensIndex(boxes[lens.box], lens); idx != -1 {
		boxes[lens.box][idx] = lens
	} else {
		boxes[lens.box] = append(boxes[lens.box], lens)
	}

	return boxes
}

func removeLens(boxes Boxes, lens Lens) Boxes {
	if idx := lensIndex(boxes[lens.box], lens); idx != -1 {
		boxes[lens.box] = slices.Delete(boxes[lens.box], idx, idx+1)
	}
	return boxes
}

func run(input string) (any, any) {
	part1 := 0
	boxes := make(Boxes, 256)

	for _, order := range strings.Split(input, ",") {
		part1 += computeHash(order)
		lens, action := extractLens(order)
		if action == ADD {
			boxes = addLens(boxes, lens)
		} else if action == REMOVE {
			boxes = removeLens(boxes, lens)
		}
	}

	return part1, computeFocusingPower(boxes)
}

func main() {
	execute.Run(run, Tests, Puzzle, true)
}
