package main

import (
	"scalc/set"
	"testing"
)

func TestParseExpression(t *testing.T) {
	str := []string{"[", "SUM", "[", "DIF", "a.txt", "b.txt", "c.txt", "]", "[", "INT", "b.txt", "c.txt", "]", "]"}

	res, err := set.ParseExpression(str)
	if err != nil {
		t.Error("error should be nil")
	}

	if len(res) != 3 {
		t.Error("length of result set should be 3")
	}

	_, ok := res[1]
	if !ok {
		t.Error("result should contain 1")
	}

	_, ok = res[3]
	if !ok {
		t.Error("result should contain 3")
	}

	_, ok = res[4]
	if !ok {
		t.Error("result should contain 4")
	}
}

func TestWrongExpression(t *testing.T) {
	str := []string{"[", "SUM", "[", "DIF", "a.txt", "b.txt", "c.txt", "]", "[", "INT", "b.txt", "c.txt", "]", ""}
	_, err := set.ParseExpression(str)
	if err == nil {
		t.Error("error should not be nil")
	}
}
