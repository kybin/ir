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

func main() {
	b, err := ioutil.ReadFile("box.geo")
	if err != nil {
		fmt.Println("cannot open file")
	}
	var f interface{}
	json.Unmarshal(b, &f)
	indiceSl := Find(f, "topology", "pointref", "indices").([]interface{})
	indices := make([]int, len(indiceSl))
	for i := range indiceSl {
		indices[i] = int(indiceSl[i].(float64))
	}
	fmt.Println(indices)
	pointAttrEl := Find(f, "attributes", "pointattributes").([]interface{})
	var poses interface{}
	for i := range pointAttrEl {
		if Find(pointAttrEl, i, 0, "name").(string) == "P" {
			poses = Find(pointAttrEl, i, 1, "values", "tuples")
			break
		}
	}
	fmt.Println(poses)
}
