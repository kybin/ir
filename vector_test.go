package main

import (
	"testing"
)

func TestVector3Transform(t *testing.T) {
	v := vector3{1,0,0}
	got := v.Transform(IdentityMatrix4())
	expect := vector3{1,0,0}
	if got != expect {
		t.Error("Vector3 transform with identity matrix should not change it's position.")
	}

	v = vector3{1,0,0}
	got = v.Transform(matrix4{0,1,0,0, 1,0,0,0, 0,0,1,0, 0,0,0,1})
	expect = vector3{0,1,0}
	if got != expect {
		t.Errorf("Vector3 does not move properly. expect:%v, got:%v", expect, got)
	}
}

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
