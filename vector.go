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

func (v vector3) Translate(v2 vector3) vector3 {
	return vector3{
		v.x + v2.x,
		v.y + v2.y,
		v.z + v2.z,
	}
}

func (v vector3) Rotate(rx, ry, rz float64, unit, order string) vector3 {
	switch unit {
		case "radian":
		case "angle":
			rx = rx/180*pi
			ry = ry/180*pi
			rz = rz/180*pi
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
