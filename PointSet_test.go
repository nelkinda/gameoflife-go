package gameoflife

import (
	"fmt"
	"testing"
)

func Test_PointsString(t *testing.T) {
	expected := "[P(0, 0), P(0, 1), P(0, 2), P(0, 3), P(0, 4)]"
	actual := fmt.Sprintf("%s", PointSetOf(P(0, 1), P(0, 0), P(0, 4), P(0, 3), P(0, 2)))
	(*T)(t).assertEquals(expected, actual)
}
