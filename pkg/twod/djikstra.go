package twod

import (
	"maps"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
)

type Path map[Vector]int

type Step struct {
	coords  Vector
	lastDir Vector
	score   int
	path    Path
}

type VisitedPoint struct {
	coords  Vector
	lastDir Vector
}

func (m Map) Djikstra(from, to Vector, scoreFunc func(curDir, nextDir Vector) int) (int, Path) {
	H, W := m.Height(), m.Width()
	priorityQueue := pq.NewWith(func(a, b interface{}) int {
		return a.(Step).score - b.(Step).score
	})
	priorityQueue.Enqueue(Step{from, RIGHT, 0, make(Path)})
	visited := make(map[VisitedPoint]bool)
	for !priorityQueue.Empty() {
		element, _ := priorityQueue.Dequeue()
		currentNode := element.(Step)
		if _, ok := visited[VisitedPoint{currentNode.coords, currentNode.lastDir}]; ok {
			continue
		}

		currentNode.path[currentNode.coords] = currentNode.score
		if currentNode.coords == to {
			return currentNode.score, currentNode.path
		}
		for _, n := range m.getNextStep(currentNode, visited, H, W, scoreFunc) {
			priorityQueue.Enqueue(n)
		}
		visited[VisitedPoint{currentNode.coords, currentNode.lastDir}] = true
	}

	return -1, make(Path)
}

func getAllowedDirections(direction Vector) []Vector {
	switch direction {
	case UP:
		return []Vector{UP, LEFT, RIGHT}
	case DOWN:
		return []Vector{DOWN, LEFT, RIGHT}
	case LEFT:
		return []Vector{LEFT, UP, RIGHT}
	case RIGHT:
		return []Vector{RIGHT, UP, DOWN}
	default:
		return []Vector{}
	}
}

func (m Map) getNextStep(current Step, visited map[VisitedPoint]bool, H, W int, scoreFunc func(curDir, nextDir Vector) int) []Step {
	possibleNext := make([]Step, 0)
	for _, dir := range getAllowedDirections(current.lastDir) {
		newPosition := current.coords + dir
		if newPosition.IsOutOfBounds(W, H) || m[newPosition] == '#' {
			continue
		}
		if _, ok := visited[VisitedPoint{newPosition, dir}]; ok {
			continue
		}
		possibleNext = append(possibleNext, Step{
			coords:  newPosition,
			lastDir: dir,
			score:   current.score + scoreFunc(current.lastDir, dir),
			path:    maps.Clone(current.path),
		})
	}
	return possibleNext
}

// GetUniqueTilesCount returns how many tiles lie on at least one cheapest path.
func (m Map) GetUniqueTilesCount(from, to Vector, scoreFunc func(curDir, nextDir Vector) int) int {
	H, W := m.Height(), m.Width()
	priorityQueue := pq.NewWith(func(a, b interface{}) int {
		return a.(Step).score - b.(Step).score
	})
	priorityQueue.Enqueue(Step{coords: from, lastDir: RIGHT})
	start := VisitedPoint{from, RIGHT}
	best := map[VisitedPoint]int{start: 0}
	previous := make(map[VisitedPoint][]VisitedPoint)
	noneVisited := make(map[VisitedPoint]bool)

	for !priorityQueue.Empty() {
		element, _ := priorityQueue.Dequeue()
		current := element.(Step)
		state := VisitedPoint{current.coords, current.lastDir}
		if current.score > best[state] {
			continue
		}
		for _, next := range m.getNextStep(current, noneVisited, H, W, scoreFunc) {
			nextState := VisitedPoint{next.coords, next.lastDir}
			known, seen := best[nextState]
			switch {
			case !seen || next.score < known:
				best[nextState] = next.score
				previous[nextState] = []VisitedPoint{state}
				priorityQueue.Enqueue(next)
			case next.score == known:
				previous[nextState] = append(previous[nextState], state)
			}
		}
	}

	var ends []VisitedPoint
	lowest := -1
	for state, score := range best {
		if state.coords != to {
			continue
		}
		if lowest == -1 || score < lowest {
			lowest, ends = score, nil
		}
		if score == lowest {
			ends = append(ends, state)
		}
	}

	tiles := make(map[Vector]bool)
	seen := make(map[VisitedPoint]bool)
	for len(ends) > 0 {
		state := ends[len(ends)-1]
		ends = ends[:len(ends)-1]
		if !seen[state] {
			seen[state] = true
			tiles[state.coords] = true
			ends = append(ends, previous[state]...)
		}
	}
	return len(tiles)
}
