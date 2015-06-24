package main

// quaternions are very useful for
// rotation one vector to another vector.
// a quaternion is consist of an axis and angle.
// we can write this term as `v + w`.
//
// video : https://www.youtube.com/watch?v=mHVwd8gYLnI

import (
	"math"
)

type quaternion struct {
	v vector3
	w float64
}

func (q quaternion) Len() float64 {
	return math.Sqrt(q.v.x*q.v.x + q.v.y*q.v.y + q.v.z*q.v.z + q.w*q.w)
}

func (q quaternion) Normalize() quaternion {
	l := q.Len()
	return quaternion{q.v.Div(l), q.w / l}
}

func (q quaternion) Neg() quaternion {
	return quaternion{q.v.Neg(), -q.w}
}

func (q quaternion) Dot(q2 quaternion) float64 {
	return q.v.Dot(q2.v) + q.w*q2.w
}

func (q quaternion) Add(q2 quaternion) quaternion {
	return quaternion{q.v.Add(q2.v), q.w + q2.w}
}

func (q quaternion) Sub(q2 quaternion) quaternion {
	return quaternion{q.v.Sub(q2.v), q.w - q2.w}
}

func (q quaternion) Mult(f float64) quaternion {
	return quaternion{q.v.Mult(f), q.w * f}
}

func (q quaternion) Div(f float64) quaternion {
	return quaternion{q.v.Div(f), q.w / f}
}

func Slerp(q, q2 quaternion, t float64) quaternion {
	cosTh := q.Dot(q2)
	if cosTh < 0 {
		q2 = q2.Neg()
		cosTh *= -1
	}
	if cosTh > 0.995 {
		// linear interpolation
		return q.Mult(1 - t).Add(q2.Mult(t)).Normalize()
	} else {
		animTh := math.Acos(cosTh) * t
		qPerpn := q2.Sub(q.Mult(cosTh))
		return q.Mult(math.Cos(animTh)).Add(qPerpn.Mult(animTh))
	}
}
