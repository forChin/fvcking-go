package main

import (
	"fmt"
)

func nilVar() {
	// You can declare nil variable.
	nil := 12
	fmt.Println("nil =", nil) // nil = 12
}

func nilType() {
	// Or declare nil type.
	type nil string
	var s nil = "Hello"
	fmt.Println(s) // "Hello"
}
