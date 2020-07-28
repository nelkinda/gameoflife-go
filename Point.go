package gameoflife

import (
	"fmt"
)

// Point in a 2-dimensional coordinate system.
type Point struct {
	x int
	y int
}

// P creates a Point.
func P(x, y int) *Point {
	return &Point{x: x, y: y}
}

var neighbors = PointSetOf(
	P(-1, 1), P(0, 1), P(1, 1),
	P(-1, 0), P(1, 0),
	P(-1, -1), P(0, -1), P(1, -1),
)

func (p *Point) plus(other *Point) *Point {
	return P(p.x+other.x, p.y+other.y)
}

func (p *Point) neighbors() PointSet {
	return neighbors.Map(p.plus)
}

func (p *Point) String() string {
	return fmt.Sprintf("P(%d, %d)", p.x, p.y)
}
