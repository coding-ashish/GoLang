package main

import "fmt"

func main() {
	fmt.Println("Welcome to the map lecture")

	var courses = make(map[string]string)
	courses["JS"] = "javascript"
	courses["RB"] = "ruby"
	courses["CPP"] = "c++"

	fmt.Println(courses)                          // does not print the values in the order they are inserted
	fmt.Println("JS stands for :", courses["JS"]) // will print javascript

}
