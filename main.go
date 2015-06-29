package main

import (
	"fmt"
	"math"
)

func main() {
	geo := loadGeometry()
	modelTransform := matrix4{
		math.Cos(pi / 4), 0, -math.Sin(pi / 4), 0,
		0, 1, 0, 0,
		math.Sin(pi / 4), 0, math.Cos(pi / 4), 0,
		1, -1, -4, 1,
	}

	// view transform should inverse of camera transform
	// for easy reversing, think it as rotation+translation.
	// inverse of translation is just negate it.
	// inverse of rotation matrix is it's transpose.
	// http://www.katjaas.nl/transpose/transpose.html
	// assume camera axis are already normalized.
	//
	// TODO : redefine camera.
	cam := &camera{
		P:     vector3{0, 0, 5},
		front: vector3{0, 0, -1},
		right: vector3{1, 0, 0},
		up:    vector3{0, 1, 0},
		focal: 50,
		aptx:  41.4214,
		resx:  1920,
		resy:  1080,
		near:  0.001,
		far:   10000,
	}
	viewTranslation := matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		-cam.P.x, -cam.P.y, -cam.P.z, 1,
	}
	viewRotation := matrix4{
		cam.right.x, cam.up.x, cam.front.x, 0,
		cam.right.y, cam.up.y, cam.front.y, 0,
		cam.right.z, cam.up.z, cam.front.z, 0,
		0, 0, 0, 1,
	}
	viewTransform := viewTranslation.Multiply(viewRotation)

	// perspective projection
	// http://ogldev.atspace.co.uk/www/tutorial12/tutorial12.html
	perspProjection := matrix4{
		1 / cam.Apty(), 0, 0, 0,
		0, 1 / cam.aptx, 0, 0,
		0, 0, (-cam.near - cam.far) / (cam.near - cam.far), 2 * cam.far * cam.near / (cam.near - cam.far),
		0, 0, 1, 0,
	}

	// TODO : to ndc?

	geo.Transform(perspProjection.Multiply(viewTransform.Multiply(modelTransform)))

	render(geo)
}

func loadGeometry() geometry {
	top := polygon{
		NewVertex(-1, 1, -1),
		NewVertex(-1, 1, 1),
		NewVertex(1, 1, 1),
		NewVertex(1, 1, -1),
	}
	front := polygon{
		NewVertex(-1, -1, 1),
		NewVertex(-1, 1, 1),
		NewVertex(1, 1, 1),
		NewVertex(1, -1, 1),
	}
	left := polygon{
		NewVertex(1, -1, -1),
		NewVertex(1, -1, 1),
		NewVertex(1, 1, 1),
		NewVertex(1, 1, -1),
	}
	right := polygon{
		NewVertex(-1, -1, -1),
		NewVertex(-1, -1, 1),
		NewVertex(-1, 1, 1),
		NewVertex(-1, 1, -1),
	}
	back := polygon{
		NewVertex(-1, -1, -1),
		NewVertex(-1, 1, -1),
		NewVertex(1, 1, -1),
		NewVertex(1, -1, -1),
	}
	bottom := polygon{
		NewVertex(-1, -1, -1),
		NewVertex(-1, -1, 1),
		NewVertex(1, -1, 1),
		NewVertex(1, -1, -1),
	}

	return geometry{
		&top,
		&front,
		&left,
		&right,
		&back,
		&bottom,
	}
}

func render(g geometry) {
	fmt.Println(g)
	for _, p := range g {
		for _, v := range *p {
			fmt.Println(*v)
		}
	}
}
