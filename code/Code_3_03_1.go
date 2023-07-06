package main

import "fmt"

func Code_3_03_1() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Scan(&A[i])
	}

	Answer := 0

	for i := 1; i <= N; i++ {
		for j := i + 1; j <= N; j++ {
			for k := j + 1; k <= N; k++ {
				for l := k + 1; l <= N; l++ {
					for m := l + 1; m <= N; m++ {
						if A[i]+A[j]+A[k]+A[l]+A[m] == 1000 {
							Answer++
						}
					}
				}
			}
		}
	}

	fmt.Println(Answer)
}
