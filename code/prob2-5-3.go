package main

import "fmt"

func Prob2_5_3() {
	var N int
	fmt.Scan(&N)
	Answer := 1

	for i := 2 ; i <= N; i++{
		Answer *= i
	}
	

	fmt.Println(Answer)
}