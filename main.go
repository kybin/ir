package main

import (
	"fmt"
)

func main() {
	geo := loadGeometry()
	geo.Transform(IdentityMatrix())
	render(geo)

}

func loadGeometry() geometry {
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

	return geometry{
			&top,
			&front,
			&left,
			&right,
			&back,
			&bottom,
	}
}

func render(g geometry) {
	fmt.Println(g)
}
