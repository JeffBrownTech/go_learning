package main

import "fmt"

func main() {
	// Array
	// Arrays are fixed size
	var arr [3]int   //array of 3 ints
	fmt.Println(arr) // [ 0 0 0 ]

	var arr1 = [3]int{1, 2, 3} //initialize the array on creation
	fmt.Println(arr1)

	fmt.Println(arr[1])  // 0
	fmt.Println(arr1[1]) // 2

	fmt.Println(len(arr1)) // 3

	// Creating array from existing performs a copy of the array
	stringArr1 := [3]string{"foo", "bar", "baz"}
	var stringArr2 = stringArr1
	stringArr1[0] = "quux"
	fmt.Println(stringArr1) // {"quux", "bar", "baz"}
	fmt.Println(stringArr2) // {"foo", "bar", "baz"}

	// Compare arrays
	// stringArr1 == stringArr2   // false

	// Slices
	// Slices don't hold their own data, they reference data held in an array
	// Slice is not a fixed size and held in memory; relfects any changes in the underlying array and vice versa
	var slices []int    //slices of ints, right now is nil
	fmt.Println(slices) // [] (nil)

	var slices1 = []int{1, 2, 3}
	fmt.Println(slices1)

	// Append to slide
	slices1 = append(slices1, 5, 10, 15)
	fmt.Println(slices1)

	// Use experimental slices package to remove indices
	slices1 = slices.Delete(slices1, 1, 3) // remove indices 1, 2 from slices1
	fmt.Println(slices1)

	// add dependencies to program
	// go get golang.org/x/exp/slices

	// Maps
	// Does not use 0-based index, you provide a key : value pair
	// Also dynamically sized
	var map1 map[string]int
	fmt.Println(map1) // map[] (nil)
	var map2 = map[string]int{"foo": 1, "bar": 2}
	fmt.Println(map2)

	fmt.Println(map2["foo"]) // looks up value for key "foo"
	map2["bar"] = 99         // update value for key "bar"
	delete(map2, "foo")
	map2["baz"] = 418 // add value to map
	// v, ok := map2["foo"] // comma okay syntax verifies presents

	// Copying maps copies references, so making changes to one changes the other
	// Use maps.Clone to truly clone a map to another variable
	// Maps are not comparable

	map3 := map[string][]string{
		"coffee": {"Coffee", "Espresso", "Cappucino"},
		"tea":    {"Hot Tea", "Chai Tea", "Chai Latte"},
	}
	fmt.Println(map3)

	// Structs

}
