package main

import (
	"testing"
)

func TestMixVector3(t *testing.T) {
	a := vector3{-1, -1, -1}
	b := vector3{1, 1, 1}
	got := mixVector3(a, b, 0.25)
	expect := vector3{-0.5, -0.5, -0.5}
	if !sameVector3(got, expect) {
		t.Errorf("not expected.")
	}
}
