package main

// quaternions are very useful for
// rotation one vector to another vector.
// a quaternion is consist of an axis and angle.
// we can write this term as `v + w`.
//
// video : https://www.youtube.com/watch?v=mHVwd8gYLnI

type struct quaternion {
	v vector3
	w float64
}

func (q quaternion) Len() float64 {
	return math.Sqrt(q.v.Dot(q.v) + q.w * q.w)
}

func (q quaternion) Normalize() quaternion {
	l = q.Len()
	return quaternion{q.v.DivF(l), q.w / l}
}

func (q quaternion) Add(q2 quaternion) quaternion {
	return quaternion{q.v.Add(q2.v), q.w + q2.w}
}

func (q quaternion) Sub(q2 quaternion) quaternion {
	return quaternion{q.v.sub(q2.v), q.w - q2.w}
}

func (q quaternion) Mult(q2 quaternion) quaternion {
	return quaternion{q.v.Mult(q2.v), q.w * q2.w}
}

func (q quaternion) MultF(f float64) quaternion {
	return quaternion{q.v.MultF(f), q.w * f}
}

func (q quaternion) Div(q2 quaternion) quaternion {
	return quaternion{q.v.Div(q2.v), q.w / q2.w}
}

func (q quaternion) DivF(f float64) quaternion {
	return quaternion{q.v.DivF(f), q.w / f}
}

func (q quaternion) Dot(q2 quaternion) float64 {
	return q.v.Dot(q2.v) + q.w * q2.w
}

func Slerp(q, q2 quaternion, t float64) quaternion {
	cosTheta := q.Dot(q2)
	if cosTheta < 0 {
		q2 = q2.Negate()
		cosTheta *= -1
	}
	if cosTheta > 0.995 {
		// linear interpolation
		return (q.MultF(1 - t).Add(q2.MultF(t))).Normalize()
	} else {
		theta := math.Acos(cosTheta)
		thetaP := theta * t
		qPerp := q2.Sub(q1.MultF(cosTheta))
		return q.Mult(math.Cos(thetap), qPerp.MultF(thetaP))
	}
}
