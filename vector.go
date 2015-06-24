package main

import (
	"math"
)

type vector3 struct {
	x, y, z float64
}

// Eqaul means almost equal.
func (v vector3) Equal(v2 vector3) bool {
	igx := ignorable > math.Abs(v.x-v2.x)
	igy := ignorable > math.Abs(v.y-v2.y)
	igz := ignorable > math.Abs(v.z-v2.z)
	if igx && igy && igz {
		return true
	}
	return false
}

func (v vector3) Len() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v vector3) Normalize() vector3 {
	l := v.Len()
	return vector3{v.x / l, v.y / l, v.z / l}
}

func (v vector3) Neg() vector3 {
	return vector3{-v.x, -v.y, -v.z}
}

func (v vector3) Dot(v2 vector3) float64 {
	return v.x*v2.x + v.y*v2.y + v.z*v2.z
}

func (v vector3) Cross(v2 vector3) vector3 {
	return vector3{
		v.y*v2.z - v2.y*v.z,
		v.z*v2.x - v2.z*v.x,
		v.x*v2.y - v2.x*v.y,
	}
}

func (v vector3) Add(v2 vector3) vector3 {
	return vector3{v.x + v2.x, v.y + v2.y, v.z + v2.z}
}

func (v vector3) Sub(v2 vector3) vector3 {
	return vector3{v.x - v2.x, v.y - v2.y, v.z - v2.z}
}

func (v vector3) Mult(f float64) vector3 {
	return vector3{v.x * f, v.y * f, v.z * f}
}

func (v vector3) Div(f float64) vector3 {
	return vector3{v.x / f, v.y / f, v.z / f}
}

func (v vector3) Rotate(rx, ry, rz float64, unit, order string) vector3 {
	switch unit {
	case "radian":
	case "angle":
		rx = rx / 180 * pi
		ry = ry / 180 * pi
		rz = rz / 180 * pi
	default:
		panic("unknown unit")
	}
	for _, ord := range order {
		switch ord {
		case 'x':
			v = v.RotateX(rx)
		case 'y':
			v = v.RotateY(ry)
		case 'z':
			v = v.RotateZ(rz)
		default:
			panic("unknown order")
		}
	}
	return v
}

func (v vector3) RotateX(theta float64) vector3 {
	return vector3{
		v.x,
		math.Cos(theta)*v.y - math.Sin(theta)*v.z,
		math.Cos(theta)*v.z + math.Sin(theta)*v.y,
	}
}

func (v vector3) RotateY(theta float64) vector3 {
	return vector3{
		math.Cos(theta)*v.x + math.Sin(theta)*v.z,
		v.y,
		math.Cos(theta)*v.z - math.Sin(theta)*v.x,
	}
}

func (v vector3) RotateZ(theta float64) vector3 {
	return vector3{
		math.Cos(theta)*v.x - math.Sin(theta)*v.y,
		math.Cos(theta)*v.y + math.Sin(theta)*v.x,
		v.z,
	}
}

type vector4 struct {
	x, y, z, w float64
}
