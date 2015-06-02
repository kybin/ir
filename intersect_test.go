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
	if !ok {
		t.Errorf("Intersection of %v and %v is %v, got no intersection.", a, b, expect)
	}
	if math.Abs(expect.x-got.x) > ignorable || math.Abs(expect.y-got.y) > ignorable {
		t.Errorf("Intersection of %v and %v is %v, got %v instead.", a, b, expect, got)
	}

	a = line2{point2{0, -1}, point2{0, 1}}
	b = line2{point2{-1, 0}, point2{1, 0}}
	got, ok = intersect2(a, b)
	expect = point2{0, 0}
	if !ok {
		t.Errorf("Intersection of %v and %v is %v, got no intersection.", a, b, expect)
	}
	if math.Abs(expect.x-got.x) > ignorable || math.Abs(expect.y-got.y) > ignorable {
		t.Errorf("Intersection of %v and %v is %v, got %v instead.", a, b, expect, got)
	}

	a = line2{point2{2,1}, point2{3,3}}
	b = line2{point2{2,-1}, point2{3,-3}}
	got, ok = intersect2(a, b)
	if ok {
		t.Errorf("Intersection of %v and %v shold not have itersection. got %v", a, b, got)
	}
}

func BenchmarkIntersect2(b *testing.B) {
	l := line2{point2{0, -1}, point2{0, 1}}
	ll := line2{point2{-1, 0}, point2{1, 0}}
	for i := 0; i < b.N; i++ {
		intersect2(l, ll)
	}
}
