package main

import (
	"fmt"
)

func emptyStructPointer() {
	// pointers to empty structs are not equal,
	// until you print them.
	p1 := &struct{}{}
	p2 := new(struct{})
	p3 := new(struct{})

	fmt.Println(p1 == p2) // false
	fmt.Println(p2 == p3) // false

	fmt.Println(p1, p2)

	fmt.Println(p1 == p2) // true
	fmt.Println(p2 == p3) // false
}
