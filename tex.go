package main

import (
	"fmt"
	"image"
	"os"
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

func TextureSample(tex image.Image, u, v float64) Color {
	bb := tex.Bounds()
	x := int(float64(bb.Min.X) + u*float64(bb.Max.X-bb.Min.X-1))
	y := int(float64(bb.Min.Y) + v*float64(bb.Max.Y-bb.Min.Y-1))
	r, g, b, a := tex.At(x, y).RGBA()
	return Color{float64(r>>8) / 255, float64(g>>8) / 255, float64(b>>8) / 255, float64(a>>8) / 255}
}
