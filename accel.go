package main

type octree struct {
	bound    bbox
	leaf     bool
	polys    []*polygon
	children [8]*octree
}

// TODO:
//   sometimes it iterates infinitely which cause terminates ir.
//   expecially when I set n-poly for the leaf is too small. fix it.

func ParseOctree(bb bbox, polys []*polygon) *octree {
	if len(polys) == 0 {
		return nil
	}

	// ray intersect check with bounding box need 64 intersect checking.
	// it the octree has polygons less than 64, make it `leaf` and don't split it more.
	if len(polys) <= 32 {
		var emptyChildren [8]*octree
		return &octree{
			bound:    bb,
			leaf:     true,
			polys:    polys,
			children: emptyChildren,
		}
	}

	center := bb.min.Add(bb.max).Div(2)

	var spaces [8]bbox
	// left, up, front
	spaces[0] = bbox{min: vector3{center.x, center.y, center.z}, max: vector3{bb.max.x, bb.max.y, bb.max.z}}
	// left, up, back
	spaces[1] = bbox{min: vector3{center.x, center.y, bb.min.z}, max: vector3{bb.max.x, bb.max.y, center.z}}
	// left, down, front
	spaces[2] = bbox{min: vector3{center.x, bb.min.y, center.z}, max: vector3{bb.max.x, center.y, bb.max.z}}
	// left, down, back
	spaces[3] = bbox{min: vector3{center.x, bb.min.y, bb.min.z}, max: vector3{bb.max.x, center.y, center.z}}
	// right, up, front
	spaces[4] = bbox{min: vector3{bb.min.x, center.y, center.z}, max: vector3{center.x, bb.max.y, bb.max.z}}
	// right, up, back
	spaces[5] = bbox{min: vector3{bb.min.x, center.y, bb.min.z}, max: vector3{center.x, bb.max.y, center.z}}
	// right, down, front
	spaces[6] = bbox{min: vector3{bb.min.x, bb.min.y, center.z}, max: vector3{center.x, center.y, bb.max.z}}
	// right, down, back
	spaces[7] = bbox{min: vector3{bb.min.x, bb.min.y, bb.min.z}, max: vector3{center.x, center.y, center.z}}

	// a polygon could live in some part of spaces (at least 1 part). check it.
	var childPolys [8][]*polygon
	for _, ply := range polys {
		// these variables indicate a polygon is inside where.
		left, right := false, false
		up, down := false, false
		front, back := false, false

		for _, pt := range ply.Points() {
			if pt.P.x >= center.x {
				left = true
			}
			if pt.P.x <= center.x {
				right = true
			}
			if pt.P.y >= center.y {
				up = true
			}
			if pt.P.y <= center.y {
				down = true
			}
			if pt.P.z >= center.z {
				front = true
			}
			if pt.P.z <= center.z {
				back = true
			}
		}

		if left && up && front {
			childPolys[0] = append(childPolys[0], ply)
		}
		if left && up && back {
			childPolys[1] = append(childPolys[1], ply)
		}
		if left && down && front {
			childPolys[2] = append(childPolys[2], ply)
		}
		if left && down && back {
			childPolys[3] = append(childPolys[3], ply)
		}
		if right && up && front {
			childPolys[4] = append(childPolys[4], ply)
		}
		if right && up && back {
			childPolys[5] = append(childPolys[5], ply)
		}
		if right && down && front {
			childPolys[6] = append(childPolys[6], ply)
		}
		if right && down && back {
			childPolys[7] = append(childPolys[7], ply)
		}
	}

	var children [8]*octree
	oct := &octree{
		bound:    bb,
		leaf:     false,
		polys:    nil,
		children: children,
	}
	for i := 0; i < 8; i++ {
		oct.children[i] = ParseOctree(spaces[i], childPolys[i])
	}
	return oct
}
