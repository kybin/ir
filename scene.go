package main

import(
	"image"
)

type scene struct {
	cam *camera
	geo *geometry // TODO geos []*geometry
	lit *dirlight // TODO lits []*light
	tex image.Image
}
