package main

import (
	"testing"
)

func TestHitP(t *testing.T) {
	r := &ray{d: vector3{0, 0.3, -1}}
	got, _, _, ok := r.HitInfo(vector3{-1, 0.2, -1.5}, vector3{1, -0.1, -1}, vector3{0.5, 1, -1.5})
	expect := vector3{0, 0.415069, -1.38356}
	if !ok || !sameVector3(got, expect) {
		t.Errorf("expect %v(true), got %v(%v)", expect, got, ok)
	}
}
