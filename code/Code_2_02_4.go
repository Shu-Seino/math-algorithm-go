package main

import (
	"fmt"
)

func Code_2_02_4() {
	var N, X, Y int
    fmt.Scanf("%d %d %d", &N, &X, &Y)

	cnt := 0
    for i := 1; i <= N; i++ {
        if i%X == 0 || i%Y == 0 {
            cnt++
        }
    }

    fmt.Println(cnt)
}