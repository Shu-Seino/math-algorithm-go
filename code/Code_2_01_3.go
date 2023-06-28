package main

import "fmt"

func Code_2_01_3() {
	var n int
	var Answer int

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		Answer += i
	}
	fmt.Println(Answer)
}
