package twod

import (
	"slices"
	"strings"
	"testing"
)

const sample = "abc\ndef"

func TestGridMatchesMapCoordinates(t *testing.T) {
	g, m := NewGridFromInput(sample), NewMapFromInput(sample)
	if g.W != 3 || g.H != 2 {
		t.Fatalf("size = %dx%d", g.W, g.H)
	}
	for pos, want := range m {
		if got := rune(g.Get(pos)); got != want {
			t.Errorf("%v: got %c, want %c", pos, got, want)
		}
	}
}

func TestGridBounds(t *testing.T) {
	g := NewGridFromInput(sample)
	for _, pos := range []Vector{NewVector(-1, 0), NewVector(3, 0), NewVector(0, -1), NewVector(0, 2)} {
		if g.InBounds(pos) {
			t.Errorf("%v reported in bounds", pos)
		}
		if got := g.Get(pos); got != 0 {
			t.Errorf("Get(%v) = %c", pos, got)
		}
	}
	defer func() {
		if recover() == nil {
			t.Error("Set outside the grid did not panic")
		}
	}()
	g.Set(NewVector(3, 0), 'x') // would otherwise write to row 1
}

func TestGridSetAndFind(t *testing.T) {
	g := NewGrid(2, 2, '.')
	g.Set(NewVector(1, 0), '#')
	g.Set(NewVector(0, 1), '@')
	want := []Vector{NewVector(1, 0), NewVector(0, 1)}
	if got := g.Find('#', '@'); !slices.Equal(got, want) {
		t.Errorf("Find = %v", got)
	}
}

func TestGridRejectsRaggedInput(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("expected a panic")
		}
	}()
	NewGridFromInput("abc\nde")
}

func TestMapSizeOfNonSquareMap(t *testing.T) {
	m := NewMapFromInput(sample)
	if m.Width() != 3 || m.Height() != 2 {
		t.Errorf("size = %dx%d", m.Width(), m.Height())
	}
	if w, h := (Map{}).Width(), (Map{}).Height(); w != 1 || h != 1 {
		t.Errorf("empty size = %dx%d", w, h)
	}
}

func TestFindIsOrderedByDistance(t *testing.T) {
	m := NewMapFromInput("x.x\n...\nx.x")
	want := m.Filter('x').SortByDist(0)
	if got := m.Find('x'); !slices.Equal(got, want) {
		t.Errorf("Find = %v, want %v", got, want)
	}
}

func TestDjikstraPath(t *testing.T) {
	m := NewMapFromInput("...\n###\nS.E")
	score, path := m.Djikstra(NewVector(0, 0), NewVector(2, 0), func(a, b Vector) int { return 1 })
	if score != 2 {
		t.Errorf("score = %d", score)
	}
	want := Path{NewVector(0, 0): 0, NewVector(1, 0): 1, NewVector(2, 0): 2}
	if len(path) != len(want) {
		t.Fatalf("path = %v", path)
	}
	for pos, s := range want {
		if path[pos] != s {
			t.Errorf("%v: score %d, want %d", pos, path[pos], s)
		}
	}
}

func benchInput() string {
	return strings.TrimRight(strings.Repeat(strings.Repeat(".", 140)+"\n", 140), "\n")
}

func BenchmarkMapScan(b *testing.B) {
	m := NewMapFromInput(benchInput())
	for b.Loop() {
		count := 0
		for y := range 140 {
			for x := range 140 {
				if m[NewVector(x, y)] == '.' {
					count++
				}
			}
		}
	}
}

func BenchmarkGridScan(b *testing.B) {
	g := NewGridFromInput(benchInput())
	for b.Loop() {
		count := 0
		for y := range 140 {
			for x := range 140 {
				if g.Get(NewVector(x, y)) == '.' {
					count++
				}
			}
		}
	}
}
