package main

import "fmt"

func main() {
	var name, pet string
	var birthYear, currYear int
	currYear = 2022

	fmt.Println("Practice using the fmt package!")
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %v!", name)
	fmt.Print("\n\n")

	fmt.Print("Enter your year of birth: ")
	fmt.Scan(&birthYear)
	age := currYear - birthYear
	ageOutput := fmt.Sprintf("You are approximately %d years old.", age)
	fmt.Println(ageOutput)
	fmt.Print("\n")

	fmt.Println("Do you like cats or dogs more?")
	fmt.Scan(&pet)
	fmt.Printf("%v are pretty cool!", pet)
	fmt.Print("\n")
}
