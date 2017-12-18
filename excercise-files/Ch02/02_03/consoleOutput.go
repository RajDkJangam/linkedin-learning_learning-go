package main

// implements a number of functions you can use to format and print strings to a variety of output targets
// see https://golang.org/pkg/fmt/

import "fmt"

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	// Print and Println functions concatenate arguments to a single string, separating with a " " (space)
	fmt.Println(str1, str2, str3)
	// go run consoleOutput.go
	// RETURN int, error object

	// Go functions can return more than one value simultaneously
	// Let's declare them for Println beforehand:
	// " := " inferred type assignment operator
	stringLength, err := fmt.Println(str1, str2, str3)
	// because we are telling Println to return those values (stringLength, err), we have to address them
	// 1. check if you get back a non-nill error object back
	// According to Go syntax rules, if you have a value that you've declare, you must address it in your code at some point
	if err == nil {
		fmt.Println("String length:", stringLength)
		// PRINT String length: 50
	}

	// if you know you are getting a variable returned but you don't want to address it, use " _ "
	// we will do that for err

	stringLength2, _ := fmt.Println(str1, str2, str3)
	fmt.Println("String length:", stringLength2)

	// You can also format strings from the fmt package
	// Printf() lets you create strings that have placeholders, known as "verbs," that define how values will be formatted.
	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", isTrue)
	// no implicit conversion
	fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber))

	// use Printf() to determine the data type
	fmt.Printf("Data types: %T, %T, %T, %T, and %T\n",
		str1, str2, str2, aNumber, isTrue)
	// RETURN string, string, string, int, and bool

	// Each function that outptuts to the console has versions for outputting to a string or other targets
	myString := fmt.Sprintf("Data types as var: %T, %T, %T, %T, and %T",
		str1, str2, str2, aNumber, isTrue)
	fmt.Print(myString)

	// see other verbs and functions at https://golang.org/pkg/fmt/
}
