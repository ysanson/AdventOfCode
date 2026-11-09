package twod

import (
	"github.com/beefsack/go-astar"
)

func (p *P) Clone() *P {
	return &P{
		Pos:           p.Pos,
		Speed:         p.Speed,
		Steps:         p.Steps,
		emptyPoints:   p.emptyPoints.Clone(),
		exitCondition: p.exitCondition,
		nodes:         p.nodes,
	}
}

func (p *P) LongestPos(emptyPoints Map) (maxPos Vector, maxDist int) {
	p.emptyPoints = emptyPoints
	p.nodes = nil
	return p.longestPath(0)
}

func (p *P) longestPath(currentDist int) (maxPos Vector, maxDist int) {
	maxPos = p.Pos
	maxDist = currentDist
	for _, neighbor := range p.PathNeighbors() {
		newPos, dist := neighbor.(*P).longestPath(currentDist + 1)
		if dist > maxDist {
			maxDist = dist
			maxPos = newPos
		}
	}
	return
}

func (p *P) ShortestPathToPos(pos Vector, emptyPoints Map) (path []*P, distance float64, found bool) {
	p.emptyPoints = emptyPoints
	p.nodes = map[Vector]*P{p.Pos: p}
	target := NewPoint(pos, 0)

	// exit condition is when current point pos is target position
	p.exitCondition = func(currentPoint *P, neighbors []astar.Pather) []astar.Pather {
		if currentPoint.Pos == pos {
			return []astar.Pather{target}
		}
		return neighbors
	}

	// find shortest path
	patherPath, dist, found := astar.Path(p, target)

	// convert path to list of point
	pointPath := []*P(nil)
	for _, step := range patherPath {
		pointPath = append(pointPath, step.(*P))
	}

	// Nodes are shared between routes, so their step data is only known now.
	// The path runs from the target back to the start, and the first entry is
	// the target marker.
	for i := len(pointPath) - 2; i > 0; i-- {
		pointPath[i].Speed = pointPath[i].Pos - pointPath[i+1].Pos
		pointPath[i].Steps = pointPath[i+1].Steps + 1
	}

	return pointPath, dist, found
}

func (p *P) PathNeighbors() []astar.Pather {
	p.emptyPoints.Render()
	var adjs []astar.Pather
	if p.nodes != nil {
		adjs = p.sharedNeighbors()
	} else {
		adjs = p.branchNeighbors()
	}

	// if this position is the end point, we found it
	if p.exitCondition != nil {
		return p.exitCondition(p, adjs)
	}

	return adjs
}

// sharedNeighbors returns the one node of each adjacent empty position.
func (p *P) sharedNeighbors() []astar.Pather {
	adjs := make([]astar.Pather, 0, 4)
	for _, dir := range FourDirections {
		pos := p.Pos + dir
		if _, exist := p.emptyPoints[pos]; !exist {
			continue
		}
		node, known := p.nodes[pos]
		if !known {
			node = &P{Pos: pos, Speed: dir, emptyPoints: p.emptyPoints, exitCondition: p.exitCondition, nodes: p.nodes}
			p.nodes[pos] = node
		}
		adjs = append(adjs, node)
	}
	return adjs
}

// branchNeighbors gives each neighbour its own copy of the map, so that every
// branch of an exhaustive search remembers only the positions it visited.
func (p *P) branchNeighbors() []astar.Pather {
	adjs := make([]astar.Pather, 0, 4)
	for _, dir := range FourDirections {

		// move the point to the adjacent position
		p.Speed = dir
		p.Move(1)

		// if the position is valid, explore it
		if _, exist := p.emptyPoints[p.Pos]; exist {
			newP := p.Clone()
			delete(newP.emptyPoints, p.Pos) // remove position already explored
			adjs = append(adjs, newP)
		}
		p.Move(-1)
	}
	return adjs
}

// PathNeighborCost returns the cost of the tube leading to Truck.
func (p *P) PathNeighborCost(to astar.Pather) float64 {
	return 1.0
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
func (p *P) PathEstimatedCost(to astar.Pather) float64 {
	return float64(p.Pos.DistFrom(to.(*P).Pos))
}
