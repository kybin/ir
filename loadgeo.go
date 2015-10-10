package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// a key should int or string.
func Find(in interface{}, keys ...interface{}) interface{} {
	for _, k := range keys {
		insl := in.([]interface{})

		i, ok := k.(int)
		if ok {
			in = insl[i]
			continue
		}

		s, ok := k.(string)
		if ok {
			found := false
			for i := 0; i < len(insl); i += 2 {
				if insl[i].(string) != s {
					continue
				}
				found = true
				in = insl[i+1]
			}
			if !found {
				panic(fmt.Sprintf("key '%s' not found", s))
			}
			continue
		}

		panic("unknown key type.")
	}

	return in
}

type attr struct {
	size   int
	typ    string
	values []interface{}
}
// TODO: create attr methods.

func attributes(rootEl interface{}, key string) map[string]attr {
	attrs := make(map[string]attr)
	el, ok := Find(rootEl, "attributes", key).([]interface{})
	if !ok {
		return attrs
	}
	for i := range el {
		name := Find(el, i, 0, "name").(string)
		size := int(Find(el, i, 1, "size").(float64))
		if size == 0 {
			continue
		}
		typ := Find(el, i, 1, "values", "storage").(string)
		values := Find(el, i, 1, "values", "tuples").([]interface{})
		attrs[name] = attr{size, typ, values}

	}
	return attrs
}

func setAttribute(ar Attributer, attrs map[string]attr, i int) {
	for name, attr := range attrs {
		val := attr.values[i]
		switch attr.typ {
		case "fpreal32", "fpreal64":
			switch attr.size {
			case 1:
				ar.SetFloatAttr(name, val.(float64))
			case 2:
				vals := val.([]interface{})
				ar.SetVectorAttr(name, vals[0].(float64), vals[1].(float64))
			case 3:
				vals := val.([]interface{})
				ar.SetVectorAttr(name, vals[0].(float64), vals[1].(float64), vals[2].(float64))
			default:
				panic("cannot parse the floats, yet.")
			}
		default:
			// TODO: find out string identifier in geo format
			panic("cannot parse other than float type yet.")
		}
	}
}

func loadGeometry(file string) *geometry {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("cannot open file")
	}
	var f interface{}
	json.Unmarshal(b, &f)

	npts := int(Find(f, "pointcount").(float64))
	pts := make([]*point, 0)
	ptAttrs := attributes(f, "pointattributes")
	ptPoses := ptAttrs["P"].values
	delete(ptAttrs, "P")
	for i := 0; i < npts; i++ {
		Pv := ptPoses[i].([]interface{})
		pt := NewPoint(vector3{
			Pv[0].(float64),
			Pv[1].(float64),
			Pv[2].(float64),
		})
		setAttribute(pt, ptAttrs, i)
		pts = append(pts, pt)
	}

	nvts := int(Find(f, "vertexcount").(float64))
	vts := make([]*vertex, 0)
	vtAttrs := attributes(f, "vertexattributes")
	indiceEl := Find(f, "topology", "pointref", "indices").([]interface{})
	for i := 0; i < nvts; i++ {
		ptid := int(indiceEl[i].(float64))
		vt := NewVertex(ptid)
		setAttribute(vt, vtAttrs, i)
		vts = append(vts, vt)
	}

	// TODO: parse all primitive types, not only polygon

	// polygon
	// TODO: set primitive attributes
	// npolys := int(Find(f, "primitivecount").(float64))
	polys := make([]*polygon, 0)
	primEl := Find(f, "primitives").([]interface{})
	for i := range primEl {
		if Find(primEl, i, 0, "type").(string) == "run" && Find(primEl, i, 0, "runtype").(string) == "Poly" {
			for _, v := range Find(primEl, i, 1).([]interface{}) {
				// vertexs of polygon
				vtids := make([]int, 0)
				idEls := v.([]interface{})[0].([]interface{})
				for _, idEl := range idEls {
					id := int(idEl.(float64))
					vtids = append(vtids, id)
				}
				ply := NewPolygon(vtids)
				polys = append(polys, ply)
			}
			break
		}
	}
	return NewGeometry(pts, vts, polys)
}

