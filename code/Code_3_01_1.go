package main

import (
	"fmt"
)

func isPrime2(N int64) bool {
	// N を 2 以上の整数とし、N が素数であれば true、素数でなければ false を返す
	for i := int64(2); i <= N-1; i++ {
		// N が i で割り切れた場合、この時点で素数ではないと分かる
		if N%i == 0 {
			return false
		}
	}
	return true
}

func Code_3_01_1() {
	N := int64(17)
	if isPrime2(N) {
		fmt.Printf("%d is a prime number\n", N)
	} else {
		fmt.Printf("%d is not a prime number\n", N)
	}
}
