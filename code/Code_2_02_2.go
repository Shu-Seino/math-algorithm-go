package main

import (
	"fmt"
)

func Code_2_02_2() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println("AND", a&b)
	fmt.Println("OR", a|b)
	fmt.Println("XOR", a^b)
}
