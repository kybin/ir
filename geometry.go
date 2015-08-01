package main

type vertex struct {
	P vector3
	w float64
	v3a map[string]vector3
	fa map[string]float64
	parent *polygon
}

func NewVertex(P vector3) *vertex {
	v3a := make(map[string]vector3)
	fa := make(map[string]float64)
	return &vertex{
		P: P,
		w: 1,
		v3a: v3a,
		fa: fa,
		parent: nil, // set when contained to polygon.
	}
}

func (v *vertex) Pos() vector3 {
	return v.P
}

func (v *vertex) Transform(m matrix4) {
	x, y, z, w := v.P.x, v.P.y, v.P.z, v.w
	v.P = vector3{
		x*m.aa + y*m.ab + z*m.ac + w*m.ad,
		x*m.ba + y*m.bb + z*m.bc + w*m.bd,
		x*m.ca + y*m.cb + z*m.cc + w*m.cd,
	}
	v.w = x*m.da + y*m.db + z*m.dc + w*m.dd
}

func (v *vertex) Normal() vector3 {
	if v.parent == nil {
		return vector3{0, 0, 0}
	}
	val, _ := v.parent.v3a["N"]
	return val
}

type polygon struct {
	vts []*vertex
	v3a map[string]vector3
	fa map[string]float64
}

func NewPolygon(vts ...*vertex) *polygon {
	v3a := make(map[string]vector3)
	fa := make(map[string]float64)
	ply := &polygon{
		vts: vts,
		v3a: v3a,
		fa: fa,
	}
	for _, v := range vts {
		v.parent = ply
	}
	return ply
}

func (p *polygon) Transform(m matrix4) {
	for _, v := range p.vts {
		v.Transform(m)
	}
}

// this method returns "N" attribute rather than exact normal.
// if you want it, call this after CalculateNormal.
func (p *polygon) Normal() vector3 {
	val, _ := p.v3a["N"]
	return val
}

func (p *polygon) CalculateNormal() {
	vts := p.vts
	switch len(vts) {
	case 0, 1, 2:
		return
	default:
		v1 := vts[1].P.Sub(vts[0].P).Normalize()
		v2 := vts[2].P.Sub(vts[1].P).Normalize()
		p.v3a["N"] = v1.Cross(v2).Normalize()
	}
}

func (p *polygon) Center() vector3 {
	switch len(p.vts) {
	case 0:
		return vector3{0, 0, 0}
	default:
		center := vector3{0, 0, 0}
		for _, v := range p.vts {
			center = center.Add(v.Pos())
		}
		center = center.Div(float64(len(p.vts)))
		return center
	}
}

func (p *polygon) BBox() bbox3 {
	switch len(p.vts) {
	case 0:
		return bbox3{vector3{0, 0, 0}, vector3{0, 0, 0}}
	case 1:
		return bbox3{p.vts[0].P, p.vts[0].P}
	default:
		min := p.vts[0].P
		max := min
		for _, v := range p.vts[1:] {
			x := v.P.x
			if min.x > x {
				min.x = x
			} else if max.x < x {
				max.x = x
			}
			y := v.P.y
			if min.y > y {
				min.y = y
			} else if max.y < y {
				max.y = y
			}
			z := v.P.z
			if min.z > z {
				min.z = z
			} else if max.z < z {
				max.z = z
			}
		}
		return bbox3{min, max}
	}
}

type geometry []*polygon // TODO : nurbs, curve

func (g *geometry) Transform(m matrix4) {
	for _, p := range *g {
		p.Transform(m)
	}
}

func (g *geometry) CalculateNormal() {
	for _, p := range *g {
		p.CalculateNormal()
	}
}
