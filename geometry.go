package main

type Attributer interface {
	SetFloatAttr(name string, x float64)
	SetVectorAttr(name string, xs ...float64)
}

type point struct {
	geo *geometry
	P   vector3
	w   float64
	v2a map[string]vector2
	v3a map[string]vector3
	fa  map[string]float64
	sa  map[string]string
}

func NewPoint(P vector3) *point {
	return &point{
		geo: nil,
		P:   P,
		w:   1,
		v2a: make(map[string]vector2),
		v3a: make(map[string]vector3),
		fa:  make(map[string]float64),
		sa:  make(map[string]string),
	}
}

func (pt *point) SetFloatAttr(name string, x float64) {
	pt.fa[name] = x
}

func (pt *point) SetVectorAttr(name string, xs ...float64) {
	switch len(xs) {
	case 2:
		pt.v2a[name] = vector2{xs[0], xs[1]}
	case 3:
		pt.v3a[name] = vector3{xs[0], xs[1], xs[2]}
	default:
		panic("can parse [2..3]string only")
	}
}

func (pt *point) Transform(m matrix4) {
	x, y, z, w := pt.P.x, pt.P.y, pt.P.z, pt.w
	pt.P = vector3{
		x*m.aa + y*m.ab + z*m.ac + w*m.ad,
		x*m.ba + y*m.bb + z*m.bc + w*m.bd,
		x*m.ca + y*m.cb + z*m.cc + w*m.cd,
	}
	pt.w = x*m.da + y*m.db + z*m.dc + w*m.dd

	for key, val := range pt.v3a {
		pt.v3a[key] = val.MultM4(m)
	}
}

type vertex struct {
	geo  *geometry
	ptid int
	v2a  map[string]vector2
	v3a  map[string]vector3
	fa   map[string]float64
	sa   map[string]string
}

func NewVertex(ptid int) *vertex {
	return &vertex{
		geo:  nil, // will be set in NewGeometry()
		ptid: ptid,
		v2a:  make(map[string]vector2),
		v3a:  make(map[string]vector3),
		fa:   make(map[string]float64),
		sa:   make(map[string]string),
	}
}

func (v *vertex) SetFloatAttr(name string, x float64) {
	v.fa[name] = x
}

func (v *vertex) SetVectorAttr(name string, xs ...float64) {
	switch len(xs) {
	case 2:
		v.v2a[name] = vector2{xs[0], xs[1]}
	case 3:
		v.v3a[name] = vector3{xs[0], xs[1], xs[2]}
	default:
		panic("can parse [2..3]string only")
	}
}

type polygon struct {
	geo   *geometry
	vtids []int
	v2a   map[string]vector2
	v3a   map[string]vector3
	fa    map[string]float64
	sa    map[string]string
}

func NewPolygon(vtids []int) *polygon {
	return &polygon{
		geo:   nil, // will be set in NewGeometry()
		vtids: vtids,
		v2a:   make(map[string]vector2),
		v3a:   make(map[string]vector3),
		fa:    make(map[string]float64),
		sa:    make(map[string]string),
	}
}

// return vertices not just their ids.
func (p *polygon) Vertices() []*vertex {
	vts := make([]*vertex, len(p.vtids))
	for i, id := range p.vtids {
		vts[i] = p.geo.vts[id]
	}
	return vts
}

// return points not just their ids. (unique, but unsorted by order)
func (p *polygon) Points() []*point {
	pts := make([]*point, 0)
	ids := make([]int, 0)
VERT:
	for _, vt := range p.Vertices() {
		for _, id := range ids {
			if id == vt.ptid {
				continue VERT
			}
		}
		ids = append(ids, vt.ptid)
		pts = append(pts, p.geo.pts[vt.ptid])
	}
	return pts
}

func (p *polygon) SetFloatAttr(name string, x float64) {
	p.fa[name] = x
}

func (p *polygon) SetVectorAttr(name string, xs ...float64) {
	switch len(xs) {
	case 2:
		p.v2a[name] = vector2{xs[0], xs[1]}
	case 3:
		p.v3a[name] = vector3{xs[0], xs[1], xs[2]}
	default:
		panic("can parse [2..3]string only")
	}
}

