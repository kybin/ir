package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type vector3 struct {
	x, y, z float64
}

type Point struct {
	P vector3
}

type Vertex struct {
	pt *Point
}

type Prim struct {
	vts []*Vertex
}

type El map[string]interface{}

func NewEl(in interface{}) El {
	insl := in.([]interface{})
	el := make(El)
	for i := 0; i < len(insl); i += 2 {
		s := insl[i].(string)
		el[s] = insl[i+1]
	}
	return el
}

func Find(in interface{}, keys ...interface{}) interface{} {
	for _, k := range keys {
		i, ok := k.(int)
		if ok {
			in = in.([]interface{})[i]
			continue
		}
		s, ok := k.(string)
		if ok {
			in = NewEl(in)[s]
			continue
		}
		panic("unknown json key type.")
	}
	return in
}

func main() {
	b, err := ioutil.ReadFile("box.geo")
	if err != nil {
		fmt.Println("cannot open file")
	}
	var f interface{}
	json.Unmarshal(b, &f)

	points := make([]*Point, 0)
	pointAttrEl := Find(f, "attributes", "pointattributes").([]interface{})
	for i := range pointAttrEl {
		if Find(pointAttrEl, i, 0, "name").(string) == "P" {
			poses := Find(pointAttrEl, i, 1, "values", "tuples").([]interface{})
			for _, p := range poses {
				p0 := p.([]interface{})[0].(float64)
				p1 := p.([]interface{})[1].(float64)
				p2 := p.([]interface{})[2].(float64)
				pt := Point{vector3{p0, p1, p2}}
				points = append(points, &pt)
			}
			break
		}
	}
	fmt.Println(points)

	verts := make([]*Vertex, 0)
	indiceSl := Find(f, "topology", "pointref", "indices").([]interface{})
	for i := range indiceSl {
		ip := int(indiceSl[i].(float64))
		verts = append(verts, &Vertex{points[ip]})
	}
	fmt.Println(verts)

	prims := make([]*Prim, 0)
	primEl := Find(f, "primitives").([]interface{})
	for i := range primEl {
		if Find(primEl, i, 0, "type").(string) == "run" && Find(primEl, i, 0, "runtype").(string) == "Poly" {
			for _, v := range Find(primEl, i, 1).([]interface{}) {
				var prim Prim
				vec := v.([]interface{})[0].([]interface{})
				for _, i := range vec {
					prim.vts = append(prim.vts, verts[int(i.(float64))])
				}
				prims = append(prims, &prim)
			}
			break
		}
	}
	fmt.Println(prims)
}

