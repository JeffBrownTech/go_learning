# Codecademy: Learning Go Introduction

## Build and Run
- ```go build``` creates executable
- ```go run``` builds and runs but does not create an executable

- View command documentation

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

## FMT Package

- Verb: combination of an ampersand (%) and a letter
- String interpolation uses ```%v```

    ```go
    selection1 := "soup"
    selection2 := "salad"
    fmt.Printf("Do I want %v or %v?", selection1, selection2)
    ```
- Different verbs
  - ```%T```: prints the type of the argument
    ```go
    specialNum := 42
    fmt.Printf("This value's type is %T.", specialNum)
    // Prints: This value's type is int.

    quote := "To do or not to do"
    fmt.Printf("This value's type is %T.", quote)
    // Prints: This value's type is string.
    ```
  - ```%d```: interpolate a number into a string
    ```go
    votingAge := 18
    fmt.Printf("You must be %d years old to vote.", votingAge)
    // Prints: You must be 18 years old to vote.
    ```
  - ```%f```: interpolate a float into a string
    ```go
    gpa := 3.8
    fmt.Printf("You're averaging: %f.", gpa)
    // Prints: You're averaging 3.800000.

    gpa := 3.8
    fmt.Printf("You're averaging: %.2f.", gpa)
    // Prints: You're averaging 3.80.
    ```
- Sprint and Sprintln
  - Sprint & Sprintln formats using the default formats for its operands and returns the
    resulting string.
  - Primary difference to think about is it doesn't print a value out, it returns one for later use
    ```go
    grade := "100"
    compliment := "Great job!"
    teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)

    fmt.Print(teacherSays)
    // Prints: You scored a 100 on the test! Great job!

    quote = fmt.Sprintln("Look ma,", "no spaces!")
    fmt.Print(quote) // Prints Look ma, no spaces!

    // Can also use verbs and store them using Sprint
    template := "I wish I had a %v."
    pet := "puppy"

    var wish string

    // Add your code below:
    wish = fmt.Sprintf(template, pet)

    fmt.Println(wish)

    ```
- User Input
  - Use Scan() to get user input
    ```go
    fmt.Println("How are you doing?")

    var response string
    fmt.Scan(&response) //Takes first value before a space and stores it in response

    fmt.Printf("I'm %v.", response)
    ```

# Conditionals

## if Statement

- if statements
  ```go
  if <condition> {
    if-true
  } else {
    if-false
  }
  ```

  ```go
  rightTime := true
  rightPlace := true

  if rightTime && rightPlace {
    fmt.Println("We're outta here!")
  } else {
    fmt.Println("Be patient...")
  }
  ```

- if/else-if statements

  ```go
  position := 2

  if position == 1 {
    fmt.Println("You won the gold!")
  } else if position == 2 {
    fmt.Println("You got the silver medal.")
  } else if position == 3 {
    fmt.Println("Great job on bronze.")
  } else {
    fmt.Println("Sorry, better luck next time?")
  }
  ```

- switch statement
  ```go
  clothingChoice := "sweater"

  switch clothingChoice {
  case "shirt":
    fmt.Println("We have shirts in S and M only.")
  case "polos":
    fmt.Println("We have polos in M, L, and XL.")
  case "sweater":
    fmt.Println("We have sweaters in S, M, L, and XL.")
  case "jackets":
    fmt.Println("We have jackets in all sizes.")
  default:
    fmt.Println("Sorry, we don't carry that")
  }
  ```

## Scope Short Declaration Statement
- Can include short variable declaration before providing a condition in an ```if``` or ```switch``` statement
- The declared variable is scoped to the statement blocks and no available any where else

  ```go
  x := 8
  y := 9
  if product := x * y; product > 60 {
    fmt.Println(product, "  is greater than 60")
  }
  ```

# Functions

  ```go
  func <name>(param1 int32, param2 string) <returnType>{
    // code

    return <variable1>, <variable2>
  }

  func main() {
    var1, var2 := 1, 2

    var1, var2 = <name>(var1, var2)
  }
  ```

  - Defer: tells Go to run a function, but at the end of the current function; delays a function call to the end of the current scope

    ```go
    func calculateTaxes(revenue, deductions, credits float64) float64 {
      defer fmt.Println("Taxes Calculated!")
      taxRate := .06143
      fmt.Println("Calculating Taxes")

      if deductions == 0 || credits == 0 {
        return revenue * taxRate
      }


      taxValue := (revenue - (deductions * credits)) * taxRate
      if taxValue >= 0 {
        return taxValue
      } else {
        return 0
      }
    }
    ```