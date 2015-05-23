package main

type vertex vector3

func (v *vertex) Transform(m matrix4) {
	*v = vertex(vector3(*v).Transform(m))
}

type polygon []*vertex

func (p polygon) Transform(m matrix4) {
	for _, v := range p {
		v.Transform(m)
	}
}

type geometry struct {
	xform matrix4
	prims []*polygon // TODO : nurbs, curve
}

func (g geometry) Transform(m matrix4) {
	for _, p := range g.prims {
		p.Transform(m)
	}
}
