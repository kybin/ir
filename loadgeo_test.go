package main

import (
	. "fmt"
	"testing"
)

func TestLoadGeo(t *testing.T) {
	geo := loadGeometry("geo/box.geo")
	Println(geo)
	for _, pt := range geo.pts {
		Println(pt)
	}
	for _, vt := range geo.vts {
		Println(vt)
	}
	for _, ply := range geo.plys {
		Println(ply)
		Println(ply.Vertices())
		Println(ply.Points())
	}
}
