package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of pointers")

	// var ptr *int
	// fmt.Println(ptr) // prints <nil>

	//Now lets try to store a memory address of a variable in a pointer

	myNumber := 7
	var ptr *int = &myNumber
	fmt.Println("The address at which pointer is pointing  is : ", ptr) // Prints the memory address of myNumber

	fmt.Println("The value stored in the ptr is : ", *ptr) // prints the value stored in the memory address stored in the ptr

	*ptr = *ptr * 2
	fmt.Println("The updated value stored in myNumber is : ", myNumber)
}
