package main

import (
	"math"
	"testing"
)

func TestNeg(t *testing.T) {
	got := vector3{1, 1, 1}.Neg()
	expect := vector3{-1, -1, -1}
	if got != expect {
		t.Errorf("negate error. expect:%v, got:%v", expect, got)
	}
	got = vector3{-1, -3, 19}.Neg()
	expect = vector3{1, 3, -19}
	if got != expect {
		t.Errorf("nagate error. expect:%v, got:%v", expect, got)
	}
}

func TestDot(t *testing.T) {
	got := vector3{0, 0, 1}.Dot(vector3{0, 0, 1})
	expect := float64(1)
	if got != expect {
		t.Errorf("dot error. expect:%v, got:%v", expect, got)
	}
	got = vector3{0, 0, 1}.Dot(vector3{0, 0, -1})
	expect = float64(-1)
	if got != expect {
		t.Errorf("dot error. expect:%v, got:%v", expect, got)
	}
	got = vector3{0, 1, 1}.Dot(vector3{0, 1, 0})
	expect = float64(1)
	if got != expect {
		t.Errorf("dot error. expect:%v, got:%v", expect, got)
	}
	got = vector3{1, 0, 1}.Dot(vector3{0, 1, 0})
	expect = float64(0)
	if got != expect {
		t.Errorf("dot error. expect:%v, got:%v", expect, got)
	}
	got = vector3{0.3, 0.4, 0.}.Dot(vector3{0.3, 0.4, 0})
	expect = float64(0.5)
	if math.Abs(expect-got) < 0.0001 {
		t.Errorf("dot error. expect:%v, got:%v", expect, got)
	}
}

func TestCross(t *testing.T) {
	got := vector3{1, 0, 0}.Cross(vector3{0, 1, 0})
	expect := vector3{0, 0, 1}
	if got != expect {
		t.Errorf("cross error. expect:%v, got:%v", expect, got)
	}
	got = vector3{0, 1, 0}.Cross(vector3{0, 0, 1})
	expect = vector3{1, 0, 0}
	if got != expect {
		t.Errorf("cross error. expect:%v, got:%v", expect, got)
	}
	got = vector3{0, 0, 1}.Cross(vector3{1, 0, 0})
	expect = vector3{0, 1, 0}
	if got != expect {
		t.Errorf("cross error. expect:%v, got:%v", expect, got)
	}
	got = vector3{0, 0, 1}.Cross(vector3{0.7, 0, 0.7}).Normalize()
	expect = vector3{0, 1, 0}
	if got != expect {
		t.Errorf("cross error. expect:%v, got:%v", expect, got)
	}
}

func Test2Cross(t *testing.T) {
	got := vector3{0, 0, 1}.Cross(vector3{0.7, 0, 0.7}).Cross(vector3{0, 0, 1}).Normalize()
	expect := vector3{1, 0, 0}
	if got != expect {
		t.Errorf("2 cross error. expect:%v, got:%v", expect, got)
	}
}

func TestLen(t *testing.T) {
	got := vector3{1, 0, 0}.Len()
	expect := float64(1)
	if got != expect {
		t.Errorf("length error. expect:%v, got:%v", expect, got)
	}
	got = vector3{0, 1, 0}.Len()
	expect = float64(1)
	if got != expect {
		t.Errorf("length error. expect:%v, got:%v", expect, got)
	}
	got = vector3{0, 0, 1}.Len()
	expect = float64(1)
	if got != expect {
		t.Errorf("length error. expect:%v, got:%v", expect, got)
	}
	got = vector3{1, 1, 1}.Len()
	expect = math.Sqrt(3)
	if got != expect {
		t.Errorf("length error. expect:%v, got:%v", expect, got)
	}
	got = vector3{-1, -1, -1}.Len()
	expect = math.Sqrt(3)
	if got != expect {
		t.Errorf("length error. expect:%v, got:%v", expect, got)
	}
}

func TestRotate(t *testing.T) {
	v := vector3{1, 0, 0}
	got := v.Rotate(0, 0, 0, "angle", "xyz")
	expect := vector3{1, 0, 0}
	if !got.Equal(expect) {
		t.Errorf("rotate error. expect:%v, got:%v", expect, got)
	}

	v = vector3{1, 0, 0}
	got = v.Rotate(0, 90, 0, "angle", "xyz")
	expect = vector3{0, 0, -1}
	if !got.Equal(expect) {
		t.Errorf("rotate error. expect:%v, got:%v", expect, got)
	}

	v = vector3{0, 1, 0}
	got = v.Rotate(0, 0, 90, "angle", "xyz")
	expect = vector3{-1, 0, 0}
	if !got.Equal(expect) {
		t.Errorf("rotate error. expect:%v, got:%v", expect, got)
	}

	v = vector3{0, 0, 1}
	got = v.Rotate(90, 0, 0, "angle", "xyz")
	expect = vector3{0, -1, 0}
	if !got.Equal(expect) {
		t.Errorf("rotate error. expect:%v, got:%v", expect, got)
	}
}
