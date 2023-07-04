package main

import (
	"fmt"
)

func Prob3_3_7() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	x := 0
	y := 0
	z := 0

	for i := 0; i < N; i++ {
		if A[i] == 1 {
			x++
		}
		if A[i] == 2 {
			y++
		}
		if A[i] == 3 {
			z++
		}
	}

	result := x*(x-1)/2 + y*(y-1)/2 + z*(z-1)/2
	fmt.Println(result)
}
