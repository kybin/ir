package main

import (
	"image"
	"math"
)

func main() {
	geo := loadGeometry("geo/rubbertoy.geo")

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
	geo.Transform(modelTransform)

	// view transform should inverse of camera transform
	// for easy reversing, think it as rotation+translation.
	// inverse of translation is just negate it.
	// inverse of rotation matrix is it's transpose.
	// http://www.katjaas.nl/transpose/transpose.html
	// assume camera axis are already normalized.
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

	lit := &dirlight{r: 1, g: 1, b: 1, dir: vector3{-0.5, -1, 0}.Normalize()}

	scn := NewScene(cam, []*geometry{geo}, []*dirlight{lit})

	texs := loadTextures([]string{"tex/uv.jpg", "tex/uv_gray.jpg"})

	render(scn, texs)
}

func loadTextures(pths []string) map[string]image.Image {
	texs := make(map[string]image.Image)
	for _, p := range pths {
		t, ok := LoadTexture(p)
		if !ok {
			continue
		}
		texs[p] = t
	}
	return texs
}

