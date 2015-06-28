package main

import (
	"testing"
)

func TestTransform(t *testing.T) {
	v := NewVertex(1, 0, 0)
	g := &geometry{&polygon{v}}

	tr := matrix4{
		0, 1, 0, 0,
		1, 0, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	g.Transform(tr)

	expect := vertex{0, 1, 0, 1}
	if *v != expect {
		t.Errorf("Vertex does not move properly. expect:%v, got:%v", expect, *v)
	}
}
