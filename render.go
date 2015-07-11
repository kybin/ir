package main

import (
	"os"
	"image"
	"image/color"
	"image/png"
)

func render(c *camera, g *geometry) {
	// TODO : copy geometry?
	// TODO : clipping
	// flatten
	for _, p := range *g {
		for _, v := range *p {
			*v = vertex{v.x / v.z, v.y / v.z, 1, 1}
		}
	}
	f, err := os.Create("hello.png")
	if err != nil {
		panic("cannot generate image file.")
	}
	defer f.Close()
	img := image.NewRGBA(image.Rect(0, 0, c.resx, c.resy))
	for py := 0; py < c.resy; py++ {
		for px := 0; px < c.resx; px++ {
			x := mix(-c.aptx/2, c.aptx/2, float64(px)/float64(c.resx))
			y := mix(c.Apty()/2, -c.Apty()/2, float64(py)/float64(c.resy))
			r := &ray{vector3{x, y, c.focal}}
			if r.Hit(g) {
				img.Set(px, py, color.RGBA{255, 255, 255, 255})
			} else {
				img.Set(px, py, color.RGBA{0, 0, 0, 255})
			}
		}
	}
	err = png.Encode(f, img)
	if err != nil {
		panic("cannot write image")
	}
}

