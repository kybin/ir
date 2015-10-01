package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func loadGeometry(file string) *geometry {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("cannot open file")
	}
	var f interface{}
	json.Unmarshal(b, &f)

	attrs := make(map[string][]interface{})
	attrSizes := make(map[string]int)
	attrTypes := make(map[string]string)
	pointAttrEl := Find(f, "attributes", "pointattributes").([]interface{})
	for i := range pointAttrEl {
		name := Find(pointAttrEl, i, 0, "name").(string)
		size := int(Find(pointAttrEl, i, 1, "size").(float64))
		typ := Find(pointAttrEl, i, 1, "values", "storage").(string)
		values := Find(pointAttrEl, i, 1, "values", "tuples").([]interface{})
		attrs[name] = values
		attrSizes[name] = size
		attrTypes[name] = typ
	}

	verts := make([]*vertex, 0)
	indiceSl := Find(f, "topology", "pointref", "indices").([]interface{})
	for i := range indiceSl {
		id := int(indiceSl[i].(float64))
		attrP := attrs["P"][id].([]interface{})
		P := vector3{
			attrP[0].(float64),
			attrP[1].(float64),
			attrP[2].(float64),
		}
		v := NewVertex(P)
		for name, values := range attrs {
			if name == "P" {
				continue
			}
			attr := values[id]
			typ := attrTypes[name]
			size := attrSizes[name]
			if typ == "fpreal32" || typ == "fpreal64" {
				if size == 1 {
					v.fa[name] = attr.(float64)
				} else if size == 2 {
					attrSlice := attr.([]interface{})
					x := attrSlice[0].(float64)
					y := attrSlice[1].(float64)
					v.v2a[name] = vector2{x, y}
				} else if size == 3 {
					attrSlice := attr.([]interface{})
					x := attrSlice[0].(float64)
					y := attrSlice[1].(float64)
					z := attrSlice[2].(float64)
					v.v3a[name] = vector3{x, y, z}
				} else {
					panic("cannot parse if not [1..3]float, yet.")
				}
			} else {
				// TODO: find out string identifier in geo format
				panic("cannot parse other than float type yet.")
			}
		}
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

