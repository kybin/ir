package main

import (
	"testing"
	"math"
)

func TestIntersect2(t *testing.T) {
	a := line2{point2{0, 0}, point2{1, 1}}
	b := line2{point2{-1, 0}, point2{2, 1}}
	got, ok := intersect2(a, b)
	expect := point2{0.5, 0.5}
	if !ok{
		t.Errorf("Intersection of %v and %v is %v, got no intersection.")
	}
	if !ok || math.Abs(expect.x-got.x) > ignorable || math.Abs(expect.y-got.y) > ignorable {
		t.Errorf("Intersection of %v and %v is %v, got %v instead.", a, b, expect, got)
	}
}
