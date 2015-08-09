package main

import (
	"math"
)

func sameVector3(a, b vector3) bool {
	return math.Abs(a.x-b.x) < 0.00001 && math.Abs(a.y-b.y) < 0.00001 && math.Abs(a.z-b.z) < 0.00001
}
