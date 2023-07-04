package main

import "fmt"

func Code_3_02_2() {
	var A, B int64
	fmt.Scan(&A, &B)
	fmt.Println(gcd2(A, B))
}

func gcd2(A, B int64) int64 {
	// 正の整数 A と B の最大公約数を返す関数
	// GCD は Greatest Common Divisor（最大公約数）の略
	for A >= 1 && B >= 1 {
		if A < B {
			B %= A // A < B の場合、大きい方 B を書き換える
		} else {
			A %= B // A >= B の場合、大きい方 A を書き換える
		}
	}
	if A >= 1 {
		return A
	}
	return B
}
