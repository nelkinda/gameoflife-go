package gameoflife

import "fmt"

// Universe in a Game of Life
type Universe struct {
	rules Rules
	life  PointSet
}

// Next calculates the next generation of the Universe.
func (u *Universe) Next() *Universe {
	return &Universe{
		rules: u.rules,
		life:  u.cellsOfNextGeneration(),
	}
}

func (u *Universe) cellsOfNextGeneration() PointSet {
	return Union(u.survivingCells(), u.bornCells())
}

func (u *Universe) survivingCells() PointSet {
	return u.life.Filter(u.survives)
}

func (u *Universe) bornCells() PointSet {
	return u.deadNeighborsOfLivingCells().Filter(u.born)
}

func (u *Universe) deadNeighborsOfLivingCells() PointSet {
	return u.life.FlatMap(u.deadNeighbors)
}

func (u *Universe) survives(cell *Point) bool {
	return u.rules.Survives(u.countLiveNeighbors(cell))
}

func (u *Universe) born(cell *Point) bool {
	return u.rules.Born(u.countLiveNeighbors(cell))
}

func (u *Universe) isAlive(cell *Point) bool {
	return u.life[*cell]
}

func (u *Universe) isDead(cell *Point) bool {
	return !u.life[*cell]
}

func (u *Universe) deadNeighbors(cell *Point) PointSet {
	return cell.neighbors().Filter(u.isDead)
}

func (u *Universe) liveNeighbors(cell *Point) PointSet {
	return cell.neighbors().Filter(u.isAlive)
}

func (u *Universe) countLiveNeighbors(cell *Point) int {
	return len(u.liveNeighbors(cell))
}

func (u *Universe) String() string {
	return fmt.Sprintf("Universe{%s\n%s}", u.rules, u.life)
}
