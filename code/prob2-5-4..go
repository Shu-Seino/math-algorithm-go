package main

import "fmt"

func isPrime(N int) bool {
	for i := 2; i < N; i++ {
		if N%i == 0 {
			return false
		}
	}
	return true
}

func Prob2_5_4() {
	var N int
	fmt.Scan(&N)

	A := []int{}
	for i := 2; i <= N; i++ {
		if isPrime(i) {
			A = append(A, i)
		}
	}
	fmt.Println(A)
}
