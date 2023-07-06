package main

import (
	"fmt"
)

func Code_3_04_2() {
	var N int
	fmt.Scan(&N)

	P := make([]int, N)
	Q := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&P[i], &Q[i])
	}

	answer := 0.0
	for i := 0; i < N; i++ {
		answer += float64(Q[i]) / float64(P[i])
	}

	fmt.Printf("%.12f\n", answer)
}
