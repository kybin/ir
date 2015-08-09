package main

import (
	"math"
)

type ray struct {
	o vector3
	d vector3
}

func (r *ray) Sample(scn *scene) (sample float64, hit bool) {
	dist := float64(1000000000)
	for _, ply := range (*scn.geo) {
		switch len(ply.vts) {
		case 3:
			p, ok := r.HitP(ply.vts[0].P, ply.vts[1].P, ply.vts[2].P)
			if !ok {
				continue
			}
			hit = true
			hitd := p.Len()
			if hitd < dist {
				dist = hitd
				sample = scn.lit.dir.Dot(ply.Normal())
			}
		case 4:
			// divide the square to 2 triangles. then we can use above (triangle) approach.
			// if ray hit any of them, ray hit the quad.
			p, ok := r.HitP(ply.vts[0].P, ply.vts[1].P, ply.vts[2].P)
			if !ok {
				p, ok = r.HitP(ply.vts[0].P, ply.vts[2].P, ply.vts[3].P)
			}
			if !ok {
				continue
			}
			hit = true
			hitd := p.Len()
			if hitd < dist {
				dist = hitd
				sample = scn.lit.dir.Dot(ply.Normal())
			}
		default:
			panic("n-gon not supported yet.")
		}
	}
	if !hit {
		return 0, false
	}
	return sample, hit
}

// does the ray hit a-b-c polygon?
func (r *ray) HitP(a, b, c vector3) (vector3, bool) {
	a = a.Sub(r.o)
	b = b.Sub(r.o)
	c = c.Sub(r.o)
	toB := b.Sub(a)
	toC := c.Sub(a)
	N := toB.Cross(toC).Normalize()
	dDotN := r.d.Dot(N)
	if math.Abs(dDotN) < 0.00001 {
		return vector3{}, false
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
		return vector3{}, false
	}
	newB := toC.Cross(dPlane).Cross(toC).Normalize()
	dotB := dPlane.Dot(newB) / toB.Dot(newB)
	if dotB < 0 || dotB > 1 {
		return vector3{}, false
	}
	if dotB + dotC > 1 {
		return vector3{}, false
	}
	return dPly.Add(r.o), true
}

