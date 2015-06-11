package main

import (
	"testing"
)

func TestRotate(t *testing.T) {
	v := vector3{1,0,0}
	got := v.Rotate(0,0,0,"angle","xyz")
	expect := vector3{1,0,0}
	if !got.Equal(expect) {
		t.Errorf("rotate error. expect:%v, got:%v", expect, got)
	}

	v = vector3{1,0,0}
	got = v.Rotate(0,90,0,"angle","xyz")
	expect = vector3{0,0,-1}
	if !got.Equal(expect) {
		t.Errorf("rotate error. expect:%v, got:%v", expect, got)
	}

	v = vector3{0,1,0}
	got = v.Rotate(0,0,90,"angle","xyz")
	expect = vector3{-1,0,0}
	if !got.Equal(expect) {
		t.Errorf("rotate error. expect:%v, got:%v", expect, got)
	}

	v = vector3{0,0,1}
	got = v.Rotate(90,0,0,"angle","xyz")
	expect = vector3{0,-1,0}
	if !got.Equal(expect) {
		t.Errorf("rotate error. expect:%v, got:%v", expect, got)
	}
}
