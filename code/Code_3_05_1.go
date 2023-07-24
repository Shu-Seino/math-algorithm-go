package main

import (
	"fmt"
	"math/rand"
)

func Code_3_05_1() {
	N := 100 // N is the number of trials (change as needed)
	M := 0

	for i := 1; i <= N; i++ {
		px := rand.Float64() // Generate a random number between 0 and 1
		py := rand.Float64() // Generate a random number between 0 and 1

		// The distance from the origin is sqrt(px * px + py * py)
		// If this value is less than or equal to 1, it's within the circle
		if px*px+py*py <= 1.0 {
			M += 1
		}
	}

	fmt.Printf("%.12f\n", 4.0*float64(M)/float64(N))
}
