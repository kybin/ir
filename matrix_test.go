package main

import (
	"testing"
)

func TestMatrixMultiply(t *testing.T) {
	id := IdentityMatrix4()
	roty := matrix4{0,0,1,0, 0,1,0,0, -1,0,0,0, 0,0,0,1}
	got := id.Multiply(roty)
	if got != roty {
		t.Errorf("matrix multiplied by identity matrix should it self. expect:%v, got:%v", roty, got)
	}
}

func TestMatrixTranspose(t *testing.T) {
	got := matrix4{11,12,13,14,21,22,23,24,31,32,33,34,41,42,43,44}.Transpose()
	expect := matrix4{11,21,31,41,12,22,32,42,13,23,33,43,14,24,34,44}
	if got != expect {
		t.Errorf("transpose failed. expect %v, got %v", expect, got)
	}
}
