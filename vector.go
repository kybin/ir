package main

type vector3 struct {
	x, y, z float64
}

func (v vector3) Transform(m matrix4) vector3 {
	a := v.x
	b := v.y
	c := v.z
	return vector3{
		a*m.aa + b*m.ba + c*m.ca,
		a*m.ab + b*m.bb + c*m.cb,
		a*m.ac + b*m.bc + c*m.cc,
	}
}
