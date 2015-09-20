package main

type scene struct {
	cam *camera
	geos []*geometry
	octs []*octree
	lits []*dirlight
}

func NewScene(cam *camera, geos []*geometry, lits []*dirlight) *scene {
	octs := make([]*octree, 0)
	for _, g := range geos {
		bb := g.BBox()
		octs = append(octs, ParseOctree(bb, g.plys))
	}
	return &scene{
		cam: cam,
		geos: geos,
		octs: octs,
		lits: lits,
	}
}
