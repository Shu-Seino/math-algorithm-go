package main
import (
	"fmt"
	"sort"
)
func Code_3_06_1() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	sort.Ints(A)
	fmt.Println(A)
}