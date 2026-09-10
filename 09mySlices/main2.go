package main

import "fmt"

func main() {
	var courses = []string{"Python", "Java", "Ruby", "GoLang", "Kotlin"}
	fmt.Println(courses)

	courses = append(courses[:2], courses[3:]...)
	fmt.Println(courses)
}

