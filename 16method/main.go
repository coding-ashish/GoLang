package main

import "fmt"

func main() {

	Ashish := User{"Ashish", "extensionops@gmail.com", true}
	Ashish.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
}

//Remember one thing that the values passed to the method is by value and npt passed by reference
func (u User) GetStatus() {
	fmt.Println("The user status is :", u.Status)
}
