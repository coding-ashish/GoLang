package main

import "fmt"

func main() {
	fmt.Println("Welcome the lecture struct")
	Ashish := User{"Ashish", "extensionops@gmail.com", true, 21}

	fmt.Println(Ashish)                           // Prints in very untidy manner
	fmt.Printf("User details are :%+v\n", Ashish) // Prints in good manner

	fmt.Printf("Name is %v and mail is %v", Ashish.Name, Ashish.Email) //tp print only specific details
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
