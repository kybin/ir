package main

import (
	"math"
)

type line2 struct {
	start, end vector2
}

// return line's slope and Y intercept.
// if the slope is stiff, it may less correct.
func (l line2) SlopeIntercept() (float64, float64) {
	slope := (l.end.y - l.start.y) / (l.end.x - l.start.x)
	intercept := l.start.y - slope*l.start.x
	return slope, intercept
}

// caculate two 2d line intersection point.
func (a line2) Intersect(b line2) bool {
	if math.Abs(a.end.y - a.start.y) > math.Abs(a.end.x - a.start.x) {
		a.start.x, a.start.y = a.start.y, a.start.x
		a.end.x, a.end.y = a.end.y, a.end.x
		b.start.x, b.start.y = b.start.y, b.start.x
		b.end.x, b.end.y = b.end.y, b.end.x
	}
	sa, ia := a.SlopeIntercept()
	// shear b to y direction.
	b.start.y = b.start.y - (sa * b.start.x) - ia
	b.end.y = b.end.y - (sa * b.end.x) - ia
	if math.Signbit(b.start.y) == math.Signbit(b.end.y) {
		return false
	}
	return true
}

// caculate two 2d line intersection point.
func (a line2) IntersectPoint(b line2) (vector2, bool) {
	swaped := false
	if math.Abs(a.end.y - a.start.y) > math.Abs(a.end.x - a.start.x) {
		swaped = true
		a.start.x, a.start.y = a.start.y, a.start.x
		a.end.x, a.end.y = a.end.y, a.end.x
		b.start.x, b.start.y = b.start.y, b.start.x
		b.end.x, b.end.y = b.end.y, b.end.x
	}
	sa, ia := a.SlopeIntercept()
	// shear b to y direction.
	b.start.y = b.start.y - (sa * b.start.x) - ia
	b.end.y = b.end.y - (sa * b.end.x) - ia
	if math.Signbit(b.start.y) == math.Signbit(b.end.y) {
		return vector2{}, false
	}
	// find x if y == 0
	tb := math.Abs(b.start.y) / math.Abs(b.end.y - b.start.y)
	x := tb * (b.end.x - b.start.x) + b.start.x
	y := sa * x + ia

	if swaped {
		x, y = y, x
	}
	return vector2{x, y}, true
}
