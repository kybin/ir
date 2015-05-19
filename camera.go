package main

import (
	"math"
)

// code here is shamelessly copied from
// http://www.sidefx.com/docs/houdini14.0/ref/cameralenses

type camera struct {
	focal float64
	apx float64
	resx, resy float64
	asp float64
}

func NewCamera(focal, apx, resx, resy, asp float64) *camera {
	return &camera{focal, apx, resx, resy, asp}
}

func (c *camera) Aperture() (float64, float64) {
	apy := (c.resy*c.apx) / (c.resx*c.asp)
	return c.apx, apy
}

func (c *camera) FOV() (float64, float64) {
	apx, apy := c.Aperture()
	fovx := 2 * math.Atan((apx/2) / c.focal)
	fovy := 2 * math.Atan((apy/2) / c.focal)
	return fovx, fovy
}
