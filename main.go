package main

import (
	"fmt"
	"math"
)

func main() {
	geo := loadGeometry()

	// TODO : rotation matrix
	// rotation by (x:30, y:30) degree.
	yRot := matrix4{
		math.Cos(pi / 6), 0, -math.Sin(pi / 6), 0,
		0, 1, 0, 0,
		math.Sin(pi / 6), 0, math.Cos(pi / 6), 0,
		0, 0, 0, 1,
	}
	xRot := matrix4{
		1, 0, 0, 0,
		0, math.Cos(pi / 6), math.Sin(pi / 6), 0,
		0, -math.Sin(pi / 6), math.Cos(pi / 6), 0,
		0, 0, 0, 1,
	}
	modelTransform := xRot.Multiply(yRot).Transpose()

	// view transform should inverse of camera transform
	// for easy reversing, think it as rotation+translation.
	// inverse of translation is just negate it.
	// inverse of rotation matrix is it's transpose.
	// http://www.katjaas.nl/transpose/transpose.html
	// assume camera axis are already normalized.
	//
	// TODO : redefine camera.
	cam := &camera{
		P:     vector3{0, 0, 10},
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
		1, 0, 0, -cam.P.x,
		0, 1, 0, -cam.P.y,
		0, 0, 1, -cam.P.z,
		0, 0, 0, 1,
	}
	viewRotation := matrix4{
		cam.right.x, cam.up.x, cam.front.x, 0,
		cam.right.y, cam.up.y, cam.front.y, 0,
		cam.right.z, cam.up.z, cam.front.z, 0,
		0, 0, 0, 1,
	}
	viewTransform := viewTranslation.Multiply(viewRotation)


	geo.Transform(viewTransform.Multiply(modelTransform))

	// TODO : do I need perspective projection in pbr renderer?
	// perspective projection
	// http://ogldev.atspace.co.uk/www/tutorial12/tutorial12.html
	// perspProjection := matrix4{
	//	1 / cam.Apty(), 0, 0, 0,
	//	0, 1 / cam.aptx, 0, 0,
	//	0, 0, (-cam.near - cam.far) / (cam.near - cam.far), 2 * cam.far * cam.near / (cam.near - cam.far),
	//	0, 0, 1, 0,
	//}

	render(cam, geo)
}

func loadGeometry() *geometry {
	top := NewPolygon(
		NewVertex(-1, 1, -1),
		NewVertex(-1, 1, 1),
		NewVertex(1, 1, 1),
		NewVertex(1, 1, -1),
	)
	front := NewPolygon(
		NewVertex(-1, -1, 1),
		NewVertex(-1, 1, 1),
		NewVertex(1, 1, 1),
		NewVertex(1, -1, 1),
	)
	left := NewPolygon(
		NewVertex(1, -1, -1),
		NewVertex(1, -1, 1),
		NewVertex(1, 1, 1),
		NewVertex(1, 1, -1),
	)
	right := NewPolygon(
		NewVertex(-1, -1, -1),
		NewVertex(-1, -1, 1),
		NewVertex(-1, 1, 1),
		NewVertex(-1, 1, -1),
	)
	back := NewPolygon(
		NewVertex(-1, -1, -1),
		NewVertex(-1, 1, -1),
		NewVertex(1, 1, -1),
		NewVertex(1, -1, -1),
	)
	bottom := NewPolygon(
		NewVertex(-1, -1, -1),
		NewVertex(-1, -1, 1),
		NewVertex(1, -1, 1),
		NewVertex(1, -1, -1),
	)

	return &geometry{
		back,
		left,
		front,
		right,
		bottom,
		top,
	}
}


func debug(c *camera, g *geometry, l *dirlight) {
	fmt.Println(*g)
	for _, p := range *g {
		fmt.Println(p.Normal())
	}
}
