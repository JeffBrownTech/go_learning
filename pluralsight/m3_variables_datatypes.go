package main

import "fmt"

func main() {
	fmt.Println("Hello, Gophers!")

	var a string
	a = "foo"
	fmt.Println(a)

	// Variable and value assignment at once
	var b int = 99
	fmt.Println(b)

	// Implicit variable type
	var c = true
	fmt.Println(c)

	// Short-hand syntax for creating, declaring type, and setting value
	d := 3.1413 //assumes float64
	fmt.Println(d)

	// Type conversion of float64 to int
	// Go does not allow implicit type conversion, must always be explicit type conversion
	var e int = int(d)
	fmt.Println(e)

	// Multiple variable declaration
	aa, bb := 10, 5
	cc := aa + bb
	fmt.Println(cc)

	// Constant
	// Constants are untyped
	const constA = 42
	const bbb float32 = constA
	const ccc float64 = constA

	fmt.Println(bbb, ccc)

	// Pointers
	s := "Hello, world!"

	p := &s
	fmt.Println(p)  //Prints the memory address of s
	fmt.Println(*p) //Dereference to use the actual value

	// Change the value in s, dereferencing operation
	*p = "Hello, gophers!"
	fmt.Println(s) // Value in s has changed

	// new() allocates memory based on the type used
	p = new(string)
	fmt.Println(p, *p) // Prints new memory address and the value, which is currently blank/null
}
