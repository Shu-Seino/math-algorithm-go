package main

import (
	"fmt"
)

func Code_3_03_4() {
	var N int
	fmt.Scan(&N)

	B := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&B[i])
	}

	R := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&R[i])
	}

	blue := 0.0
	red := 0.0
	for i := 0; i < N; i++ {
		blue += float64(B[i]) / float64(N)
		red += float64(R[i]) / float64(N)
	}

	result := blue + red
	fmt.Printf("%.12f\n", result)
}
