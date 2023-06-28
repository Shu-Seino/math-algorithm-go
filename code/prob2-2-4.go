package main

import "fmt"

func Prob2_2_4() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Scan(&A[i])
	}

	Answer := 0
	for i := 1; i <= N; i++ {
		Answer += A[i]
	}

	fmt.Println(Answer % 100)
}
