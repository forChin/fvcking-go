package main

import "fmt"

type logger struct{}

func (*logger) log(msg string) {
	fmt.Println("[LOG]", msg)
}

func staticMethod() {
	// It is some kind of static method
	(*logger).log(nil, "hello")
}
