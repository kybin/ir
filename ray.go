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
		if !(bb.min.x <= pt.x && pt.x <= bb.max.x && bb.min.y <= pt.y && pt.y <= bb.max.y) {
			return false
		}
		switch len(*p) {
		case 3:
			// triangle's center should always inside of the triangle.
			// draw a line from center to pt and check it is intersect with any polygon edge.
			// if so, the ray does not hit the triangle.
			ct := p.Center()
			pp := *p
			l := line2{vector2{ct.x, ct.y}, pt}
			a := line2{vector2{pp[0].x, pp[0].y}, vector2{pp[1].x, pp[1].y}}
			b := line2{vector2{pp[1].x, pp[1].y}, vector2{pp[2].x, pp[2].y}}
			c := line2{vector2{pp[2].x, pp[2].y}, vector2{pp[0].x, pp[0].y}}
			if !(l.Intersect(a) || l.Intersect(b) || l.Intersect(c)) {
				return true
			}
			return false
		case 4:
			// divide the square to 2 triangles. then we can use above (triangle) approach.
			// if any triangle contains the point, it means the square contains the point.
			pp := *p
			ct := (&polygon{pp[0], pp[1], pp[2]}).Center()
			l := line2{vector2{ct.x, ct.y}, pt}
			a := line2{vector2{pp[0].x, pp[0].y}, vector2{pp[1].x, pp[1].y}}
			b := line2{vector2{pp[1].x, pp[1].y}, vector2{pp[2].x, pp[2].y}}
			c := line2{vector2{pp[2].x, pp[2].y}, vector2{pp[0].x, pp[0].y}}
			if !(l.Intersect(a) || l.Intersect(b) || l.Intersect(c)) {
				return true
			}
			ct = (&polygon{(*p)[0], (*p)[2], (*p)[3]}).Center()
			l = line2{vector2{ct.x, ct.y}, pt}
			a = line2{vector2{pp[0].x, pp[0].y}, vector2{pp[1].x, pp[1].y}}
			b = line2{vector2{pp[1].x, pp[1].y}, vector2{pp[2].x, pp[2].y}}
			c = line2{vector2{pp[2].x, pp[2].y}, vector2{pp[0].x, pp[0].y}}
			if !(l.Intersect(a) || l.Intersect(b) || l.Intersect(c)) {
				return true
			}
			return false
		default:
			panic("n-gon not supported yet.")
		}
	}
	return false
}

