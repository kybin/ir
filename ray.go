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
	for i, oct := range scn.octs {
		leafs := r.HitOctreeLeafs(oct)
		// inside bounding sphere. check more.
		for _, oct := range leafs {
			for _, ply := range oct.polys {
				if !r.HitBSphere(ply.BBox().BSphere()) {
					continue
				}
				hitP, u, v, ok := r.HitPolyInfo(ply)
				if !ok {
					continue
				}
				hit = true
				hitd := r.o.Sub(hitP).Len()
				if hitd < dist {
					dist = hitd
					clr = HitColor(r, ply, u, v, scn.geos[i], scn.lits, texs)
				}
			}
		}
	}
	// TODO: return clr, hit
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

func (r *ray) HitBBox(bb bbox3) bool {
	left, right := bb.min.x, bb.max.x
	bottom, top := bb.min.y, bb.max.y
	back, front := bb.min.z, bb.max.z

	leftTopFront := vector3{left, top, front}
	leftTopBack := vector3{left, top, back}
	leftBottomFront := vector3{left, bottom, front}
	leftBottomBack := vector3{left, bottom, back}
	rightTopFront := vector3{right, top, front}
	rightTopBack := vector3{right, top, back}
	rightBottomFront := vector3{right, bottom, front}
	rightBottomBack := vector3{right, bottom, back}

	// TODO: make it clock / anti-clock wise?
	// topPlane
	if r.Hit(leftTopFront, leftTopBack, rightTopBack) {
		return true
	}
	if r.Hit(leftTopFront, rightTopBack, rightTopFront) {
		return true
	}
	// bottomPlane
	if r.Hit(leftBottomFront, leftBottomBack, rightBottomBack) {
		return true
	}
	if r.Hit(leftBottomFront, rightBottomBack, rightBottomFront) {
		return true
	}
	// leftPlane
	if r.Hit(leftBottomFront, leftBottomBack, leftTopBack) {
		return true
	}
	if r.Hit(leftBottomFront, leftTopBack, leftTopFront) {
		return true
	}
	// rightPlane
	if r.Hit(rightBottomFront, rightBottomBack, rightTopBack) {
		return true
	}
	if r.Hit(rightBottomFront, rightTopBack, rightTopFront) {
		return true
	}
	// frontPlane
	if r.Hit(leftBottomFront, leftTopFront, rightTopFront) {
		return true
	}
	if r.Hit(leftBottomFront, rightTopFront, rightBottomFront) {
		return true
	}
	// backPlane
	if r.Hit(leftBottomBack, leftTopBack, rightTopBack) {
		return true
	}
	if r.Hit(leftBottomBack, rightTopBack, rightBottomBack) {
		return true
	}

	return false
}

func (r *ray) HitOctreeLeafs(oct *octree) []*octree {
	leafs := make([]*octree, 0)
	if !r.HitBSphere(oct.bound.BSphere()) {
		return leafs
	}
	if !r.HitBBox(oct.bound) {
		return leafs
	}
	for _, child := range oct.children {
		if child == nil {
			continue
		}
		// hit the octree.
		if child.leaf {
			leafs = append(leafs, child)
		} else {
			childleafs := r.HitOctreeLeafs(child)
			leafs = append(leafs, childleafs...)
		}
	}
	return leafs
}

func (r *ray) HitPolyInfo(ply *polygon) (p vector3, u, v float64, ok bool) {
	switch len(ply.vts) {
	case 3:
		p, u, v, ok := r.HitInfo(ply.vts[0].P, ply.vts[1].P, ply.vts[2].P)
		return p, u, v, ok
	case 4:
		// divide the square to 2 triangles. then we can use above (triangle) approach.
		// if ray hit any of them, ray hit the quad.
		p, u, v, ok := r.HitInfo(ply.vts[0].P, ply.vts[1].P, ply.vts[3].P)
		if !ok {
			p, u, v, ok = r.HitInfo(ply.vts[2].P, ply.vts[3].P, ply.vts[1].P)
			u, v = 1 - u, 1 - v
		}
		return p, u, v, ok
	default:
		panic("n-gon not supported yet.")
	}
}

// TODO: make better version, not use HitInfo.
func (r *ray) Hit(a, b, c vector3) bool {
	_, _, _, ok := r.HitInfo(a, b, c)
	return ok
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

// TODO: if N does not exist.
func HitNormal(r *ray, ply *polygon, u, v float64) vector3 {
	switch len(ply.vts) {
	case 3:
		// find dist from verticies to u,v pos?
		// (sqrt(2)-distA)/sqrt(2)*attA + (1-distB)*attB + (1-distC)*attC
		panic("not implemented yet.")
	case 4:
		// blend x then y.
		Na := mixVector3(ply.vts[0].v3a["N"], ply.vts[1].v3a["N"], u)
		Nb := mixVector3(ply.vts[3].v3a["N"], ply.vts[2].v3a["N"], u)
		return mixVector3(Na, Nb, v).Normalize()
	default:
		panic("n-gon not supported yet.")
	}
}

func HitColor(rr *ray, ply *polygon, u, v float64, geo *geometry, lits []*dirlight, texs map[string]image.Image) Color {
	clr := Color{1, 1, 1, 1}

	texpath := ply.sa["texture"]
	if texpath == "" {
		texpath = geo.sa["texture"]
	}

	if texpath != "" {
		tex, ok := texs[texpath]
		if ok {
			clr = TextureSample(tex, u, v)
		}
	}

	N := HitNormal(rr, ply, u, v)

	var r, g, b float64
	for _, l := range lits {
		dot := maxval(l.dir.Neg().Dot(N), 0)
		r += l.r * dot
		g += l.g * dot
		b += l.b * dot
	}
	return Color{clr.r * r, clr.g * g, clr.b * b, clr.a}
}

