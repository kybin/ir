package main

import (
	"fmt"
)

func main() {
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

	box := geometry{&top, &front, &left, &right, &back, &bottom}
	fmt.Println(box)
}
