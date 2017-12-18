// start with a package declaration
// do not use quotes. this is an identified, not a string
// the package name doesn't have to match the directory or source code file like Java
package main

// import the package you want to use in the application
// package names get wrapped in quotes

import (
	"fmt"
	"strings"
)

// import single package:
/*
	import "fmt"
*/

// import multiple packages:
/*
	import (
		"fmt"
		"package2"
	)
*/

// the main package should have a function named main
// main() is run on start automatically
func main() {
	// Println() has been exported
	// we know this because it is uppercase
	fmt.Println("Hello from Go!")
	// go run hellowworld.go
	fmt.Println(strings.ToUpper("Hello really loud!"))
	// go run hellowworld.go
}
