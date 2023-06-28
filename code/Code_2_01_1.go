package main

import "fmt"

func Code_2_01_1() {
	var num int
	fmt.Scan(&num) // データを格納する変数のアドレスを指定
	sum := 5 + num
	fmt.Println(sum)
}
