package main

import (
	"fmt"
)

func isPrime3(N int64) bool {
	// N を 2 以上の整数とし、N が素数であれば true、素数でなければ false を返す
	for i := int64(2); i*i <= N; i++ {
		if N%i == 0 {
			return false
		}
	}
	return true
}

func Code_3_01_2() {
	N := int64(17)
	if isPrime3(N) {
		fmt.Printf("%d is a prime number\n", N)
	} else {
		fmt.Printf("%d is not a prime number\n", N)
	}
}