/*
// this method returns geometric normal.
func (p *polygon) Normal() vector3 {
	switch len(p.pts) {
	case 0, 1, 2:
		return vector3{}
	default:
		v1 := p.pts[1].P.Sub(p.pts[0].P).Normalize()
		v2 := p.pts[2].P.Sub(p.pts[1].P).Normalize()
		return v1.Cross(v2).Normalize()
	}
}

func (p *polygon) Center() vector3 {
	switch len(p.pts) {
	case 0:
		return vector3{0, 0, 0}
	default:
		center := p.pts[0].P
		for _, pt := range p.pts[1:] {
			center = center.Add(pt.P)
		}
		center = center.Div(float64(len(p.pts)))
		return center
	}
}
*/

func (p *polygon) BBox() bbox {
	pts := p.Points()
	switch len(pts) {
	case 0:
		panic("no vetices in polygon")
	case 1:
		return bbox{pts[0].P, pts[0].P}
	default:
		min := pts[0].P
		max := min
		for _, pt := range pts[1:] {
			x := pt.P.x
			if min.x > x {
				min.x = x
			} else if max.x < x {
				max.x = x
			}
			y := pt.P.y
			if min.y > y {
				min.y = y
			} else if max.y < y {
				max.y = y
			}
			z := pt.P.z
			if min.z > z {
				min.z = z
			} else if max.z < z {
				max.z = z
			}
		}
		return bbox{min, max}
	}
}

func (p *polygon) BSphere() bsphere {
	pts := p.Points()
	switch len(pts) {
	case 0:
		panic("no vetices in polygon")
	case 1:
		return bsphere{pts[0].P, 0}
	default:
		o := vector3{}
		for _, pt := range pts {
			o = o.Add(pt.P)
		}
		o = o.Div(float64(len(pts)))
		var r float64 = 0
		for _, pt := range pts {
			rr := o.Sub(pt.P).Len()
			if rr > r {
				r = rr
			}
		}
		return bsphere{o, r}
	}
}

// TODO : nurbs, curve
type geometry struct {
	pts  []*point
	vts  []*vertex
	plys []*polygon
	v2a  map[string]vector2
	v3a  map[string]vector3
	fa   map[string]float64
	sa   map[string]string
	bb   bbox
}

func NewGeometry(pts []*point, vts []*vertex, plys []*polygon) *geometry {
	g := &geometry{
		pts:  pts,
		vts:  vts,
		plys: plys,
		v2a:  make(map[string]vector2),
		v3a:  make(map[string]vector3),
		fa:   make(map[string]float64),
		sa:   make(map[string]string),
	}

	// set parent
	for _, pt := range pts {
		pt.geo = g
	}
	for _, vt := range vts {
		vt.geo = g
	}
	for _, ply := range plys {
		ply.geo = g
	}

	// bbox
	if len(plys) > 0 {
		g.bb = plys[0].BBox()
	}
	for _, p := range plys[1:] {
		g.bb = g.bb.Union(p.BBox())
	}
	return g
}

func (g *geometry) SetFloatAttr(name string, x float64) {
	g.fa[name] = x
}

func (g *geometry) SetVectorAttr(name string, xs ...float64) {
	switch len(xs) {
	case 2:
		g.v2a[name] = vector2{xs[0], xs[1]}
	case 3:
		g.v3a[name] = vector3{xs[0], xs[1], xs[2]}
	default:
		panic("can parse [2..3]string only")
	}
}

func (g *geometry) Transform(m matrix4) {
	for _, pt := range g.pts {
		pt.Transform(m)
	}
	for key, val := range g.v3a {
		g.v3a[key] = val.MultM4(m)
	}
	// TODO: MultM4 for child element attributes as well (ply, vt)
}

func (g *geometry) BBox() bbox {
	if len(g.plys) == 0 {
		panic("cannot check bounding box for geomtry. no polygons in the geomtry")
	}
	bb := g.plys[0].BBox()
	for _, p := range g.plys[1:] {
		bb = bb.Union(p.BBox())
	}
	return bb
}
