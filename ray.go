package main

import (
	"math"
	"image"
)

type ray struct {
	o vector3
	d vector3
}

// NewRay return a new ray.
// Given 'd' will normalized.
func NewRay(o, d vector3) *ray {
	return &ray{o, d.Normalize()}
}

func (r *ray) Sample(scn *scene, texs map[string]image.Image) (clr Color, hit bool) {
	dist := float64(1000000000)
	for _, geo := range scn.geos {
		// ray not hit the bounding sphere means not hit the geometry.
		bs := geo.bb.BSphere()
		if !r.HitBSphere(bs) {
			continue
		}
		// inside bounding sphere. check more.
		for _, ply := range geo.plys {
			if !r.HitBSphere(ply.BBox().BSphere()) {
				continue
			}
			switch len(ply.vts) {
			case 3:
				p, u, v, ok := r.HitInfo(ply.vts[0].P, ply.vts[1].P, ply.vts[2].P)
				if !ok {
					continue
				}
				hit = true
				hitd := r.o.Sub(p).Len()
				if hitd < dist {
					dist = hitd
					clr = HitColor(r, ply, u, v, geo, scn.lits, texs)
				}
			case 4:
				// divide the square to 2 triangles. then we can use above (triangle) approach.
				// if ray hit any of them, ray hit the quad.
				p, u, v, ok := r.HitInfo(ply.vts[0].P, ply.vts[1].P, ply.vts[3].P)
				if !ok {
					p, u, v, ok = r.HitInfo(ply.vts[2].P, ply.vts[3].P, ply.vts[1].P)
					u, v = 1 - u, 1 - v
				}
				if !ok {
					continue
				}
				hit = true
				hitd := r.o.Sub(p).Len()
				if hitd < dist {
					dist = hitd
					clr = HitColor(r, ply, u, v, geo, scn.lits, texs)
				}
			default:
				panic("n-gon not supported yet.")
			}
		}
	}
	if !hit {
		return Color{}, false
	}
	return clr, true
}

// HitBSphere checks the ray hit the bounding sphere.
func (r *ray) HitBSphere(bs bsphere) bool {
		toBs := bs.o.Sub(r.o)
		dist := toBs.Sub(r.d.Mult(toBs.Dot(r.d))).Len()
		if dist > bs.r {
			return false
		}
		return true
}

// does the ray hit a-b-c polygon?
func (r *ray) HitInfo(a, b, c vector3) (p vector3, u, v float64, ok bool) {
	a = a.Sub(r.o)
	b = b.Sub(r.o)
	c = c.Sub(r.o)
	toB := b.Sub(a)
	toC := c.Sub(a)
	N := toB.Cross(toC).Normalize()
	dDotN := r.d.Dot(N)
	if math.Abs(dDotN) < 0.00001 {
		return
	}
	if dDotN < 0 {
		N, dDotN = N.Neg(), -dDotN
	}
	aDotN := a.Dot(N)
	if aDotN < 0 {
		return
	}
	dPly := r.d.Mult(aDotN/dDotN)
	dPlane := dPly.Sub(a)
	newC := toB.Cross(dPlane).Cross(toB).Normalize()
	dotC := dPlane.Dot(newC) / toC.Dot(newC)
	if dotC < 0 || dotC > 1 {
		return vector3{}, 0, 0, false
	}
	newB := toC.Cross(dPlane).Cross(toC).Normalize()
	dotB := dPlane.Dot(newB) / toB.Dot(newB)
	if dotB < 0 || dotB > 1 {
		return
	}
	if dotB + dotC > 1 {
		return
	}
	return dPly.Add(r.o), dotB, dotC, true
}

func HitColor(rr *ray, ply *polygon, u, v float64, geo *geometry, lits []*dirlight, texs map[string]image.Image) Color {
	var clr Color
	pth, ok := ply.sa["texture"]
	if !ok {
		pth, ok = geo.sa["texture"]
		if !ok {
			clr = Color{1, 1, 1, 1}
		}
	} else {
		tex, ok := texs[pth]
		if !ok {
			clr = Color{1, 1, 1, 1}
		} else {
			clr = TextureSample(tex, u, v)
		}
	}

	// TODO: Use intersect point's normal instead.
	N := ply.Normal()
	if N.Dot(rr.d) > 0 {
		N = N.Neg()
	}

	var r, g, b float64
	for _, l := range lits {
		dot := maxval(l.dir.Dot(N), 0)
		r += l.r * dot
		g += l.g * dot
		b += l.b * dot
	}

	return Color{clr.r * r, clr.g * g, clr.b * b, clr.a}
}

