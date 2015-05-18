package main

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	id := IdentityMatrix4()
	roty := matrix4{0,0,1,0, 0,1,0,0, -1,0,0,0, 0,0,0,1}
	got := id.Multiply(roty)
	if got != roty {
		t.Errorf("matrix multiplied by identity matrix should it self. expect:%v, got:%v", roty, got)
	}
}
