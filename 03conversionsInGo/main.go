package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our app.")
	fmt.Println("Please enter your rating for our pizza (1-5) : ")

	rating := bufio.NewReader(os.Stdin)
	input, _ := rating.ReadString('\n')

	input = strings.TrimSpace(input)
	numRating, _ := strconv.ParseFloat(input, 64)
	
	//Now suppose we want to add 1 to whatever rating user has given us.

	fmt.Println("Your rating is : ", numRating+1)
	// with the above code we will get error because in the provided striong to stdconv function there is \n is also there. which we have to cut down

}
