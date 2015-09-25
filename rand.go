package main

import (
	"math/rand"
)

func rand1D() chan float64 {
	rng := make(chan float64)
	go func() {
		for {
			rng <- rand.Float64()
		}
	}()
	return rng
}
