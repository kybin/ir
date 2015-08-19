package main

type Color struct {
	r, g, b, a float64
}

func (c Color) Add(c2 Color) Color {
	return Color{c.r+c2.r, c.g+c2.g, c.b+c2.b, c.a+c2.a}
}

func (c Color) Div(f float64) Color {
	return Color{c.r/f, c.g/f, c.b/f, c.a/f}
}

