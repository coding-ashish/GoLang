package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// Generating a random number between 1 and 6
	num := rand.IntN(6) + 1
	fmt.Println("The number is :", num)

	switch num {
	case 6:
		fmt.Println("Woohoo! You rolled a 6. You can unlock a new token from the base!")
		// fallthrough forces the code to immediately execute the next case block
		fallthrough 
	case 1, 2, 3, 4, 5:
		fmt.Println("You can move an active token forward by", num, "steps.")
	default:
		fmt.Println("Invalid dice roll!")
	}
}