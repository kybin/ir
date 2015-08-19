package main

import (
	"os"
	"image"
	"image/color"
	"fmt"
)

func LoadTexture(pth string) (image.Image, bool) {
	f, err := os.Open(pth)
	if err != nil {
		fmt.Printf("cannot open texture : %v", pth)
		return nil, false
	}
	defer f.Close()
	tex, _, err := image.Decode(f)
	if err != nil {
		fmt.Printf("cannot decode texture : %v", pth)
		return nil, false
	}
	return tex, true
}

func TextureSample(tex image.Image, u, v float64) color.RGBA {
	bb := tex.Bounds()
	x := int(float64(bb.Min.X) + u*float64(bb.Max.X-bb.Min.X))
	y := int(float64(bb.Min.Y) + v*float64(bb.Max.Y-bb.Min.Y))
	r, g, b, a := tex.At(x, y).RGBA()
	return color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
}

