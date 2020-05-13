package set

import (
	"testing"
)

func TestUnion(t *testing.T) {
	setA := NewSet(1, 2, 3)
	setB := NewSet(2, 3, 4)
	setC := NewSet(3, 4, 5)

	res := Union(setA, setB, setC)

	if len(res) != 5 {
		t.Error("Length of result have to be equal 5")
	}
}

func TestIntersection(t *testing.T) {
	setA := NewSet(1, 2, 3)
	setB := NewSet(2, 3, 4)
	setC := NewSet(3, 4, 5)

	res := Intersect(setA, setB, setC)

	if len(res) != 1 {
		t.Error("Length of result have to be equal 1")
	}

	_, ok := res[3]
	if !ok {
		t.Error("res have to be queal 3")
	}
}

func TestDiff(t *testing.T) {
	setA := NewSet(1, 2, 3)
	setB := NewSet(2, 3, 4)
	setC := NewSet(3, 4, 5)

	res := Diff(setA, setB, setC)
	if len(res) != 1 {
		t.Error("Length of result have to be equal 1")
	}

	_, ok := res[1]
	if !ok {
		t.Error("res have to be queal 1")
	}
}
