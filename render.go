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

	nsample := 1
	rng := rand1D()

	l := c.right.Mult(-c.aptx / 2)
	r := c.right.Mult(c.aptx / 2)
	t := c.up.Mult(c.Apty() / 2)
	b := c.up.Mult(-c.Apty() / 2)
	f := c.front.Mult(c.focal)

	img := image.NewRGBA(image.Rect(0, 0, c.resx, c.resy))
	for py := 0; py < c.resy; py++ {
		for px := 0; px < c.resx; px++ {
			clr := Color{}
			for i := 0; i < nsample; i++ {
				offx := <-rng
				offy := <-rng
				lr := mixVector3(l, r, (float64(px)+offx)/float64(c.resx-1))
				tb := mixVector3(t, b, (float64(py)+offy)/float64(c.resy-1))
				r := &ray{o:c.P, d:lr.Add(tb).Add(f).Normalize()}
				sc, _ := r.Sample(scn, texs)
				clr = clr.Add(sc)
			}
			clr = clr.Div(float64(nsample))
			go func(img *image.RGBA, px, py int, clr color.RGBA) {
				img.Set(px, py, clr)
			}(img, px, py, color.RGBA{uint8(255*clr.r), uint8(255*clr.g), uint8(255*clr.b), uint8(255*clr.a)})
		}
	}

	fd, err := os.Create("hello.png")
	if err != nil {
		panic("cannot generate output image file.")
	}
	defer fd.Close()
	err = png.Encode(fd, img)
	if err != nil {
		panic("cannot write image")
	}
}


