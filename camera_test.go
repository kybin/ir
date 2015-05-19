package main

import (
	"testing"
	"math"
)

func TestCamera(t *testing.T) {
	cam := NewCamera(50, 41.4214, 320, 243, 1)
	// aperture
	expx, expy := 41.4214, 31.454375625
	gotx, goty := cam.Aperture()
	ign := 0.0001
	if math.Abs(expx - gotx) > ign || math.Abs(expy - goty) > ign {
		t.Errorf("camera's aperture not matched. expect:%v,%v got:%v,%v", expx, expy, gotx, goty)
	}
	// fov
	expx, expy = 0.7853989104731898, 0.6094914217849252
	gotx, goty = cam.FOV()
	if math.Abs(expx - gotx) > ign || math.Abs(expy - goty) > ign {
		t.Errorf("camera's fov not matched. expect:%v,%v got:%v,%v", expx, expy, gotx, goty)
	}
}
