package main

import (
	"testing"
	"math"
)

func TestIntersectPoint(t *testing.T) {
	a := line2{vector2{0, 0}, vector2{1, 1}}
	b := line2{vector2{-1, 0}, vector2{2, 1}}
	got, ok := a.IntersectPoint(b)
	expect := vector2{0.5, 0.5}
	if !ok {
		t.Errorf("Intersection of %v and %v is %v, got no intersection.", a, b, expect)
	} else if math.Abs(expect.x-got.x) > ignorable || math.Abs(expect.y-got.y) > ignorable {
		t.Errorf("Intersection of %v and %v is %v, got %v instead.", a, b, expect, got)
	}

	a = line2{vector2{0, -1}, vector2{0, 1}}
	b = line2{vector2{-1, 0}, vector2{1, 0}}
	got, ok = a.IntersectPoint(b)
	expect = vector2{0, 0}
	if !ok {
		t.Errorf("Intersection of %v and %v is %v, got no intersection.", a, b, expect)
	} else  if math.Abs(expect.x-got.x) > ignorable || math.Abs(expect.y-got.y) > ignorable {
		t.Errorf("Intersection of %v and %v is %v, got %v instead.", a, b, expect, got)
	}

	a = line2{vector2{0, 0}, vector2{1, 0.5}}
	b = line2{vector2{1, 0}, vector2{0.5, 1}}
	got, ok = a.IntersectPoint(b)
	expect = vector2{0.8, 0.4}
	if !ok {
		t.Errorf("Intersection of %v and %v is %v, got no intersection.", a, b, expect)
	} else if math.Abs(expect.x-got.x) > ignorable || math.Abs(expect.y-got.y) > ignorable {
		t.Errorf("Intersection of %v and %v is %v, got %v instead.", a, b, expect, got)
	}

	a = line2{vector2{-3, 1}, vector2{2, 3}}
	b = line2{vector2{-1, 0}, vector2{-1, 5}}
	got, ok = a.IntersectPoint(b)
	expect = vector2{-1, 1.8}
	if !ok {
		t.Errorf("Intersection of %v and %v is %v, got no intersection.", a, b, expect)
	} else if math.Abs(expect.x-got.x) > ignorable || math.Abs(expect.y-got.y) > ignorable {
		t.Errorf("Intersection of %v and %v is %v, got %v instead.", a, b, expect, got)
	}

	a = line2{vector2{2,1}, vector2{3,3}}
	b = line2{vector2{2,-1}, vector2{3,-3}}
	got, ok = a.IntersectPoint(b)
	if ok {
		t.Errorf("Intersection of %v and %v shold not have itersection. got %v", a, b, got)
	}

}

func TestIntersect(t *testing.T) {
	a := line2{vector2{0.66, 0.33}, vector2{0.2, 0.3}}
	b := line2{vector2{0, 0}, vector2{1, 1}}
	if !a.Intersect(b) {
		t.Errorf("%v and %v shold intersect. got no intersect", a, b)
	}
	a = line2{vector2{0.33, 0.66}, vector2{0.2, 0.3}}
	b = line2{vector2{0, 0}, vector2{0, 1}}
	if a.Intersect(b) {
		t.Errorf("%v and %v shold not intersect. got intersect", a, b)
	}
	a = line2{vector2{0.33, 0.66}, vector2{0.2, 0.3}}
	b = line2{vector2{1, 1}, vector2{0, 0}}
	if a.Intersect(b) {
		t.Errorf("%v and %v should not itersect. got intersect", a, b)
	}
	a = line2{vector2{0.33, 0.66}, vector2{0.2, 0.3}}
	b = line2{vector2{0, 1}, vector2{1, 1}}
	if a.Intersect(b) {
		t.Errorf("%v and %v shold not intersect. got intersect", a, b)
	}
	a = line2{vector2{0.66, 0.33}, vector2{0.2, 0.3}}
	b = line2{vector2{1, 1}, vector2{1, 0}}
	if a.Intersect(b) {
		t.Errorf("%v and %v shold not intersect. got intersect", a, b)
	}
}

func BenchmarkIntersect(b *testing.B) {
	l := line2{vector2{0, -1}, vector2{0, 1}}
	ll := line2{vector2{-1, 0}, vector2{1, 0}}
	for i := 0; i < b.N; i++ {
		l.Intersect(ll)
	}
}
