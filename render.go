package main

import (
	"os"
	"image"
	"image/color"
	"image/png"
	"math/rand"
)

func render(c *camera, g *geometry) {
	// TODO : copy geometry?
	// TODO : clipping
	// flatten
	for _, p := range *g {
		for _, v := range p.vts {
			v.P = vector3{v.P.x/v.P.z, v.P.y/v.P.z, 1}
		}
	}
	dx := c.aptx / float64(c.resx)
	dy := c.Apty() / float64(c.resy)
	rnd := rand.New(rand.NewSource(99))
	nsample := 9
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
			var cl float64
			for i := 0; i < nsample; i++ {
				r := &ray{vector3{x + dx * rnd.Float64(), y + dy * rnd.Float64(), c.focal}}
				if r.Hit(g) {
					cl += 1
				}
			}
			cl /= float64(nsample)
			img.Set(px, py, color.RGBA{uint8(255*cl), uint8(255*cl), uint8(255*cl), 255})
		}
	}
	err = png.Encode(f, img)
	if err != nil {
		panic("cannot write image")
	}
}

