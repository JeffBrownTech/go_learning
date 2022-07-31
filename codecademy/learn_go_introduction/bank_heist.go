package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	isHeistOn := true
	eludedGuards := rand.Intn(100)

	fmt.Println("It's time to start the bank heist!")
	fmt.Println("The first step is to sneak past the guards.")

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("You were caught by the guards! Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("You successfully opened the vault. Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("You are no match for this vault's combination lock!")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Unfortunately, you were caught by the guards on the way out.")
		case 1:
			isHeistOn = false
			fmt.Println("The vault door shut behind you, locking you in.")
		case 2:
			isHeistOn = false
			fmt.Println("The silent alarm was tripped and police arrived to capture you.")
		case 3:
			isHeistOn = false
			fmt.Println("The getaway vehicle broke down and you were captured.")
		default:
			fmt.Println("Start the getaway car!")
		}

		if isHeistOn {
			amtStolen := 10000 + rand.Intn(100000)
			fmt.Printf("You stole $%d dollars from the bank!", amtStolen)
		} else {
			fmt.Println("Crime doesn't pay!")
		}
	}
}
