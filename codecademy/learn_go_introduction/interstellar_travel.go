package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Println("Remaining Fuel:", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}

	return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
	if planet != "Earth" {
		fmt.Println("Welcome to", planet, "!")
	} else {
		fmt.Println("Welcome back to Earth!")
	}
}

// Create the function cantFly() here
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel

	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}

	return fuelRemaining
}

func flyToEarth(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel

	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet("Earth")
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}

	return fuelRemaining
}

func main() {
	// Test your functions!

	// Create `planetChoice` and `fuel`
	fuel := 1000000
	var currentPlanet, planetChoice string
	currentPlanet = "Earth"

	fmt.Println("Where would you like to go?")
	fmt.Scan(&planetChoice)

	// And then liftoff!
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
	currentPlanet = planetChoice

	fmt.Println("Returning to Earth...")
	fuel = flyToEarth(currentPlanet, fuel)
	fuelGauge(fuel)
}
