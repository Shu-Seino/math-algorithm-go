package main

import (
	"fmt"
)

func Code_2_02_5() {
	var N, X int
	Answer := 0
	fmt.Scanf("%d %d", &N, &X)

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if i+j <= X {
				Answer += 1
			}
		}
	}

	fmt.Println(Answer)
}
