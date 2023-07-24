//select sort
package main
import (
	"fmt"
)

func Code_3_06_2() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	for i := 0; i < N; i++ {
		minj := i
		for j := i; j < N; j++ {
			if A[j] < A[minj] {
				minj = j
			}
		}
		A[i], A[minj] = A[minj], A[i]
	}

	fmt.Println(A)
}