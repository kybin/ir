package main

type vertex vector3

func (v *vertex) Transform(m matrix4) {
	vv := vector3{v.x, v.y, v.z}.Transform(m)
	v =  &vertex{vv.x, vv.y, vv.z}
}

type polygon []*vertex

func (p polygon) Transform(m matrix4) {
	for _, v := range p {
		v.Transform(m)
	}
}

type geometry []*polygon

func (g geometry) Transform(m matrix4) {
	for _, p := range g {
		p.Transform(m)
	}
}
