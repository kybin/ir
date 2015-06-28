package main

import (
	"math"
	"testing"
)

func TestSlerp(t *testing.T) {
	q1 := quaternion{vector3{1, 0, 0}, 0}
	q2 := quaternion{vector3{0, 1, 0}, 0}
	got := Slerp(q1, q2, 0.5)
	expect := quaternion{vector3{1/math.Sqrt(2), 1/math.Sqrt(2), 0}, 0}
	if !got.Equal(expect) {
		t.Errorf("expect %v, got %v", expect, got)
	}
}
