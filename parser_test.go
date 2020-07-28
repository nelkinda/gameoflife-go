package gameoflife

import (
	"fmt"
	"reflect"
	"testing"
)

//noinspection GoSnakeCaseUsage
func parseSimplifiedLife1_05(spec string) *Universe {
	var cells = make(PointSet)
	var line = 0
	var column = 0
	for _, c := range spec {
		if c == '\n' {
			line++
			column = 0
		} else if c == '*' {
			cells[*P(column, line)] = true
			column++
		} else if c == '.' {
			column++
		} else {
			panic(fmt.Sprintf("Unexpected character '%v' at line %d, column %d", c, line, column))
		}
	}
	return &Universe{rules: ConwayRules, life: cells}
}

func newUniverse(points ...*Point) *Universe {
	return &Universe{rules: ConwayRules, life: PointSetOf(points...)}
}

func parses(t *testing.T, spec string, points ...*Point) {
	if !reflect.DeepEqual(newUniverse(points...), parseSimplifiedLife1_05(spec)) {
		t.Fail()
	}
}

func Test_Parser(t *testing.T) {
	parses(t, "")
	parses(t, "*", P(0, 0))
	parses(t, "**", P(0, 0), P(1, 0))
	parses(t, "*\n*", P(0, 0), P(0, 1))
	parses(t, "*.*", P(0, 0), P(2, 0))
}
