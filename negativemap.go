package main

import "fmt"

func negativeMapSize() {
	n := -10
	m := make(map[string]int, n)
	fmt.Println(len(m)) // 0
}
