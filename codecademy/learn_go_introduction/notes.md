# Codecademy: Learning Go Introduction

## Build and Run
- ```go build``` creates executable
- ```go run``` builds and runs but does not create an executable

View command documentation

```go
go doc <command>
go doc fmt
go doc fmt.Println
```

## Variables
- Literal (unnamed) values are values written into code as is (strings, numbers) (13.89, "Hello, World")
- Named values
    - Constants: ```const``` keyword
    - Variables: ```var``` keyword, followed by name of variable, followed by type (var lengthOfSong uint16 = 600)
- String concatenation

```go
nameOfSong + " is by: " + nameOfArtist + "."

fmt.Println("Purchase of", coolSneakers + niceNecklace, "with 8.875% sales tax", taxCalculation, "equal to", coolSneakers + niceNecklace + taxCalculation)
```

- All numeric variables have a value of 0 before assignment
- String variables have a default of "" or an empty string
- Boolean are false by default

- Inferring Type
    - Use declaration ```:=``` to declare a variable without explicitly stating its type
    - Do not need to use the "var" keyword

```go
nuclearMeltdownOccurring := true  //bool
radiumInGroundWater := 4.521  //float
daysSinceLastWorkplaceCatastrophe := 0  //int
externalMessage := "Everything is normal. Keep calm and carry on."  //string
```

- Can also use "var" keyword and just assign a value to infer type

```go
var nuclearMeltdownOccurring = true
var radiumInGroundWater = 4.521
var daysSinceLastWorkplaceCatastrophe = 0
var externalMessage = "Everything is normal. Keep calm and carry on."
```

```go
package main

import "fmt"

func main() {
  // Define daysOnVacation using := below:
  daysOnVacation := 7

  // Define hoursInDay using var and = below:
  var hoursInDay = 24

  fmt.Println("You have spent", daysOnVacation * hoursInDay, "hours on vacation.")
}
```

- Multiple Variable Declaration

```go
var part1, part2 string
part1 = "To be..."
part2 = "Not to be..."

quote, fact := "Bears, Beets, Battlestar Galactica", true
fmt.Println(quote) // Prints: Bears, Beets, Battlestar Galactica
fmt.Println(fact) // Prints: true
```