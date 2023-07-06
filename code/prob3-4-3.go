package main

import (
	"fmt"
)

func Prob3_4_3() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	for i := 0; i < N; i++ {
		fmt.Scan(&B[i])
	}

	Answer := 0.0
	for i := 0; i < N; i++ {
		Answer += (1.0 / 3.0) * float64(A[i]) + (2.0 / 3.0) * float64(B[i])
	}

	fmt.Printf("%.12f\n", Answer)
}
