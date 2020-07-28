package gameoflife

import (
	"fmt"
	"testing"
)

func Test_UniverseString(t *testing.T) {
	expected := "Universe{R 23/3\n[P(0, 1)]}"
	actual := fmt.Sprintf("%s", newUniverse(P(0, 1)))
	if expected != actual {
		t.Errorf("Expected: <%s>\nActual: <%s>", expected, actual)
	}
}
