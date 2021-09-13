package main

import (
	. "fmt"
)

func withoutPackagePrefix() {
	// You can use exported function of package without package prefix
	// if specify alias for imported package as dot.

	Println("Hello World!")
}
