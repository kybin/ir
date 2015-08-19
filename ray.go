package main

import (
	"math"
	"image"
)

type ray struct {
	o vector3
	d vector3
}

func (r *ray) Sample(scn *scene, texs map[string]image.Image) (clr Color, hit bool) {
	dist := float64(1000000000)
	for _, geo := range scn.geos {
		for _, ply := range geo.plys {
			switch len(ply.vts) {
			case 3:
				p, u, v, ok := r.HitInfo(ply.vts[0].P, ply.vts[1].P, ply.vts[2].P)
				if !ok {
					continue
				}
				hit = true
				hitd := p.Len()
				if hitd < dist {
					dist = hitd
					clr = HitColor(ply, u, v, geo, scn.lits, texs)
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
				hitd := p.Len()
				if hitd < dist {
					dist = hitd
					clr = HitColor(ply, u, v, geo, scn.lits, texs)
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
		return vector3{}, 0, 0, false
	}
	if dDotN < 0 {
		N, dDotN = N.Neg(), -dDotN
	}
	aDotN := a.Dot(N)
	if aDotN < 0 {
		panic("aDot is smaller than 0.")
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
		return vector3{}, 0, 0, false
	}
	if dotB + dotC > 1 {
		return vector3{}, 0, 0, false
	}
	return dPly.Add(r.o), dotB, dotC, true
}

func HitColor(ply *polygon, u, v float64, geo *geometry, lits []*dirlight, texs map[string]image.Image) Color {
	pth, ok := ply.sa["texture"]
	if !ok {
		pth, ok = geo.sa["texture"]
		if !ok {
			return Color{1, 1, 1, 1}
		}
	}
	tex, ok := texs[pth]
	if !ok {
		return Color{1, 1, 1, 1}
	}
	clr := TextureSample(tex, u, v)

	var r, g, b float64
	for _, lit := range lits {
		dot := lit.dir.Dot(ply.Normal())
		r += lit.r * dot
		g += lit.g * dot
		b += lit.b * dot
	}

	return Color{float64(clr.R)/255 * r, float64(clr.G)/255 * g, float64(clr.B)/255 * b, float64(clr.A)/255}
}

