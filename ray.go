package main

type ray struct {
	dir vector3
}

// TODO :
//  better hit algorithm. with current method, it's hard to find hit position.
//  current hit position is not accurate.
func (r *ray) Sample(scn *scene) (sample float64, hit bool) {
	if r.dir.z == 0 {
		return 0, false
	}
	// flatten the ray dir so ray.z is 1.
	// after that, we can compare x, y components.
	dir := r.dir.Div(r.dir.z)
	pt := vector2{dir.x, dir.y}
	hitpos := float64(-1000000000)
	for _, p := range (*scn.geo) {
		// check if the pt is place outside of bound.
		pz := p.Center().z
		switch len(p.vts) {
		case 3:
			// triangle's center should always inside of the triangle.
			// draw a line from center to pt and check it is intersect with any polygon edge.
			// if so, the ray does not hit the triangle.
			v0 := vector3{p.vts[0].P.x/p.vts[0].P.z, p.vts[0].P.y/p.vts[0].P.z, 1}
			v1 := vector3{p.vts[1].P.x/p.vts[1].P.z, p.vts[1].P.y/p.vts[1].P.z, 1}
			v2 := vector3{p.vts[2].P.x/p.vts[2].P.z, p.vts[2].P.y/p.vts[2].P.z, 1}
			np := NewPolygon(NewVertex(v0), NewVertex(v1), NewVertex(v2))
			bb := np.BBox()
			if !((bb.min.x <= pt.x && pt.x <= bb.max.x) && (bb.min.y <= pt.y && pt.y <= bb.max.y)) {
				continue
			}
			ct := np.Center()
			l := line2{vector2{ct.x, ct.y}, pt}
			a := line2{vector2{v0.x, v0.y}, vector2{v1.x, v1.y}}
			b := line2{vector2{v1.x, v1.y}, vector2{v2.x, v2.y}}
			c := line2{vector2{v2.x, v2.y}, vector2{v0.x, v0.y}}
			if !(l.Intersect(a) || l.Intersect(b) || l.Intersect(c)) {
				hit = true
				if pz > hitpos {
					hitpos = pz
					sample = scn.lit.dir.Dot(p.Normal())
				}
			}
			continue
		case 4:
			// divide the square to 2 triangles. then we can use above (triangle) approach.
			// if any triangle contains the point, it means the square contains the point.
			v0 := vector3{p.vts[0].P.x/p.vts[0].P.z, p.vts[0].P.y/p.vts[0].P.z, 1}
			v1 := vector3{p.vts[1].P.x/p.vts[1].P.z, p.vts[1].P.y/p.vts[1].P.z, 1}
			v2 := vector3{p.vts[2].P.x/p.vts[2].P.z, p.vts[2].P.y/p.vts[2].P.z, 1}
			v3 := vector3{p.vts[3].P.x/p.vts[3].P.z, p.vts[3].P.y/p.vts[3].P.z, 1}
			np := NewPolygon(NewVertex(v0), NewVertex(v1), NewVertex(v2), NewVertex(v3))
			bb := np.BBox()
			if !((bb.min.x <= pt.x && pt.x <= bb.max.x) && (bb.min.y <= pt.y && pt.y <= bb.max.y)) {
				continue
			}
			ct := NewPolygon(NewVertex(v0), NewVertex(v1), NewVertex(v2)).Center()
			l := line2{vector2{ct.x, ct.y}, pt}
			a := line2{vector2{v0.x, v0.y}, vector2{v1.x, v1.y}}
			b := line2{vector2{v1.x, v1.y}, vector2{v2.x, v2.y}}
			c := line2{vector2{v2.x, v2.y}, vector2{v0.x, v0.y}}
			if !(l.Intersect(a) || l.Intersect(b) || l.Intersect(c)) {
				hit = true
				if pz > hitpos {
					hitpos = pz
					sample = scn.lit.dir.Dot(p.Normal())
				}
			}
			ct = NewPolygon(NewVertex(v0), NewVertex(v2), NewVertex(v3)).Center()
			l = line2{vector2{ct.x, ct.y}, pt}
			a = line2{vector2{v0.x, v0.y}, vector2{v2.x, v2.y}}
			b = line2{vector2{v2.x, v2.y}, vector2{v3.x, v3.y}}
			c = line2{vector2{v3.x, v3.y}, vector2{v0.x, v0.y}}
			if !(l.Intersect(a) || l.Intersect(b) || l.Intersect(c)) {
				hit = true
				if pz > hitpos {
					hitpos = pz
					sample = scn.lit.dir.Dot(p.Normal())
				}
			}
			continue
		default:
			panic("n-gon not supported yet.")
		}
	}
	if !hit {
		return 0, false
	}
	return sample, hit
}

