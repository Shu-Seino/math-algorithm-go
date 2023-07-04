package main

import "fmt"

func Prob3_1_2() {
	// 入力
	// var N int64
	// fmt.Scan(&N)
	N := int64(100)
	// 素因数分解（空白区切りで出力）
	flag := false
	for i := int64(2); i*i <= N; i++ {
		for N%i == 0 {
			if flag {
				fmt.Print(" ")
			}
			flag = true
			N /= i
			fmt.Print(i)
		}
	}
	if N >= 2 {
		if flag {
			fmt.Print(" ")
		}
		flag = true
		fmt.Print(N)
	}
	fmt.Println()
}
