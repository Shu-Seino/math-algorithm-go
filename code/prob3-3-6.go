package main

import (
	"fmt"
)

func Prob3_3_6() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	a := 0
	b := 0
	c := 0
	d := 0
	for i := 0; i < N; i++ {
		if A[i] == 100 {
			a++
		}
		if A[i] == 200 {
			b++
		}
		if A[i] == 300 {
			c++
		}
		if A[i] == 400 {
			d++
		}
	}

	result := a*d + b*c
	fmt.Println(result)
}
