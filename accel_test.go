package main

// it's OK to see the result. but I need better design.

import (
	"fmt"
	"testing"
)

func TestParseOctree(t *testing.T) {
	geo := loadGeometry("geo/rubbertoy.geo")
	bb := geo.BBox()
	fmt.Println(bb)
	oct := ParseOctree(bb, geo.plys)
	for i, o := range oct.children {
		if o == nil {
			continue
		}
		fmt.Printf("%v: %v, %v, %v\n", i, o.leaf, len(o.polys), o.children)
		//fmt.Println(o.bound)
	}
}
