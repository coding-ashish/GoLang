package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input" //Walrus operator
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hey, please enter the rating of pizza : ")

	// comma ok syntax or comma error syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating : ", input)
}
