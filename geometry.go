package main

type vertex vector4

func NewVertex(x, y, z float64) *vertex {
	return &vertex{x, y, z, 1}
}

func (v *vertex) Transform(m matrix4) {
	*v = vertex{
		v.x*m.aa + v.y*m.ab + v.z*m.ac + v.w*m.ad,
		v.x*m.ba + v.y*m.bb + v.z*m.bc + v.w*m.bd,
		v.x*m.ca + v.y*m.cb + v.z*m.cc + v.w*m.cd,
		v.x*m.da + v.y*m.db + v.z*m.dc + v.w*m.dd,
	}
}

type polygon []*vertex

func (p polygon) Transform(m matrix4) {
	for _, v := range p {
		v.Transform(m)
	}
}

type geometry []*polygon // TODO : nurbs, curve

func (g geometry) Transform(m matrix4) {
	for _, p := range g {
		p.Transform(m)
	}
}
