package main

import (
	"math"
	"testing"
)

func TestCamera(t *testing.T) {
	cam := camera{
		P:     vector3{0, 0, 0},
		front: vector3{0, 0, 1},
		right: vector3{1, 0, 0},
		up:    vector3{0, 1, 0},
		focal: 50,
		aptx:  41.4214,
		resx:  320,
		resy:  243,
		near:  0.0001,
		far:   10000,
	}
	// aperture
	exp := 31.454376
	got := cam.Apty()
	if math.Abs(exp-got) > 0.0001 {
		t.Errorf("camera's Y-aperture not matched. expect:%v, got:%v", exp, got)
	}
	// fov
	exp = 0.609491
	got = cam.FOV()
	if math.Abs(exp-got) > 0.0001 {
		t.Errorf("camera's fov not matched. expect:%v, got:%v", exp, got)
	}
}
