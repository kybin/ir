package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

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

func loadGeometry() *geometry {
	b, err := ioutil.ReadFile("geo/rubbertoy.geo")
	if err != nil {
		fmt.Println("cannot open file")
	}
	var f interface{}
	json.Unmarshal(b, &f)

	poses := make([]vector3, 0)
	pointAttrEl := Find(f, "attributes", "pointattributes").([]interface{})
	for i := range pointAttrEl {
		if Find(pointAttrEl, i, 0, "name").(string) == "P" {
			poseEl := Find(pointAttrEl, i, 1, "values", "tuples").([]interface{})
			for _, p := range poseEl {
				p0 := p.([]interface{})[0].(float64)
				p1 := p.([]interface{})[1].(float64)
				p2 := p.([]interface{})[2].(float64)
				pos := vector3{p0, p1, p2}
				poses = append(poses, pos)
			}
			break
		}
	}

	verts := make([]*vertex, 0)
	indiceSl := Find(f, "topology", "pointref", "indices").([]interface{})
	for i := range indiceSl {
		ip := int(indiceSl[i].(float64))
		v := NewVertex(poses[ip])
		verts = append(verts, v)
	}

	polys := make([]*polygon, 0)
	primEl := Find(f, "primitives").([]interface{})
	for i := range primEl {
		if Find(primEl, i, 0, "type").(string) == "run" && Find(primEl, i, 0, "runtype").(string) == "Poly" {
			for _, v := range Find(primEl, i, 1).([]interface{}) {
				vts := make([]*vertex, 0)
				vec := v.([]interface{})[0].([]interface{})
				for _, i := range vec {
					vts = append(vts, verts[int(i.(float64))])
				}
				ply := NewPolygon(vts...)
				polys = append(polys, ply)
			}
			break
		}
	}
	return NewGeometry(polys...)
}

