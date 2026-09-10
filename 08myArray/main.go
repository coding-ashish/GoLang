package main

import "fmt"

func main() {
	var fruitList [3]string

	fruitList[0] = "Apple"
	// fruitList[1]
	fruitList[2] = "Mango"

	fmt.Println("Our Array is : ", fruitList) // we noticed that there is an extra gap after Apple to indicate that there is a missing space there.

	//Lets see what it prints when we try to print fruitList[1]
	fmt.Println(fruitList[1])   // it orints just an empty space
	fmt.Println(len(fruitList)) // give the ans to be 3

	var vegList = []string{"Tomato", "Lauki", "Bhindi"} //If we don't tell the size of the array we are initialising, it becomes slice(a data type in go)
	// Also by this syntax to in itialise slice , we have to necessasiraliy initialise it at the time of declaration
	fmt.Println("The vegList is : ", vegList)
	fmt.Println("Size of the array is : ", len(vegList))
}
