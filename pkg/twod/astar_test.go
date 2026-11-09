package twod

import (
	"strings"
	"testing"
)

func TestShortestPathOnOpenMapWithManyRoutes(t *testing.T) {
	const size = 60
	m := NewMapFromInput(strings.TrimRight(strings.Repeat(strings.Repeat(".", size)+"\n", size), "\n"))
	before := len(m)
	start, end := NewVector(0, 0), NewVector(size-1, size-1)

	path, distance, found := NewPoint(start, RIGHT).ShortestPathToPos(end, m)

	if !found {
		t.Fatal("no path found")
	}
	// One extra hop reaches the target marker, as before.
	if want := float64(2*(size-1) + 1); distance != want {
		t.Errorf("distance = %v, want %v", distance, want)
	}
	if len(m) != before {
		t.Errorf("search changed the map: %d cells left", len(m))
	}
	if len(path) != 2*(size-1)+2 || path[0].Pos != end || path[1].Pos != end || path[len(path)-1].Pos != start {
		t.Fatalf("unexpected path of %d points", len(path))
	}
	for i := len(path) - 2; i > 0; i-- {
		if got := path[i].Pos.DistFrom(path[i+1].Pos); got != 1 {
			t.Errorf("hop %d has length %d", i, got)
		}
		if path[i].Pos-path[i].Speed != path[i+1].Pos {
			t.Errorf("point %d has speed %v", i, path[i].Speed)
		}
		if want := len(path) - 1 - i; path[i].Steps != want {
			t.Errorf("point %d has %d steps, want %d", i, path[i].Steps, want)
		}
	}
}

func TestShortestPathGoesAroundWalls(t *testing.T) {
	m := NewMapFromInput("S#E\n.#.\n...")
	free := m.FilterOut('#')
	_, distance, found := NewPoint(m.Find('S')[0], RIGHT).ShortestPathToPos(m.Find('E')[0], free)
	if !found || distance != 7 {
		t.Errorf("distance = %v, found = %v", distance, found)
	}
}

func TestShortestPathReportsNoRoute(t *testing.T) {
	m := NewMapFromInput("S#E")
	free := m.FilterOut('#')
	if _, _, found := NewPoint(m.Find('S')[0], RIGHT).ShortestPathToPos(m.Find('E')[0], free); found {
		t.Error("found a path through a wall")
	}
}

func TestLongestPosExploresEachBranch(t *testing.T) {
	m := NewMapFromInput("...\n#.#")
	p := NewPoint(NewVector(1, 0), UP)
	_, dist := p.LongestPos(m.FilterOut('#'))
	// From the bottom cell: up, then one side. Only a per-branch history
	// allows both sides to be explored.
	if dist != 2 {
		t.Errorf("dist = %d, want 2", dist)
	}
}
