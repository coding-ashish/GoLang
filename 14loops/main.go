package main

import "fmt"

func main() {
	var days []string
	days = append(days, "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday")

	// this for loop we use when we have to traverse through some limited range and not full array
	for d := 0; d < len(days); d++ {
		if days[d] == "Friday" {
			fmt.Println("Wohoo! Its friday night we will party all night")
			break
		}
		fmt.Println(days[d])
	}

	//When we have to traverse the entiore array
	for i := range days {
		fmt.Println(days[i])
	}

	// Another for of a for loop is using it like a while loop
	val := 1
	for val <= 10 {
		fmt.Println(val)
		val++
	}
}
