package main

import (
	"testing"
)

func TestTransform(t *testing.T) {
	v := NewVertex(vector3{1, 0, 0})
	g := NewGeometry(NewPolygon(v))

	tr := matrix4{
		0, 1, 0, 0,
		1, 0, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	g.Transform(tr)

	expect := vector3{0, 1, 0}
	if v.P != expect {
		t.Errorf("Vertex does not move properly. expect:%v, got:%v", expect, *v)
	}
}
