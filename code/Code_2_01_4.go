package main

import "fmt"

func Code_2_01_4() {
	var n int
	var Answer string

	fmt.Scan(&n)
	for i := 0; n >= 1; i++ {
		if n%2 == 0 {
			Answer = "0" + Answer
		} else if n%2 == 1 {
			Answer = "1" + Answer
		}
		n = n / 2
	}

	fmt.Println(Answer)

}
