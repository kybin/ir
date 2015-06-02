package main

import (
	"math"
)

type point2 struct {
	x, y float64
}

type line2 struct {
	start, end point2
}

// return line's slope and Y intercept.
// if the slope is stiff, it may less correct.
func slopeIntercept(l line2) (float64, float64) {
	slope := (l.end.y - l.start.y) / (l.end.x - l.start.x)
	intercept := l.start.y - slope*l.start.x
	return slope, intercept
}

func swapXY(l line2) line2 {
	l.start.x, l.start.y = l.start.y, l.start.x
	l.end.x, l.end.y = l.end.y, l.end.x
	return l
}

// caculate two 2d line intersection point.
func intersect2(a, b line2) (point2, bool) {
	swaped := false
	xlen := math.Abs(a.end.x-a.start.x)
	ylen := math.Abs(a.end.y-a.start.y)
	if ylen > xlen {
		a = swapXY(a)
		b = swapXY(b)
		swaped = true
	}
	sa, ia := slopeIntercept(a)
	// shear b to y direction.
	b.start.y = b.start.y - (sa*b.start.x) - ia
	b.end.y = b.end.y - (sa*b.end.x) - ia
	if math.Signbit(b.start.y) == math.Signbit(b.end.y) {
		return point2{}, false
	}
	// find x if y == 0
	tb := b.start.y / math.Abs(b.end.y - b.start.y)
	x := tb * math.Abs(b.end.x - b.start.x) + b.start.x
	// b is sheared. so find y with a.
	ta := x / math.Abs(a.end.x - a.start.x)
	y := ta / (a.end.x - a.start.x)
	if swaped {
		x, y = y, x
	}
	return point2{x, y}, true
}
