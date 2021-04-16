package main

import "fmt"

func nilToNULL() {
	// You can not just assign nil to variable, but
	// you can do your own NULL.
	null := interface{}(nil)

	// Variable null behaves exactly like nil.
	fmt.Printf("%T -- %v\n", null, null) // <nil> -- <nil>
	fmt.Printf("%T -- %v\n", nil, nil)   // <nil> -- <nil>
	fmt.Println(null == nil)             // true
}
