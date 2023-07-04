package main

import "fmt"

func Code_3_01_3() {
	var N int64
	fmt.Scan(&N)
	for i := int64(1); i*i <= N; i++ {
		if N%i != 0 {
			continue
		}
		fmt.Println(i) // i を約数に追加
		if i != N/i {
			fmt.Println(N / i) // i ≠ N/i のとき、N/i も約数に追加
		}
	}
}
