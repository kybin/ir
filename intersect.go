package main

type point2 struct {
	x, y float64
}

type line2 struct {
	start, end point2
}

// but what should I do if the slope is infinity? (l.end.x == l.start.x)
// it still could have intersection point.
func slopeIntercept(l line2) (float64, float64) {
	slope := (l.end.y - l.start.y) / (l.end.x - l.start.x)
	intercept := l.start.y - slope*l.start.x
	return slope, intercept
}

// caculate two 2d line intersection point.
func intersect2(a, b line2) (point2, bool) {
	sa, ia := slopeIntercept(a)
	sb, ib := slopeIntercept(b)
	if sa == sb {
		return point2{}, false
	}
	slope := sb - sa
	intercept := ib - ia
	x := -intercept/slope
	if x < a.start.x || x > a.end.x || x < b.start.x || x > b.end.x {
		return point2{}, false
	}
	y := sa*x + ia
	return point2{x, y}, true
}
