package main

type vertex vector4

func NewVertex(x, y, z float64) *vertex {
	return &vertex{x, y, z, 1}
}

func (v *vertex) Position() vector3 {
	return vector3{v.x, v.y, v.z}
}

func (v *vertex) Transform(m matrix4) {
	*v = vertex{
		v.x*m.aa + v.y*m.ab + v.z*m.ac + v.w*m.ad,
		v.x*m.ba + v.y*m.bb + v.z*m.bc + v.w*m.bd,
		v.x*m.ca + v.y*m.cb + v.z*m.cc + v.w*m.cd,
		v.x*m.da + v.y*m.db + v.z*m.dc + v.w*m.dd,
	}
}

// Note :
//  some polygon methods will panic if polygon has no vertex.
//  how to avoid it? (polygon should have vertex?)
type polygon []*vertex

func (p *polygon) Transform(m matrix4) {
	for _, v := range *p {
		v.Transform(m)
	}
}

func (p *polygon) Center() vector3 {
	center := (*p)[0].Position()
	for _, v := range (*p)[1:] {
		center = center.Add(v.Position())
	}
	center = center.Div(float64(len(*p)))
	return center
}

func (p *polygon) BBox() bbox3 {
	min := (*p)[0].Position()
	max := (*p)[0].Position()
	for _, v := range (*p)[1:] {
		if min.x > v.x {
			min.x = v.x
		} else if max.x < v.x {
			max.x = v.x
		}
		if min.y > v.y {
			min.y = v.y
		} else if max.y < v.y {
			max.y = v.y
		}
		if min.z > v.z {
			min.z = v.z
		} else if max.z < v.z {
			max.z = v.z
		}
	}
	return bbox3{min, max}
}
type geometry []*polygon // TODO : nurbs, curve

func (g geometry) Transform(m matrix4) {
	for _, p := range g {
		p.Transform(m)
	}
}
