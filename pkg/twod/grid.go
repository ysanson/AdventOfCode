package twod

import (
	"fmt"
	"strings"
)

// Grid is a dense rectangular map of single-byte cells stored in one slice.
// It uses the same coordinates as Map: the bottom left of the input is the
// origin and UP increases Y. Prefer Map for sparse data, coordinates that can
// go negative, or cells holding anything other than a byte.
type Grid struct {
	W, H  int
	Cells []byte // row by row from the bottom: Cells[y*W+x]
}

// NewGrid returns a W by H grid with every cell set to fill.
func NewGrid(w, h int, fill byte) Grid {
	g := Grid{W: w, H: h, Cells: make([]byte, w*h)}
	for i := range g.Cells {
		g.Cells[i] = fill
	}
	return g
}

// NewGridFromInput parses a rectangular input. It panics if lines differ in length.
func NewGridFromInput(input string) Grid {
	lines := strings.Split(input, "\n")
	g := Grid{W: len(lines[0]), H: len(lines), Cells: make([]byte, 0, len(lines[0])*len(lines))}
	for j := range lines {
		line := lines[len(lines)-1-j]
		if len(line) != g.W {
			panic(fmt.Sprintf("twod: line %d has length %d, expected %d", len(lines)-1-j, len(line), g.W))
		}
		g.Cells = append(g.Cells, line...)
	}
	return g
}

// InBounds reports whether pos is a cell of the grid.
func (g Grid) InBounds(pos Vector) bool {
	return !pos.IsOutOfBounds(g.W, g.H)
}

// Get returns the cell at pos, or 0 outside the grid.
func (g Grid) Get(pos Vector) byte {
	x, y := pos.X(), pos.Y()
	if x < 0 || x >= g.W || y < 0 || y >= g.H {
		return 0
	}
	return g.Cells[y*g.W+x]
}

// Set writes the cell at pos. It panics outside the grid, where the index
// could otherwise land on a different row.
func (g Grid) Set(pos Vector, value byte) {
	if !g.InBounds(pos) {
		panic(fmt.Sprintf("twod: %v is outside the %dx%d grid", pos, g.W, g.H))
	}
	g.Cells[pos.Y()*g.W+pos.X()] = value
}

// Find returns the positions holding any of the values, from the bottom row up
// and left to right.
func (g Grid) Find(values ...byte) []Vector {
	var founds []Vector
	for i, cell := range g.Cells {
		for _, value := range values {
			if cell == value {
				founds = append(founds, NewVector(i%g.W, i/g.W))
				break
			}
		}
	}
	return founds
}
