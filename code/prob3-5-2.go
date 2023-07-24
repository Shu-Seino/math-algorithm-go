package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Prob_3_5_2() {
	N := 1000000
	M := 0

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= N; i++ {
		px := 6.0 * rand.Float64()
		py := 9.0 * rand.Float64()

		// Distance from point (3, 3)
		dist_33 := math.Sqrt((px-3.0)*(px-3.0) + (py-3.0)*(py-3.0))

		// Distance from point (3, 7)
		dist_37 := math.Sqrt((px-3.0)*(px-3.0) + (py-7.0)*(py-7.0))

		// Condition check
		if dist_33 <= 3.0 || dist_37 <= 2.0 {
			M++
		}
	}

	// Print the number of points inside or on the boundary of the circles
	fmt.Println(M)
}
