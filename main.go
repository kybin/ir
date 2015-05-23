package main

import (
	"fmt"
	"math"
)

func main() {
	geo := loadGeometry()
	geo.Transform(matrix4{
						math.Cos(pi/4), 0, -math.Sin(pi/4), 0,
						0, 1, 0, 0,
						math.Sin(pi/4), 0, math.Cos(pi/4), 0,
						1, -1, -4, 1,
	})
	render(geo)
}

func loadGeometry() *geometry {
	top := polygon{
			&vertex{-1, 1, -1},
			&vertex{-1, 1, 1},
			&vertex{1, 1, 1},
			&vertex{1, 1, -1},
	}
	front := polygon{
			&vertex{-1, -1, 1},
			&vertex{-1, 1, 1},
			&vertex{1, 1, 1},
			&vertex{1, -1, 1},
	}
	left := polygon{
			&vertex{1, -1, -1},
			&vertex{1, -1, 1},
			&vertex{1, 1, 1},
			&vertex{1, 1, -1},
	}
	right := polygon{
			&vertex{-1, -1, -1},
			&vertex{-1, -1, 1},
			&vertex{-1, 1, 1},
			&vertex{-1, 1, -1},
	}
	back := polygon{
			&vertex{-1, -1, -1},
			&vertex{-1, 1, -1},
			&vertex{1, 1, -1},
			&vertex{1, -1, -1},
	}
	bottom := polygon{
			&vertex{-1, -1, -1},
			&vertex{-1, -1, 1},
			&vertex{1, -1, 1},
			&vertex{1, -1, -1},
	}

	return &geometry{
			&top,
			&front,
			&left,
			&right,
			&back,
			&bottom,
	}
}

func render(g *geometry) {
	fmt.Println(g)
}
