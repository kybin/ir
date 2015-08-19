package main

import (
	"os"
	"image"
	"image/color"
	"image/png"
	_ "image/jpeg"
)

func render(scn *scene, texs map[string]image.Image) {
	// TODO : copy geometry?
	// TODO : clipping
	c := scn.cam
	f, err := os.Create("hello.png")
	if err != nil {
		panic("cannot generate output image file.")
	}
	defer f.Close()

	nsample := 1
	rng := rand1D()
	dx, dy := c.aptx/float64(c.resx), c.Apty()/float64(c.resy)

	img := image.NewRGBA(image.Rect(0, 0, c.resx, c.resy))
	for py := 0; py < c.resy; py++ {
		for px := 0; px < c.resx; px++ {
			clr := Color{}
			for i := 0; i < nsample; i++ {
				x := mix(-c.aptx/2, c.aptx/2, float64(px)/float64(c.resx))
				y := mix(c.Apty()/2, -c.Apty()/2, float64(py)/float64(c.resy))
				offx := dx * <-rng
				offy := dy * <-rng
				r := &ray{d:vector3{x+offx, y+offy, -c.focal}}
				c, _ := r.Sample(scn, texs)
				clr = clr.Add(c)
			}
			clr = clr.Div(float64(nsample))
			go func(img *image.RGBA, px, py int, clr color.RGBA) {
				img.Set(px, py, clr)
			}(img, px, py, color.RGBA{uint8(255*clr.r), uint8(255*clr.g), uint8(255*clr.b), uint8(255*clr.a)})
		}
	}
	err = png.Encode(f, img)
	if err != nil {
		panic("cannot write image")
	}
}


