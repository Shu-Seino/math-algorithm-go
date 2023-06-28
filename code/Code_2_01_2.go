package main

import "fmt"

func Code_2_01_2() {
	var n [3]int
	fmt.Scan(&n[0], &n[1], &n[2])
	sum := n[0] + n[1] + n[2]
	fmt.Println(sum)
}
