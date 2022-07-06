// .go files are organized into packages; package command tells which package this .go file belongs
// If 'package main' then this program is compiled into an executable
package main

// Import other packages, can use aliases
import (
	"fmt"
	t "time"
)

// Function declaration
func main() {
	fmt.Println("Hello World")
	fmt.Println(t.Now())
}
