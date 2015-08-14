package main

type scene struct {
	cam *camera
	geos []*geometry
	lits []*dirlight
}

func NewScene(cam *camera, geos []*geometry, lits []*dirlight) *scene {
	return &scene{
		cam: cam,
		geos: geos,
		lits: lits,
	}
}
