package main

type bbox3 struct {
	min vector3
	max vector3
}

func (b bbox3) Add(b2 bbox3) bbox3 {
	var min, max vector3
	min.x = minval(b.min.x, b2.min.x)
	min.y = minval(b.min.y, b2.min.y)
	min.z = minval(b.min.z, b2.min.z)
	max.x = maxval(b.max.x, b2.max.x)
	max.y = maxval(b.max.y, b2.max.y)
	max.z = maxval(b.max.z, b2.max.z)
	return bbox3{min, max}
}

func (b bbox3) BSphere() bsphere {
	o := b.min.Add(b.max).Div(2)
	r := b.max.Sub(o).Len()
	return bsphere{o, r}
}

type bsphere struct {
	o vector3
	r float64
}

