package main

import (
	"testing"
)

func TestRayHit(t *testing.T) {
	geo := &geometry{&polygon{NewVertex(0, 0, 1), NewVertex(0, 1, 1), NewVertex(1, 1, 1), NewVertex(1, 0, 1)}}
	r := &ray{vector3{0.8, 0, 0.6}}
	got := r.Hit(geo)
	if got != false {
		t.Errorf("ray hit?!")
	}
	r = &ray{vector3{0.2, 0.3, 1}}
	got = r.Hit(geo)
	if got != true {
		t.Errorf("ray not hit?!")
	}
}
