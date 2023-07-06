package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	Answer := 0.0
	for i := 1; i <= N; i++ {
		Answer += 1.0 * float64(N) / float64(i)
	}

	fmt.Printf("%.12f\n", Answer)
}
