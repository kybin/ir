package main

import (
	"testing"
)

func TestTransform(t *testing.T) {
	v := &vertex{1,0,0}
	g := &geometry{&polygon{v}}
	g.Transform(matrix4{0,1,0,0, 1,0,0,0, 0,0,1,0, 0,0,0,1})
	expect := vertex{0,1,0}
	if *v != expect {
		t.Errorf("Vertex does not move properly. expect:%v, got:%v", expect, *v)
	}
}

