package main

import (
	"fmt"
)

func Code_2_02_6() {
	var N, S int
	fmt.Scanf("%d %d", &N, &S)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])

	}

	answer := "No"
	for i := 0; i < (1 << N); i++ {
		partsum := 0
		for j := 0; j < N; j++ {
			if (i & (1 << j)) != 0 {
				partsum += A[j]
			}
		}
		if partsum == S {
			answer = "Yes"
			break
		}
	}

	fmt.Println(answer)
}
