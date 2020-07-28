package gameoflife

import "fmt"

type intSet map[int]bool

// Rules describes how a Game of Life universe determines its next state.
type Rules interface {
	// Survives returns true if a live cell with the specified number of liveNeighbors survives to the next generation.
	Survives(liveNeighbors int) bool

	// Born returns true if a dead cell with the specified number of liveNeighbors is born into the next generation.
	Born(liveNeighbors int) bool
}

type standardRules struct {
	liveNeighborsForSurvival intSet
	liveNeighborsForBirth    intSet
}

// ConwayRules are the rules for Conway's Game of Life (2 or 3 live neighbors for survival, 3 live neighbors for birth)
var ConwayRules = standardRules{
	liveNeighborsForSurvival: intSetOf(2, 3),
	liveNeighborsForBirth:    intSetOf(3),
}

// HighlifeRules are the rules for Highlife (2 or 3 live neighbors for survival, 3 or 6 live neighbors for birth)
var HighlifeRules = standardRules{
	liveNeighborsForSurvival: intSetOf(2, 3),
	liveNeighborsForBirth:    intSetOf(3, 6),
}

func intSetOf(numbers ...int) intSet {
	m := make(intSet)
	for _, number := range numbers {
		m[number] = true
	}
	return m
}

func (rules standardRules) Survives(liveNeighbors int) bool {
	return rules.liveNeighborsForSurvival[liveNeighbors]
}

func (rules standardRules) Born(liveNeighbors int) bool {
	return rules.liveNeighborsForBirth[liveNeighbors]
}

func (rules standardRules) String() string {
	return fmt.Sprintf("R 23/3")
}
