package main

type ray struct {
	dir vector3
}

// input geometry should flattened before loaded to here. (for efficiency)
// that means the geomtry's z value will treated as 1.
func (r *ray) Hit(g *geometry) bool {
	if r.dir.z == 0 {
		return false
	}
	// flatten the ray dir so ray.z is 1.
	// after that, we can compare x, y components.
	dir := r.dir.Div(r.dir.z)
	pt := vector2{dir.x, dir.y}
	for _, p := range *g {
		// check if the pt is place outside of bound.
		bb := p.BBox()
		if !((bb.min.x <= pt.x && pt.x <= bb.max.x) && (bb.min.y <= pt.y && pt.y <= bb.max.y)) {
			continue
		}
		switch len(p.vts) {
		case 3:
			// triangle's center should always inside of the triangle.
			// draw a line from center to pt and check it is intersect with any polygon edge.
			// if so, the ray does not hit the triangle.
			ct := p.Center()
			l := line2{vector2{ct.x, ct.y}, pt}
			a := line2{vector2{p.vts[0].P.x, p.vts[0].P.y}, vector2{p.vts[1].P.x, p.vts[1].P.y}}
			b := line2{vector2{p.vts[1].P.x, p.vts[1].P.y}, vector2{p.vts[2].P.x, p.vts[2].P.y}}
			c := line2{vector2{p.vts[2].P.x, p.vts[2].P.y}, vector2{p.vts[0].P.x, p.vts[0].P.y}}
			if !(l.Intersect(a) || l.Intersect(b) || l.Intersect(c)) {
				return true
			}
			continue
		case 4:
			// divide the square to 2 triangles. then we can use above (triangle) approach.
			// if any triangle contains the point, it means the square contains the point.
			ct := NewPolygon(p.vts[0], p.vts[1], p.vts[2]).Center()
			l := line2{vector2{ct.x, ct.y}, pt}
			a := line2{vector2{p.vts[0].P.x, p.vts[0].P.y}, vector2{p.vts[1].P.x, p.vts[1].P.y}}
			b := line2{vector2{p.vts[1].P.x, p.vts[1].P.y}, vector2{p.vts[2].P.x, p.vts[2].P.y}}
			c := line2{vector2{p.vts[2].P.x, p.vts[2].P.y}, vector2{p.vts[0].P.x, p.vts[0].P.y}}
			if !(l.Intersect(a) || l.Intersect(b) || l.Intersect(c)) {
				return true
			}
			ct = NewPolygon(p.vts[0], p.vts[2], p.vts[3]).Center()
			l = line2{vector2{ct.x, ct.y}, pt}
			a = line2{vector2{p.vts[0].P.x, p.vts[0].P.y}, vector2{p.vts[2].P.x, p.vts[2].P.y}}
			b = line2{vector2{p.vts[2].P.x, p.vts[2].P.y}, vector2{p.vts[3].P.x, p.vts[3].P.y}}
			c = line2{vector2{p.vts[3].P.x, p.vts[3].P.y}, vector2{p.vts[0].P.x, p.vts[0].P.y}}
			if !(l.Intersect(a) || l.Intersect(b) || l.Intersect(c)) {
				return true
			}
			continue
		default:
			panic("n-gon not supported yet.")
		}
	}
	return false
}

