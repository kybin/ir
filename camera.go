package main

import (
	"math"
)

// ref
// http://www.sidefx.com/docs/houdini14.0/ref/cameralenses

type camera struct {
	P vector3
	front, right, up vector3
	focal float64
	aptx float64 // aperture x
	resx, resy int
	near, far float64
}

func (c *camera) Apty() float64 {
	return float64(c.resy) / float64(c.resx) * c.aptx
}

func (c *camera) FOV() float64 {
	return 2 * math.Atan((c.Apty()/2) / c.focal)
}

