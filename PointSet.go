package gameoflife

import "sort"

// PointSet is a set of Point.
type PointSet map[Point]bool

// PointSetOf creates a new PointSet based on the provided points.
func PointSetOf(points ...*Point) PointSet {
	var pointSet = make(PointSet)
	for _, point := range points {
		pointSet[*point] = true
	}
	return pointSet
}

// Union a new PointSet that is a union of all the provided pointSets.
func Union(pointSets ...PointSet) PointSet {
	union := make(PointSet)
	for _, points := range pointSets {
		union.addAll(points)
	}
	return union
}

func (pointSet PointSet) addAll(other PointSet) {
	for point := range other {
		pointSet[point] = true
	}
}

// Map maps all points of pointSet to a new PointSet using the given transform.
func (pointSet PointSet) Map(transform func(*Point) *Point) PointSet {
	mapped := make(PointSet)
	for point := range pointSet {
		mapped[*transform(&point)] = true
	}
	return mapped
}

// FlatMap maps all points of a pointSet to a new PointSet using the given transform.
func (pointSet PointSet) FlatMap(transform func(*Point) PointSet) PointSet {
	flattened := make(PointSet)
	for point := range pointSet {
		for p := range transform(&point) {
			flattened[p] = true
		}
	}
	return flattened
}

// Filter returns a new PointSet containing all points from the pointSet that pass the given filter.
func (pointSet PointSet) Filter(filter func(*Point) bool) PointSet {
	filtered := make(PointSet)
	for point := range pointSet {
		if filter(&point) {
			filtered[point] = true
		}
	}
	return filtered
}

func (pointSet PointSet) String() string {
	s := "["
	first := true
	for _, point := range pointSet.sortedKeys() {
		if !first {
			s += ", "
		}
		s += point.String()
		first = false
	}
	s += "]"
	return s
}

type points []Point

func (p points) Len() int {
	return len(p)
}

func (p points) Less(i int, j int) bool {
	p1 := p[i]
	p2 := p[j]
	return p1.x < p2.x || p1.x == p2.x && p1.y < p2.y
}

func (p points) Swap(i int, j int) {
	tmp := p[i]
	p[i] = p[j]
	p[j] = tmp
}

func (p points) Sort() {
	sort.Sort(p)
}

func (pointSet PointSet) sortedKeys() points {
	keys := make(points, 0)
	for point := range pointSet {
		keys = append(keys, point)
	}
	keys.Sort()
	return keys
}
