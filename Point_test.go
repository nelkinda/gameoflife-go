package gameoflife

import (
	"fmt"
	"reflect"
	"testing"
)

type T testing.T

func (t *T) assertEquals(expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: <%s>\nActual: <%s>", expected, actual)
	}
}

func Test_PointString(t *testing.T) {
	expected := "P(0, 1)"
	actual := fmt.Sprintf("%s", P(0, 1))
	(*T)(t).assertEquals(expected, actual)
}

func Test_Equals(t *testing.T) {
	(*T)(t).assertEquals(P(0, 0), P(0, 0))
}

func Test_NotEquals(t *testing.T) {
	if P(0, 0) == P(0, 1) {
		t.Fail()
	}
}
