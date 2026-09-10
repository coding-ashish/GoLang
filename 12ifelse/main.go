package main

import "fmt"

func main() {
	fmt.Println("Lecture on if else statement")

	var num int

	fmt.Println("Enter the val of num :")
	// fmt.Scan()(&num)

	_, err := fmt.Scan(&num)

	if err != nil {
		// fmt.Printf("The number is %v:", num)
		fmt.Printf("An error occured")
	} else {
		fmt.Println("The entered number is :",num)
	}
}
